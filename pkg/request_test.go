package pkg

import "testing"

func TestDoRequest(t *testing.T) {
	response, err := DoRequest("7108551682732100877")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response.Body)
}
