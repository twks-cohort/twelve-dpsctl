package clients

import (
  // "encoding/json"
	"io/ioutil"
	"fmt"
	"dpsctl/clients/models"
	"net/http"
	"log"
	"strings"

	//"github.com/spf13/viper"
	//"github.com/tidwall/gjson"
)

func RequestDeviceCode() models.DeviceCode {
	fmt.Println("Requesting device code...")

	url := "https://dev-twdpsio.us.auth0.com/oauth/device/code"

	payload := strings.NewReader("client_id=B4jm7Wv4fjOEPqg1gjXIUUxEa6eg1HvB&scope=openid offline_access email&audience=https://mapi.twdps.digital/v1")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))




	// payload := strings.NewReader(fmt.Sprintf("client_id=%s&scope=%s&audience=%s",
	// 	viper.Get("LoginClientId").(string),
	// 	viper.Get("LoginScope").(string),
	// 	viper.Get("LoginAudience").(string)))

	// body, statusCode := submitPostRequest(viper.Get("DeviceCodeUrl").(string), payload)

	// fmt.Println("Status code: ", statusCode)
	// fmt.Println("Body: ", string(body))
	// // if statusCode != http.StatusOK {
	// // 	log.Fatalf("Status: %d\n%s: %s\n", statusCode, gjson.Get(string(body), "error"), gjson.Get(string(body), "error_description"))
	// // }

	var deviceCode models.DeviceCode
	// //json.Unmarshal(body, &deviceCode)

	return deviceCode
}

func submitPostRequest(url string, payload *strings.Reader) ([]byte, int) {
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	return body, res.StatusCode
}