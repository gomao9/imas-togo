package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/knakk/rdf"
	"github.com/knakk/sparql"
)

func main() {
	fmt.Println(idols())
}

func idols() []map[string]rdf.Term {
	const endpoint = "https://sparql.crssnky.xyz/spql/imas"
	const filename = "./query.rq"
	repo, err := sparql.NewRepo(endpoint)
	if err != nil {
		log.Fatal(err)
	}

	res, err := repo.Query(buildQuery(filename))
	if err != nil {
		log.Fatal(err)
	}

	return res.Solutions()
}

func buildQuery(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}
