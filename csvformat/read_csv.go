package csvformat

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

type Movie struct {
	Title    string
	Director string
	Year     int
}

// ReadCSV processes CSV data passed in as an io.Reader
func ReadCSV(b io.Reader) ([]Movie, error) {
	r := csv.NewReader(b)
	// Optional configuration options
	r.Comma = ';'
	r.Comment = '-'
	var movies []Movie

	// Ignore the header for now
	_, err := r.Read()
	if err != nil && err != io.EOF {
		return nil, err
	}

	// Loop until all records are processed
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		// Convert year to int
		year, err := strconv.Atoi(record[2])
		if err != nil {
			return nil, err
		}

		m := Movie{
			Title:    record[0],
			Director: record[1],
			Year:     year,
		}
		movies = append(movies, m)
	}

	return movies, nil
}

// AddMoviesFromText uses the CSV parser with a string
func AddMoviesFromText() error {
	// Example input string
	in := `Guardians of the Galaxy Vol. 2;James Gunn;2017
	Star Wars: Episode VIII;Rian Johnson;2017`
	b := bytes.NewBufferString(in)
	m, err := ReadCSV(b)
	if err != nil {
		return err
	}

	fmt.Printf("%#v\n", m)
	return nil
}
