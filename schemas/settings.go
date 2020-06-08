package schemas

import "github.com/hashicorp/terraform/helper/schema"

type Settings struct {
	SsoOneloginSecret     interface{} `json:"sso_onelogin_secret"`
	OraclePublicKey       string      `json:"oracle_public_key"`
	RoutedSubnet6         interface{} `json:"routed_subnet6"`
	RoutedSubnet6Wg       interface{} `json:"routed_subnet6_wg"`
	SsoSamlIssuerURL      interface{} `json:"sso_saml_issuer_url"`
	SsoAuthzeroAppID      interface{} `json:"sso_authzero_app_id"`
	Disabled              bool        `json:"disabled"`
	UsGovEast1AccessKey   interface{} `json:"us_gov_east_1_access_key"`
	ApNortheast2SecretKey interface{} `json:"ap_northeast_2_secret_key"`
	EmailServer           interface{} `json:"email_server"`
	Auditing              interface{} `json:"auditing"`
	Route53Zone           interface{} `json:"route53_zone"`
	SsoAuthzeroAppSecret  interface{} `json:"sso_authzero_app_secret"`
	EuCentral1AccessKey   interface{} `json:"eu_central_1_access_key"`
	AuthAPI               bool        `json:"auth_api"`
	SsoOktaAppID          interface{} `json:"sso_okta_app_id"`
	YubikeyID             interface{} `json:"yubikey_id"`
	Theme                 string      `json:"theme"`
	UsEast2AccessKey      interface{} `json:"us_east_2_access_key"`
	Secret                string      `json:"secret"`
	UsWest1AccessKey      interface{} `json:"us_west_1_access_key"`
	SsoOneloginMode       string      `json:"sso_onelogin_mode"`
	PinMode               string      `json:"pin_mode"`
	SsoOrg                string      `json:"sso_org"`
	SsoRadiusHost         interface{} `json:"sso_radius_host"`
	SsoOneloginID         interface{} `json:"sso_onelogin_id"`
	Username              string      `json:"username"`
	ServerCert            string      `json:"server_cert"`
	OtpSecret             interface{} `json:"otp_secret"`
	ServerPort            int         `json:"server_port"`
	SsoYubicoClient       string      `json:"sso_yubico_client"`
	UsWest2AccessKey      interface{} `json:"us_west_2_access_key"`
	ClientReconnect       bool        `json:"client_reconnect"`
	Sso                   string      `json:"sso"`
	ServerKey             string      `json:"server_key"`
	SsoClientCache        bool        `json:"sso_client_cache"`
	ApEast1SecretKey      interface{} `json:"ap_east_1_secret_key"`
	SaEast1SecretKey      interface{} `json:"sa_east_1_secret_key"`
	SsoDuoHost            interface{} `json:"sso_duo_host"`
	ApSoutheast2AccessKey interface{} `json:"ap_southeast_2_access_key"`
	Default               interface{} `json:"default"`
	OracleUserOcid        interface{} `json:"oracle_user_ocid"`
	SsoDuoToken           interface{} `json:"sso_duo_token"`
	CnNorth1AccessKey     interface{} `json:"cn_north_1_access_key"`
	SsoYubicoSecret       interface{} `json:"sso_yubico_secret"`
	SsoSamlURL            interface{} `json:"sso_saml_url"`
	Token                 string      `json:"token"`
	SsoCache              bool        `json:"sso_cache"`
	CnNorthwest1SecretKey interface{} `json:"cn_northwest_1_secret_key"`
	AcmeDomain            interface{} `json:"acme_domain"`
	OtpAuth               bool        `json:"otp_auth"`
	ApSoutheast2SecretKey interface{} `json:"ap_southeast_2_secret_key"`
	Route53Region         interface{} `json:"route53_region"`
	ApNortheast1SecretKey interface{} `json:"ap_northeast_1_secret_key"`
	EuNorth1SecretKey     interface{} `json:"eu_north_1_secret_key"`
	EuNorth1AccessKey     interface{} `json:"eu_north_1_access_key"`
	SsoMatch              []string    `json:"sso_match"`
	EuWest1SecretKey      interface{} `json:"eu_west_1_secret_key"`
	CaCentral1AccessKey   interface{} `json:"ca_central_1_access_key"`
	UsEast1AccessKey      interface{} `json:"us_east_1_access_key"`
	ReverseProxy          bool        `json:"reverse_proxy"`
	SsoOneloginAppID      interface{} `json:"sso_onelogin_app_id"`
	ApSouth1AccessKey     interface{} `json:"ap_south_1_access_key"`
	EmailPassword         interface{} `json:"email_password"`
	SsoGoogleKey          interface{} `json:"sso_google_key"`
	EuWest2AccessKey      interface{} `json:"eu_west_2_access_key"`
	ID                    string      `json:"id"`
	EuWest3AccessKey      interface{} `json:"eu_west_3_access_key"`
	RestrictImport        bool        `json:"restrict_import"`
	UsGovWest1AccessKey   interface{} `json:"us_gov_west_1_access_key"`
	Monitoring            interface{} `json:"monitoring"`
	ApNortheast2AccessKey interface{} `json:"ap_northeast_2_access_key"`
	InfluxdbURI           interface{} `json:"influxdb_uri"`
	UsEast1SecretKey      interface{} `json:"us_east_1_secret_key"`
	SsoSamlCert           interface{} `json:"sso_saml_cert"`
	SuperUser             bool        `json:"super_user"`
	SsoAuthzeroDomain     interface{} `json:"sso_authzero_domain"`
	SsoAzureAppSecret     interface{} `json:"sso_azure_app_secret"`
	ApEast1AccessKey      interface{} `json:"ap_east_1_access_key"`
	UsGovWest1SecretKey   interface{} `json:"us_gov_west_1_secret_key"`
	EuWest2SecretKey      interface{} `json:"eu_west_2_secret_key"`
	SsoOktaMode           string      `json:"sso_okta_mode"`
	SsoGoogleEmail        string      `json:"sso_google_email"`
	CnNorth1SecretKey     interface{} `json:"cn_north_1_secret_key"`
	PublicAddress         string      `json:"public_address"`
	CaCentral1SecretKey   interface{} `json:"ca_central_1_secret_key"`
	EmailUsername         interface{} `json:"email_username"`
	SsoDuoSecret          interface{} `json:"sso_duo_secret"`
	UsWest1SecretKey      interface{} `json:"us_west_1_secret_key"`
	CloudProvider         string      `json:"cloud_provider"`
	UsWest2SecretKey      interface{} `json:"us_west_2_secret_key"`
	CnNorthwest1AccessKey interface{} `json:"cn_northwest_1_access_key"`
	ApNortheast1AccessKey interface{} `json:"ap_northeast_1_access_key"`
	ApSouth1SecretKey     interface{} `json:"ap_south_1_secret_key"`
	UsGovEast1SecretKey   interface{} `json:"us_gov_east_1_secret_key"`
	EuWest3SecretKey      interface{} `json:"eu_west_3_secret_key"`
	ApSoutheast1AccessKey interface{} `json:"ap_southeast_1_access_key"`
	EuCentral1SecretKey   interface{} `json:"eu_central_1_secret_key"`
	SsoOktaToken          interface{} `json:"sso_okta_token"`
	UsEast2SecretKey      interface{} `json:"us_east_2_secret_key"`
	SsoAzureAppID         interface{} `json:"sso_azure_app_id"`
	ApSoutheast1SecretKey interface{} `json:"ap_southeast_1_secret_key"`
	PublicAddress6        interface{} `json:"public_address6"`
	EmailFrom             interface{} `json:"email_from"`
	SsoDuoMode            interface{} `json:"sso_duo_mode"`
	EuWest1AccessKey      interface{} `json:"eu_west_1_access_key"`
	SaEast1AccessKey      interface{} `json:"sa_east_1_access_key"`
	SsoAzureDirectoryID   interface{} `json:"sso_azure_directory_id"`
	SsoRadiusSecret       interface{} `json:"sso_radius_secret"`
}

func LoadSettings(d *schema.ResourceData) (sch *Settings) {
	sch = &Settings{
		SsoGoogleEmail: d.Get("sso_google_email").(string),
		Sso:            d.Get("sso").(string),
		SsoOrg:         d.Get("sso_org").(string),
		ID:             d.Id(),
	}

	ssoMatch := d.Get("sso_match").([]interface{})
	if ssoMatch != nil {
		sch.SsoMatch = []string{}
		for _, ssoM := range ssoMatch {
			sch.SsoMatch = append(sch.SsoMatch, ssoM.(string))
		}
	}
	return
}
