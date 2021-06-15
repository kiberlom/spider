package parsjsonagent

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type AgentReq struct {
	Agent []string `json:"agent"`
}

type Agent struct {
	Agent string `json:"agent"`
}

func Parse() {
	// open the file
	file, err := os.Open("../data/agent_request.json")

	//handle errors while opening
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	fileScanner := bufio.NewScanner(file)
	// handle first encountered error while reading
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	ag := AgentReq{}

	ij := 0
	ia := 0
	// read line by line
	for fileScanner.Scan() {

		ij++

		a := &Agent{}
		if err := json.Unmarshal([]byte(fileScanner.Text()), a); err != nil {
			log.Println(err)
			continue
		}

		if err := ag.duble(a.Agent); err == nil {
			ag.Agent = append(ag.Agent, a.Agent)
			ia++
		}

	}
	fmt.Println(ag.Agent)
	fmt.Printf("Всего: %d  Уникальных: %d", ij, ia)

	jsonString, _ := json.Marshal(ag)
	ioutil.WriteFile("../data/agent.json", jsonString, os.ModePerm)

	file.Close()

}

func (a *AgentReq) duble(s string) error {

	for _, v := range a.Agent {
		if v == s {
			return fmt.Errorf("12")
		}
	}

	return nil

}
