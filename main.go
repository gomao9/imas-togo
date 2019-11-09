package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/b4b4r07/go-finder"
	"github.com/knakk/sparql"
	"github.com/mitchellh/mapstructure"
)

func main() {
	fzf, err := finder.New("fzf")
	if err != nil {
		log.Fatal(err)
	}

	items := finder.NewItems()
	for _, idol := range idols() {
		items.Add(idol.displayName(), idol)
	}

	selectedItems, err := fzf.Select(items)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range selectedItems {
		fmt.Printf("height:%s\n", item.(idol).Height)
	}
}

type idol struct {
	Name              string
	NameKana          string
	AlternateName     string
	AlternateNameKana string
	GivenName         string
	GivenNameKana     string
	Height            string
	Weight            string
}

func (i *idol) displayName() string {
	var names = []string{}

	if i.Name != "" {
		names = append(names, fmt.Sprintf("%s(%s)", i.Name, i.NameKana))
	}
	if i.AlternateName != "" {
		names = append(names, fmt.Sprintf("%s(%s)", i.AlternateName, i.AlternateNameKana))
	}
	if len(names) == 0 && i.GivenName != "" {
		names = append(names, fmt.Sprintf("%s(%s)", i.GivenName, i.GivenNameKana))
	}
	return strings.Join(names, " ")
}

func mapToStruct(m map[string]string, val *idol) error {
	tmp, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = json.Unmarshal(tmp, val)
	if err != nil {
		return err
	}
	return nil
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
		newMap := map[string]string{}

		for key, value := range m {
			newMap[key] = value.String()
		}
		var result idol

		err := mapstructure.Decode(newMap, &result)
		if err != nil {
			log.Fatal(err)
		}
		idols = append(idols, result)
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
