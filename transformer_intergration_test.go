//go:build manualintegration
// +build manualintegration

package bodytransformer

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"
)

var (
	apiKey            = flag.String("apiKey", "", "API key for accessing CAPI, used for public content API")
	basicAuthUser     = flag.String("basicAuthUser", "", "basic auth user for accessing UPP clusters, used for getting content from document store")
	basicAuthPassword = flag.String("basicAuthPassword", "", "basic auth password for accessing UPP clusters, used for getting content from document store")
)

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}

// TestBodyTransformation tests iterates over a list of content uuids and checks whether the output of the library is
// the same as the output body form public content API.
//
// There are the following known differences between the transformation done byt the current library and the one
// from public content API/enriched content API:
// - if FT specific tags like ft-concept or ft-content are self closed (<ft-concept ... /> or <ft-contetn ... />)
// in the document store api version of the body the current implementation leaves them self closed
// as opposed to the output of the public content API where they are transformed to start and end tag (<ft-concept type="..." url="..."></ft-concept>).
// E.g. "72eebb8e-0bf0-11e8-8eb7-42f857ea9f09", "83ad52a8-59a5-11e7-9bc8-8055f264aa8b", "6cf68edc-f686-11e9-9ef3-eca8fc8f2d65"
// "5dc60b96-669c-11ea-800d-da70cff6e4d3", "f1dcb508-3bc7-11e7-ac89-b01cc67cfeec"
//
// - if the body of a content contains html escape sequences the current implementation html-unescapes them as opposed to
// the output of the public content API where only some characters are escaped; we are un-escaping them because if the user
// of the library needs to escape all characters consistently, they will end up escaping the already escaped characters
// them 2 times and would have invalid character sequence as a result
// E.g. "2617b220-1d2b-430c-842d-ebcc5be7e169", "d28aa2c6-a598-11db-a4e0-0000779e2340", "261cd90e-b620-11df-a784-00144feabdc0"
// "144e1502-3762-11e6-a780-b48ed7b6126f", "28d2f611-2e19-4c7a-8e19-06228f567cb8", "34905308-81a4-11e0-8a54-00144feabdc0"
func TestBodyTransformation(t *testing.T) {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	// set of uuids from random time periods
	testUUIDs := []string{
		"aabd25b0-0be7-11e6-b0f1-61f222853ff3",
		"c7108d38-929e-11da-977b-0000779e2340",
		"389ebb32-3824-11dd-aabb-0000779fd2ac",
		"937f7f84-5a2e-11dc-9bcd-0000779fd2ac",
		"ccf1e97a-f2a3-11db-a454-000b5df10621",
		"5a069168-55e8-11e3-96f5-00144feabdc0",
		"e7e41134-e6d4-11db-9034-000b5df10621",
		"cdbbe2ea-acc0-11e4-beeb-00144feab7de",
		"cdd1ea06-7cc0-11e0-994d-00144feabdc0",
		"14ffa01a-e4d9-11e9-9743-db5a370481bc",
		"a529b3f4-2cf9-11e8-a34a-7e7563b0b0f4",
		"2a4fff9a-b0a2-11dd-8915-0000779fd18c",
		"2448256e-5a6c-11e2-a02e-00144feab49a",
		"3bc0a8de-25c8-11dd-b510-000077b07658",
		"7460443e-b4f1-11e7-8007-554f9eaa90ba",
		"d03898c2-19b3-11e1-ba5d-00144feabdc0",
		"6a61d458-e7c7-11e2-babb-00144feabdc0",
		"ef546278-21a6-11db-b650-0000779e2340",
		"f5deb440-2f46-11e7-9555-23ef563ecf9a",
		"6ca69342-4d2b-11e8-97e4-13afc22d86d4",
		"59c0ed22-7e52-11e5-a1fe-567b37f80b64",
		"fd38117c-1fcf-11e7-a454-ab04428977f9",
		"fa27388e-523f-11e3-8c42-00144feabdc0",
		"09c4e090-3865-11e0-959c-00144feabdc0",
		"69612f36-755c-11e8-aa31-31da4279a601",
		"b93be716-c981-11de-a071-00144feabdc0",
		"1d7abc10-409e-11de-8f18-00144feabdc0",
		"9dffdb8f-f00e-4305-a69a-158b845f6970",
	}

	for _, uuid := range testUUIDs {
		publicAPIBody := getPublicContentBody(t, uuid, httpClient)
		docStoreBody := getDocumentStoreBody(t, uuid, httpClient)

		transfomedBody, err := TransformBody(docStoreBody)
		if err != nil {
			t.Fatalf("failed to transform content body: %v", err)
		}

		// Remove empty new lines for easier comparison later
		publicAPIBody = removeEmptyLines(publicAPIBody)

		if publicAPIBody != transfomedBody {
			t.Errorf(
				"for %s transformed body is different than the content body returned by public content API \n expected \n %s \n got \n %s \n",
				uuid,
				publicAPIBody,
				transfomedBody,
			)
		}
	}
}

func getDocumentStoreBody(t *testing.T, uuid string, client *http.Client) string {
	t.Helper()

	urlTempl := "https://upp-prod-delivery-eu.ft.com/__document-store-api/content/%s"
	url := fmt.Sprintf(urlTempl, uuid)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatalf("failed to create document store get request: %v", err)
	}

	req.SetBasicAuth(*basicAuthUser, *basicAuthPassword)

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("failed to perform document store get request: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("failed to get item '%s' from document store: '%s'", uuid, resp.Status)
	}

	content := struct {
		BodyXML string `json:"bodyXML"`
		Body    string `json:"body"`
	}{}
	err = json.NewDecoder(resp.Body).Decode(&content)
	if err != nil {
		t.Fatalf("failed to decode document store response body: %v", err)
	}

	if content.BodyXML != "" {
		return content.BodyXML
	}
	if content.Body != "" {
		return content.Body
	}

	t.Fatalf("no body or body xml fields returned from document store for uuid %v", uuid)
	return ""
}

func getPublicContentBody(t *testing.T, uuid string, client *http.Client) string {
	t.Helper()

	urlTempl := "https://api.ft.com/content/%s?apiKey=%s"
	url := fmt.Sprintf(urlTempl, uuid, *apiKey)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatalf("failed to create public content get request: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("failed to perform public content get request: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("failed to get item '%s' from public content api: '%s'", uuid, resp.Status)
	}

	content := struct {
		BodyXML string `json:"bodyXML"`
	}{}
	err = json.NewDecoder(resp.Body).Decode(&content)
	if err != nil {
		t.Fatalf("failed to decode public content reponse: %v", err)
	}

	return content.BodyXML
}
