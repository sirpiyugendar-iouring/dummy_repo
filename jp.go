package main

import (
	Bytes "bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type ApiResponse struct {
	SchemeCode     int     `json:"schemecode"`
	InvDate        string  `json:"invdate"`
	InvEndData     string  `json:"invenddate"`
	SrNo           int     `json:"srno"`
	Fincode        int     `json:"fincode"`
	Asect_code     int     `json:"ASECT_CODE"`
	NoShares       int     `json:"noshares"`
	MktVal         float64 `json:"mktval"`
	Aum            float64 `json:"aum"`
	HoldPercentage float64 `json:"holdpercentage"`
	CompName       string  `json:"compname"`
	Sect_name      string  `json:"sect_name"`
	Asect_name     string  `json:"asect_name"`
	Rating         string  `json:"rating"`
	Flag           string  `json:"flag"`
	// Sect_code      int     `json:"sect_code"`
}

type ApiResponses struct {
	Url   string        `json:"url"`
	Table []ApiResponse `json:"Table"`
}

// holds the name of the output file. passed as a flag
var outFile string = "outFile"

func Main() {
	var responses ApiResponses
	jsonDir := os.ExpandEnv("$HOME/franklin_json")
	Mf_portfolioFile := filepath.Join(jsonDir, "Mf_portfolio.json")
	if outFile == "null" {
		log.Fatal("need a valid name")
	}

	// file name outFile is the temporary filename, will be deleted during make steps.
	franklinFile := outFile + ".json"
	franklinFile = filepath.Join(jsonDir, franklinFile)

	bytes, er := os.ReadFile(Mf_portfolioFile)
	if er != nil {
		log.Fatal(er)
	}
	er = json.Unmarshal(bytes, &responses)
	if er != nil {
		log.Fatal("unmarshal error", er)
	}
	file, er := os.OpenFile(franklinFile, os.O_WRONLY|os.O_CREATE, 0600)
	if er != nil {
		log.Fatal("can't prepare out file", er)
	}
	var fraknlinResponses ApiResponses
	for _, response := range responses.Table {
		if response.SchemeCode != 7864 {
			continue
		}

		fraknlinResponses.Table = append(fraknlinResponses.Table, response)

	}
	franklinBytes, er := json.Marshal(fraknlinResponses)
	if er != nil {
		log.Fatal("can't marshal", er)
	}
	var buf Bytes.Buffer
	buf.Write(franklinBytes)
	n, er := file.WriteString(buf.String())
	if er != nil {
		log.Fatal("can't write", er)
	}
	fmt.Println("Wrote", n, " bytes to", franklinFile)

}
