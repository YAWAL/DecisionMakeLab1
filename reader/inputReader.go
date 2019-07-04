package reader

import (
	"encoding/csv"
	"os"
	"strconv"
)

func ReadCSV(filename string) (input [3][3]int, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return [3][3]int{}, err
	}
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return [3][3]int{}, err
	}
	//var intMax [3][3]int
	for i, record := range records {
		for j, elem := range record {
			val, err := strconv.Atoi(elem)
			if err != nil {
				return [3][3]int{}, err
			}
			input[i][j] = val
			//intMax[i][j] = val
		}
	}
	return input, nil
	//return intMax, nil
}
