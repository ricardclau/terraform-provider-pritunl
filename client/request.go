package client

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

	"gopkg.in/mgo.v2/bson"
)

type PritunlClient struct {
	client *http.Client
	host   string
	token  string
	secret string
}

type Request struct {
	Method string
	Path   string
	Query  map[string]string
	Json   interface{}
}

func NewPritunlClient(host string, token string, secret string, httpClient *http.Client) *PritunlClient {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &PritunlClient{
		host:   host,
		token:  token,
		secret: secret,
		client: httpClient,
	}
}

func (c *PritunlClient) Do(r Request, respVal interface{}) error {

	url := "https://" + c.host + r.Path

	authTimestamp := strconv.FormatInt(time.Now().Unix(), 10)
	authNonce := bson.NewObjectId().Hex()
	authString := strings.Join([]string{
		c.token,
		authTimestamp,
		authNonce,
		r.Method,
		r.Path,
	}, "&")

	hashFunc := hmac.New(sha256.New, []byte(c.secret))
	hashFunc.Write([]byte(authString))
	rawSignature := hashFunc.Sum(nil)
	authSig := base64.StdEncoding.EncodeToString(rawSignature)

	var body io.Reader
	if r.Json != nil {
		data, e := json.Marshal(r.Json)
		if e != nil {
			return fmt.Errorf("client: Json marshal error: %v", e)
		}

		body = bytes.NewBuffer(data)
	}

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	req, err := http.NewRequest(r.Method, url, body)
	if err != nil {
		return fmt.Errorf("client: Failed to create client: %v", err)
	}

	if r.Query != nil {
		query := req.URL.Query()

		for key, val := range r.Query {
			query.Add(key, val)
		}

		req.URL.RawQuery = query.Encode()
	}

	req.Header.Set("Auth-Token", c.token)
	req.Header.Set("Auth-Timestamp", authTimestamp)
	req.Header.Set("Auth-Nonce", authNonce)
	req.Header.Set("Auth-Signature", authSig)

	if r.Json != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	log.Println(fmt.Sprintf("[DEBUG] Ricard Request: %s", req))

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("client: Request error: %v", err)
	}
	defer resp.Body.Close()

	info, _ := ioutil.ReadAll(resp.Body)
	log.Println(fmt.Sprintf("[DEBUG] Ricard Response: %v Body: %s", resp.StatusCode, string(info)))

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("client: Bad response status %d for req: %v", resp.StatusCode, req)
	}

	if respVal != nil {
		err = json.Unmarshal(info, &respVal)
		if err != nil {
			return fmt.Errorf("client: Failed to parse response, %v, Body: %s", err, info)
		}
	}

	return nil
}
