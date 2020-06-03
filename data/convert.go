package main

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	convertedPrefix = "converted-"
)

func main() {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if !strings.HasSuffix(f.Name(), ".csv") ||
			strings.HasPrefix(f.Name(), convertedPrefix) {
			continue
		}

		log.Printf("convert raw file: %s", f.Name())

		cf, _ := os.Open(f.Name())
		defer cf.Close()

		ncf, err := os.Create(convertedPrefix + f.Name())
		defer ncf.Close()
		if err != nil {
			log.Fatal(err)
		}

		r := csv.NewReader(cf)
		w := csv.NewWriter(ncf)
		recs, _ := r.ReadAll()
		for _, rec := range recs {
			if t, err := time.ParseDuration(rec[len(rec)-1]); err == nil {
				rec[len(rec)-1] = strconv.FormatInt(t.Nanoseconds(), 10)
			}

			w.Write(rec)
			w.Flush()
		}
	}
}
