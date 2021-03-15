package readenv

import (
	"io"
	"io/ioutil"
	"os"
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

func TestEnvDataGet(t *testing.T) {
	tests := []struct {
		name        string
		reader      io.Reader
		key         string
		wantVal     string
		wantErr     bool
		expectedErr error
	}{
		{
			name:        "get env variable by valid key",
			reader:      strings.NewReader("dbtype=postgres"),
			key:         "dbtype",
			wantVal:     "postgres",
			wantErr:     false,
			expectedErr: nil,
		},
		{
			name:        "get env variable by invalid key",
			reader:      strings.NewReader("trip=chancetherapper"),
			key:         "trips",
			wantVal:     "",
			wantErr:     true,
			expectedErr: errKeyNotInEnvData("trips"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ReadEnv(tc.reader)
			if gotVal, _ := got.Get(tc.key); gotVal != tc.wantVal {
				t.Errorf("wanted %v, got %v", tc.wantVal, gotVal)
			}
			if tc.wantErr {
				if err := tc.expectedErr; err == nil {
					t.Errorf("expected an error %v, got nil", tc.expectedErr)
				}
			}
			checkErr(t, err)
		})
	}
}

func TestFile(t *testing.T) {
	err := os.Mkdir("tmp", 0777)
	checkErr(t, err)
	tmpFile, err := ioutil.TempFile("tmp", "*.env")
	checkErr(t, err)
	_, err = tmpFile.WriteString("TEST_USER=johndoe")
	checkErr(t, err)
	defer os.Remove("tmp")
	defer os.Remove(tmpFile.Name())

	got, err := File(tmpFile.Name())
	checkErr(t, err)
	want := EnvData{"TEST_USER": "johndoe"}

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
