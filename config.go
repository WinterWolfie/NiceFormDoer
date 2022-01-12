package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"io/ioutil"
	"strings"
)

func GenConfig(title, form string) Config {
	c := colly.NewCollector(
	//colly.AllowedDomains("https://docs.google.com"),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.StatusCode)
	})

	var config = Config{
		Title: title,
		Form:  form,
	}

	c.OnHTML(".freebirdFormviewerComponentsQuestionBaseRoot", func(e *colly.HTMLElement) {

		var question Question

		e.ForEach(".freebirdFormviewerComponentsQuestionBaseHeader", func(_ int, e *colly.HTMLElement) {
			question.Label = e.Text

			//fmt.Println(e.Text)
		})

		e.ForEach("input", func(_ int, e *colly.HTMLElement) {
			question.PostKey = strings.ReplaceAll(e.Attr("name"), "_sentinel", "")

			//fmt.Println(e.Attr("name"))
		})

		e.ForEach(".docssharedWizToggleLabeledLabelText", func(_ int, e *colly.HTMLElement) {
			var option = Option{
				Value:  e.Text,
				Chance: 1,
			}
			question.Options = append(question.Options, option)

			//fmt.Println(e.Text)
		})

		config.Questions = append(config.Questions, question)
	})

	err := c.Visit(form + "/viewform")
	if err != nil {
		panic(err)
	}

	SaveConfig(config)

	return config
}

func SaveConfig(config Config) {
	file, _ := json.MarshalIndent(config, "", " ")
	_ = ioutil.WriteFile("configs/"+config.Title+".json", file, 0644)
}

func ReadConfigGen() (string, string) {
	data, err := ioutil.ReadFile("configGen.json")
	if err != nil {
		panic(err)
	}

	type Info struct {
		Title string
		Form  string
	}
	// json data
	var obj Info

	// unmarshall it
	err = json.Unmarshal(data, &obj)
	if err != nil {
		panic(err)
	}

	return obj.Title, obj.Form
}

func ReadConfig(configName string) Config {
	data, err := ioutil.ReadFile("configs/" + configName + ".json")
	if err != nil {
		panic(err)
	}

	var config Config

	err = json.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	return config
}
