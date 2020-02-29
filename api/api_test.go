package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andresdb91/test-meli-magneto/db/dbtest"
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
}

func TestHttpGetStats(t *testing.T) {}
