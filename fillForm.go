package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

func FillForm(config Config) {
	for i := 0; i < config.Quantity; i++ {

		params := CalcParams(config.Questions)

		fmt.Println(i, "request: ", params)

		PostRequest(Request{
			Form:   config.Form,
			Params: params,
		})

	}
}

func CalcParams(questions []Question) []Param {
	var params []Param
	//log.Println("what", questions)

	for _, question := range questions {

		var chance []string
		var totalChance int

		var param Param
		param.Key = question.PostKey

		//log.Println(question)

		for _, option := range question.Options {

			//log.Println("question.Options", totalChance)

			for i := 0; i < option.Chance; i++ {

				chance = append(chance, option.Value)
				totalChance++

				//log.Println("chance", chance)
			}
		}

		//log.Println("chance", totalChance)

		rand.Seed(time.Now().UnixNano())
		param.Value = chance[rand.Intn(totalChance)]

		//log.Println("err", param)

		params = append(params, param)

	}

	return params
}

func PostRequest(request Request) {
	//log.Println(request)

	params := url.Values{}

	for _, param := range request.Params {
		params.Add(param.Key, param.Value)
	}

	log.Println("params ", params)

	resp, err := http.PostForm(request.Form+"/formResponse", params)

	if err != nil {
		log.Printf("Request Failed: %s", err)
		return
	}

	err = resp.Body.Close()
	if err != nil {
		log.Printf("close failed: %s", err)
		return
	}

	/*body, err := ioutil.ReadAll(resp.Body)  // Log the request body
	bodyString := string(body)
	log.Print(bodyString)  // Unmarshal result*/
}
