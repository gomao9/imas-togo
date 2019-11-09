package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/knakk/sparql"
)

func readQuery() string {
	content, err := ioutil.ReadFile("./query.rq")
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func main() {
	repo, err := sparql.NewRepo("https://sparql.crssnky.xyz/spql/imas")
	if err != nil {
		log.Fatal(err)
	}

	res, err := repo.Query(readQuery())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.Solutions())
}
