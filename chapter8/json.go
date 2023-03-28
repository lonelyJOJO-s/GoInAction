package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type (
	gResult struct {
		GsearchResultClass string `json:"GsearchResultClass"`
		UnescapedURL       string `json:"unescapedUrl"`
		URL                string `json:"url"`
		VisibleURL         string `json:"visibleUrl"`
		CacheURL           string `json:"cacheUrl"`
		Title              string `json:"title"`
		TitleNoFormatting  string `json:"titleNoFormatting"`
		Content            string `json:"content"`
	}

	gResponse struct {
		ResponseData struct {
			Results []gResult `json:"results"`
		} `json:"responseData"`
	}
)

func main() {
	url := "http://baidu.com"

	resp, err := http.Get(url)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	defer resp.Body.Close()

	var gr gResponse
	err = json.NewDecoder(resp.Body).Decode(&gr)
	if err != nil {
		log.Println("error: ", err)
		return
	}

	fmt.Println(gr)

}


// unmarshal function deserialization the json string