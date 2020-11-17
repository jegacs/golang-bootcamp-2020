package models

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const COUNTRIES_API = "https://restcountries.eu/rest/v2/alpha/"
const FILE_PATH = "./"

type Country struct {
	Name        string
	Capital     string
	Area        float64
	Gini        float64
	NativeName  string
	NumericCode string
	Population  int
	Region      string
	SubRegion   string
}

var rows = []string{
	"Name",
	"Capital",
	"Area",
	"Gini",
	"NativeName",
	"NumericCode",
	"Population",
	"Region",
	"SubRegion",
}

func NewCountry(name string) *Country {
	return &Country{
		Name: name,
	}

}

func (c *Country) Store() error {
	f, err := os.Create(FILE_PATH + c.Name + ".csv")
	if err != nil {
		return err
	}

	defer f.Close()
	attributes := []string{}
	attributes = append(attributes, c.Name)
	attributes = append(attributes, c.Capital)
	attributes = append(attributes, fmt.Sprint(c.Area))
	attributes = append(attributes, fmt.Sprint(c.Gini))
	attributes = append(attributes, c.NativeName)
	attributes = append(attributes, c.NumericCode)
	attributes = append(attributes, fmt.Sprint(c.Population))
	attributes = append(attributes, c.Region)
	attributes = append(attributes, c.SubRegion)

	w := csv.NewWriter(f)
	err = w.Write(rows)
	if err != nil {
		return err
	}

	err = w.Write(attributes)
	if err != nil {
		return err
	}

	w.Flush()

	return nil
}

func (c *Country) Fetch() error {
	resp, err := http.Get(COUNTRIES_API + c.Name)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, c)
	if err != nil {
		return err
	}

	return nil
}

func (c *Country) ReadFile() error {
	return nil
}
