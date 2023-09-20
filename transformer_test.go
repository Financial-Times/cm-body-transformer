package bodytransformer

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestTransformBody(t *testing.T) {
	tests := map[string]struct {
		BodyXMLFixture  string
		ExpectedFixture string
	}{
		"experimental": {
			BodyXMLFixture:  "testdata/10979399-ba25-45b9-b85d-776c1b75bfea/content.html",
			ExpectedFixture: "testdata/10979399-ba25-45b9-b85d-776c1b75bfea/expected.html",
		},
		"scrollable": {
			BodyXMLFixture:  "testdata/c0ac9d59-2285-4efc-b786-355a10ff3661/content.html",
			ExpectedFixture: "testdata/c0ac9d59-2285-4efc-b786-355a10ff3661/expected.html",
		},
		"recommended": {
			BodyXMLFixture:  "testdata/1bd99ff1-c8c3-4f28-b011-e2f8aeaba833/content.html",
			ExpectedFixture: "testdata/1bd99ff1-c8c3-4f28-b011-e2f8aeaba833/expected.html",
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			bodyXML := readFile(t, test.BodyXMLFixture)
			expected := readFile(t, test.ExpectedFixture)
			got, err := TransformBody(bodyXML)
			if err != nil {
				t.Fatalf("unexpected transformation error: %s", err.Error())
			}
			if strings.Compare(expected, got) != 0 {
				t.Fatalf("expected:\n%s\ngot:\n%s\n", expected, got)
			}
		})
	}
}

func readFile(t *testing.T, filename string) string {
	t.Helper()
	f, err := os.Open(filename)
	if err != nil {
		t.Fatalf("failed to open file %s: %s", filename, err.Error())
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		t.Fatalf("failed to read file %s: %s", filename, err.Error())
	}
	return string(data)
}
