package models

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

const COUNTRIES_API = "https://restcountries.eu/rest/v2/alpha/"

var FILE_PATH = os.Getenv("COUNTRIES_FILE")

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
	Code        string
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

func NewCountry(code string) *Country {
	return &Country{
		Code: code,
	}

}

func (c *Country) Store() error {
	f, err := os.Create(FILE_PATH + c.Code + ".csv")
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
	resp, err := http.Get(COUNTRIES_API + c.Code)
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
	fileName := c.Code + ".csv"
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	r := csv.NewReader(file)
	i := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if i == 1 {
			c.Name = record[0]
			c.Capital = record[1]
			area, _ := strconv.ParseFloat(record[2], 64)
			c.Area = area
			gini, _ := strconv.ParseFloat(record[3], 64)
			c.Gini = gini
			c.NativeName = record[4]
			c.NumericCode = record[5]
			population, _ := strconv.ParseInt(record[6], 10, 64)
			c.Population = int(population)
			c.Region = record[7]
			c.SubRegion = record[8]
		}
		i++
	}
	return nil
}
