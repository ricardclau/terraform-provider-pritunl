package request

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/pritunl/terraform-provider-pritunl/schemas"
	"gopkg.in/mgo.v2/bson"
)

var client = &http.Client{
	Timeout: 2 * time.Minute,
}

type Request struct {
	Method string
	Path   string
	Query  map[string]string
	Json   interface{}
}

func (r *Request) Do(prvdr *schemas.Provider, respVal interface{}) (resp *http.Response, err error) {

	url := "https://" + prvdr.PritunlHost + r.Path

	authTimestamp := strconv.FormatInt(time.Now().Unix(), 10)
	authNonce := bson.NewObjectId().Hex()
	authString := strings.Join([]string{
		prvdr.PritunlToken,
		authTimestamp,
		authNonce,
		r.Method,
		r.Path,
	}, "&")

	hashFunc := hmac.New(sha256.New, []byte(prvdr.PritunlSecret))
	hashFunc.Write([]byte(authString))
	rawSignature := hashFunc.Sum(nil)
	authSig := base64.StdEncoding.EncodeToString(rawSignature)

	var body io.Reader
	if r.Json != nil {
		data, e := json.Marshal(r.Json)
		if e != nil {
			err = fmt.Errorf("request: Json marshal error: %v", e)

			return
		}

		body = bytes.NewBuffer(data)
	}

	// Disable SSL Check for local testing
	// if prvdr.PritunlHost == "localhost" {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	// }

	req, err := http.NewRequest(r.Method, url, body)
	if err != nil {
		err = fmt.Errorf("request: Failed to create request: %v", err)

		return
	}

	if r.Query != nil {
		query := req.URL.Query()

		for key, val := range r.Query {
			query.Add(key, val)
		}

		req.URL.RawQuery = query.Encode()
	}

	req.Header.Set("Auth-Token", prvdr.PritunlToken)
	req.Header.Set("Auth-Timestamp", authTimestamp)
	req.Header.Set("Auth-Nonce", authNonce)
	req.Header.Set("Auth-Signature", authSig)

	if r.Json != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	log.Printf("[DEBUG] Sending Request: %s", req)

	resp, err = client.Do(req)
	if err != nil {
		err = fmt.Errorf("request: Request error: %v", err)

		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 || resp.StatusCode == 401 {
		return
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		err = fmt.Errorf("request: Bad response status %d", resp.StatusCode)

		return
	}

	if respVal != nil {
		info, _ := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(info, &respVal)
		log.Printf("[DEBUG] Returned Request: %s", respVal)
		// if r.Path == "/settings" {
		// 	var settingsResp *schemas.Settings
		// 	mapstructure.Decode(respVal, &settingsResp)
		// 	log.Printf("[DEBUG] Returned Response settings: %s", settingsResp)
		// 	return
		// }
		if err != nil {
			err = fmt.Errorf("request: Failed to parse response, %v", err)

			return
		}
	}

	return
}
