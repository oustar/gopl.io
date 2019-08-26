package baidutsn

// https://openapi.baidu.com/oauth/2.0/token
// 	?grant_type=client_credentials&
// 	client_id=PxhL9BAMxFxBetzvk6jRWzgW&
// 	client_secret=RkrffGG1UCFodWqVUWWYaBuAzGjp5jVS

// {
// 	"access_token":"24.65a77813be752d44786f0642913719b7.2592000.1568873165.282335-16772699",
// 	 "session_key":"9mzdDcOu2hAUbh5Gcc3C05rvb5h7RkBcXN+FgxvoSsxITz1IXKik2dbq+snnPDSzNcJKd3uVy0qbBY2bUWiPUftJ5wz27g==",
// 	 "scope":"audio_voice_assistant_get brain_enhanced_asr audio_tts_post public brain_all_scope wise_adapt lebo_resource_base lightservice_public hetu_basic lightcms_map_poi kaidian_kaidian ApsMisTest_Test\u6743\u9650 vis-classify_flower lpq_\u5f00\u653e cop_helloScope ApsMis_fangdi_permission smartapp_snsapi_base iop_autocar oauth_tp_app smartapp_smart_game_openapi oauth_sessionkey smartapp_swanid_verify smartapp_opensource_openapi smartapp_opensource_recapi fake_face_detect_\u5f00\u653eScope",
// 	 "refresh_token":"25.b1a227efc03cc7a7c4cd8e7e6ef0a90d.315360000.1881641165.282335-16772699",
// 	 "session_secret":"c1adc6863aefacbad322774c5c8e7459",
// 	 "expires_in":2592000
// 	}

// TokenURL is the url of token
const TokenURL = "https://openapi.baidu.com/oauth/2.0/token"

// ClientID is the API KEY
const ClientID = "PxhL9BAMxFxBetzvk6jRWzgW"

// ClientSecret is the API Secret
const ClientSecret = "RkrffGG1UCFodWqVUWWYaBuAzGjp5jVS"

// Token is the response of OAUTH
type Token struct {
	AccessToken string `json:"access_token"`
	SessionKey  string `json:"session_key"`
	Scope       string
}

// TsnURL is the text2audio URL.
const TsnURL = "http://tsn.baidu.com/text2audio"
