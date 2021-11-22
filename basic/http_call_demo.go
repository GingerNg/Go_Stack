package basic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

/*
https://www.cnblogs.com/mafeng/p/7068837.html
*/

func HttpDemo() {
	//d_json := map[string]string{
	//	"texts1": bytelist2string(texts1),
	//	"texts2": bytelist2string(texts2),
	//}
	//httpPost("http://192.168.1.226:35998/")
	httpPostForm("http://192.168.1.226:35998/")
}

func httpPostForm(t_url string) {
	d_param := map[string]string{
		"query": "12121",
	}
	d_param_byte, _ := json.Marshal(d_param)
	resp, err := http.PostForm(t_url,
		url.Values{"p": {string(d_param_byte)}})

	if err != nil {
		// handle error
		fmt.Println("----")
		fmt.Println(err)
		fmt.Println(resp)
	}
	if resp != nil{
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// handle error
			fmt.Println(err)
		}
		fmt.Println(string(body))
	}
}

func httpPost(url string) {
	resp, err := http.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader("name=cjb"))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func httpRequestPost(url string) {

}
