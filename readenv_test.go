package readenv

import (
	"reflect"
	"strings"
	"testing"
)

func TestReadEnv(t *testing.T) {
	got, err := ReadEnv(strings.NewReader("TEST_USER=JOHNDOE\nTEST_DB=postgres"))
	checkErr(t, err)
	want := EnvData(map[string]string{
		"TEST_USER": "JOHNDOE",
		"TEST_DB":   "postgres",
	})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func checkErr(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
