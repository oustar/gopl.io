package baidutsn

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// GetToken get user acess token
func GetToken() (string, error) {
	v := url.Values{}

	v.Add("grant_type", "client_credentials")
	v.Add("client_id", ClientID)
	v.Add("client_secret", ClientSecret)

	req, err := http.NewRequest("GET", TokenURL, nil)
	if err != nil {
		return "", err
	}
	req.URL.RawQuery = v.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("get token failed: %s", resp.Status)
	}

	var data Token
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}
	return data.AccessToken, nil
}

// tex	必填	合成的文本，使用UTF-8编码。小于2048个中文字或者英文数字。（
//		文本在百度服务器内转换为GBK后，长度必须小于4096字节）
// tok	必填	开放平台获取到的开发者access_token（见上面的“鉴权认证机制”段落）
// cuid	必填	用户唯一标识，用来计算UV值。建议填写能区分用户的机器 MAC 地址或 IMEI 码，长度为60字符以内
// ctp	必填	客户端类型选择，web端填写固定值1
// lan	必填	固定值zh。语言选择,目前只有中英文混合模式，填写固定值zh
// spd	选填	语速，取值0-15，默认为5中语速
// pit	选填	音调，取值0-15，默认为5中语调
// vol	选填	音量，取值0-15，默认为5中音量
// per（基础音库）	选填	度小宇=1，度小美=0，度逍遥=3，度丫丫=4
// per（精品音库）	选填	度博文=106，度小童=110，度小萌=111，度米朵=103，度小娇=5
// aue	选填	3为mp3格式(默认)； 4为pcm-16k；5为pcm-8k；6为wav（内容同pcm-16k）;
//		注意aue=4或者6是语音识别要求的格式，但是音频内容不是语音识别要求的自然人发音，所以识别效果会受影响。

// Transfer transfter text to audio.
func Transfer(text string) ([]byte, error) {
	var v = url.Values{
		"tex":  {""},
		"tok":  {""},
		"cuid": {"owen-ubuntu"},
		"ctp":  {"1"},
		"lan":  {"zh"},
	}

	token, err := GetToken()
	if err != nil {
		return nil, err
	}

	v.Set("tex", text)
	v.Set("tok", token)

	req, err := http.NewRequest("POST", TsnURL, strings.NewReader(v.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get token failed: %s", resp.Status)
	}

	if strings.Contains(resp.Header.Get("Content-Type"), "audio") {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return body, nil
	}

	return nil, fmt.Errorf("transfer error")

}
