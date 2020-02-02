package sqrt

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"testing"
)

func almostEqual(v1, v2 float64) bool {
	return Abs(v1-v2) <= 0.001
}

func TestSimple(t *testing.T) {
	file, err := os.Open("sqrt_cases.csv")
	if err != nil {
		t.Fatalf("can`t open case file - %s\n", err)
	}
	defer file.Close()
	rdr := csv.NewReader(file)
	for {
		record, err := rdr.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			t.Fatalf("error reading case file - %s\n", err)
		}

		val, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			t.Fatalf("bad value - %s\n", record[0])
		}

		expected, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			t.Fatalf("bad value - %s\n", record[1])
		}
		t.Run(fmt.Sprintf("%f", val), func(t *testing.T) {
			out, err := Sqrt(val)
			if err != nil {
				t.Fatal(err)
			}
			if !almostEqual(out, expected) {
				t.Fatalf("%f != %f", out, expected)
			}
		})
	}
	// val, err := Sqrt(2)
	//
	// if err != nil {
	// 	t.Fatalf("error in calculation - %s\n", err)
	// }
	//
	// if !almostEqual(val, 1.414214) {
	// 	t.Fatalf("bad value - %f\n", val)
	// }
}

// type testCase struct {
// 	value    float64
// 	expected float64
// }
//
// func TestMany(t *testing.T) {
// 	testCase := []testCase{
// 		{0.0, 0.0},
// 		{2.0, 1.414214},
// 		{9.0, 3.0},
// 	}
//
// 	for _, tc := range testCase {
// 		t.Run(fmt.Sprintf("%f", tc.value), func(t *testing.T) {
// 			out, err := Sqrt(tc.value)
// 			if err != nil {
// 				t.Fatal("error")
// 			}
// 			if !almostEqual(out, tc.expected) {
// 				t.Fatalf("%f ! = %f", out, tc.expected)
// 			}
// 		})
// 	}
// }
