package api

import (
	"bytes"
	"encoding/json"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andresdb91/test-meli-magneto/db/dbtest"
	"github.com/andresdb91/test-meli-magneto/hll"
	"github.com/andresdb91/test-meli-magneto/hll/hlltest"
)

func TestHttpPostMutant(t *testing.T) {
	cases := []struct {
		in   []byte
		want int
	}{
		{
			[]byte(`{"dna":["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]}`),
			http.StatusOK,
		},
		{
			[]byte(`{"dna":["ATGCGA","CTGTAC","TTATGT","AGAAGG","CCACTA","TCACTG"]}`),
			http.StatusForbidden,
		},
	}

	router := setupRouter()
	hll.Client = hlltest.SetupMockRedis()
	dbtest.SetupMockDb()

	for _, c := range cases {
		req, err := http.NewRequest("POST", "/mutant", bytes.NewBuffer(c.in))

		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		if status := rec.Code; status != c.want {
			t.Errorf("checkMutant returned wrong code, got: %v, want: %v", status, c.want)
		}
	}

	dbtest.CleanupMockDb()
	hlltest.CleanupMockRedis(hll.Client)
}

func TestMutantFormatCheck(t *testing.T) {
	cases := []struct {
		in   []byte
		want int
	}{
		{
			[]byte(`{"dna":["CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]}`),
			http.StatusBadRequest,
		},
		{
			[]byte(`{"dna":["ATGCGA","CTGTAC","TTATGT","AGA","CCACTA","TCATG"]}`),
			http.StatusBadRequest,
		},
		{
			[]byte(`{"dna":["ATGCRA","CTGTAC","TTATGT","AGA","CCACTA","TCATG"]}`),
			http.StatusBadRequest,
		},
	}

	router := setupRouter()

	for _, c := range cases {
		req, err := http.NewRequest("POST", "/mutant", bytes.NewBuffer(c.in))

		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		if status := rec.Code; status != c.want {
			t.Errorf("checkMutant returned wrong code, got: %v, want: %v", status, c.want)
		}
	}
}

func TestHttpGetStats(t *testing.T) {
	hll.Client = hlltest.SetupMockRedis()
	countH, countM := hlltest.PopulateMockRedis(hll.Client)
	ratio := float64(countM) / float64(countH)

	expected := []struct {
		key   string
		value float64
	}{
		{
			key:   "count_human_dna",
			value: float64(countH),
		},
		{
			key:   "count_mutant_dna",
			value: float64(countM),
		},
		{
			key:   "ratio",
			value: math.Round(ratio*10) / 10,
		},
	}

	router := setupRouter()

	req, err := http.NewRequest("GET", "/stats", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("getStats returned wrong code, got: %v, want: %v\n", status, http.StatusOK)
	}

	var res map[string]float64
	err = json.Unmarshal([]byte(rec.Body.String()), &res)

	for _, c := range expected {
		value, exists := res[c.key]
		if exists {
			if value != c.value {
				t.Errorf("Incorrect value for key: %v, got: %v, want: %v\n", c.key, value, c.value)
			}
		} else {
			t.Errorf("Expected key not found: %v\n", c.key)
		}
	}

	hlltest.CleanupMockRedis(hll.Client)
}
