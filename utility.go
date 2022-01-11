package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

func DoIt(request Request2, quantity int) {

	for i := 0; i < quantity; i++{
		params := CalcParams(request.Questions)

		PostRequest(Request{
			Form:   request.Form,
			Params: params,
		})

	}
}



func AppendOption(answer *Question, value string, chance int) {
	answer.Options = append(answer.Options, Option{
		Value:  value,
		Chance: chance,
	})
}

func CalcParams(questions []Question) []Param {

	var params []Param


	//log.Println("what", questions)

	for _, question := range questions {

		var chance []string
		var totalChance int

		var param Param
		param.Key = question.Key

		//log.Println(question)

		for _, option := range question.Options {

			//log.Println("question.Options", totalChance)

			for i := 0; i < option.Chance; i++{

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
		log.Println(request)


		params := url.Values{}

		for _, param := range request.Params {
			params.Add(param.Key, param.Value)
		}

		resp, err := http.PostForm(request.Form, params)

		if err != nil {
			log.Printf("Request Failed: %s", err)
			return
		}

		err = resp.Body.Close()
		if err != nil {
			log.Printf("close failed: %s", err)
			return
		}

		body, err := ioutil.ReadAll(resp.Body)  // Log the request body
		bodyString := string(body)
		log.Print(bodyString)  // Unmarshal result


}