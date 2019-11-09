package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/b4b4r07/go-finder"
	"github.com/knakk/rdf"
	"github.com/knakk/sparql"
)

func main() {
	fzf, err := finder.New("fzf")
	if err != nil {
		log.Fatal(err)
	}

	items := finder.NewItems()
	for _, idol := range idols() {
		items.Add(idol.name, idol)
	}

	selectedItems, err := fzf.Select(items)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range selectedItems {
		fmt.Printf("name:%s\n", item.(map[string]rdf.Term)["name"].String())
	}

	// fmt.Println(idols())
}

type idol struct {
	name string
}

func name(idol map[string]rdf.Term) string {
	var name rdf.Term
	switch {
	case idol["name"] != nil:
		name = idol["name"]
	case idol["alternateName"] != nil:
		name = idol["alternateName"]
	case idol["givenName"] != nil:
		name = idol["givenName"]
	default:
		return ""
	}
	return name.String()
}

func idols() (idols []idol) {
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

	for _, m := range res.Solutions() {
		idols = append(idols, idol{name: name(m)})
	}

	return
}

func buildQuery(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}
