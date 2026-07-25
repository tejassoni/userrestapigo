package requests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"regexp"
)

var birthdateValuePattern = regexp.MustCompile(`"birthdate"\s*:\s*(\d{4}-\d{2}-\d{2})`)

func DecodeCreateUserRequest(r *http.Request) (CreateUserRequest, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return CreateUserRequest{}, err
	}

	decode := func(data []byte) (CreateUserRequest, error) {
		var req CreateUserRequest
		err := json.NewDecoder(bytes.NewReader(data)).Decode(&req)
		return req, err
	}

	req, err := decode(body)
	if err == nil {
		return req, nil
	}

	repaired := birthdateValuePattern.ReplaceAllString(
		string(body),
		`"birthdate":"$1"`,
	)

	if repaired != string(body) {
		return decode([]byte(repaired))
	}

	return CreateUserRequest{}, err
}

func DecodeUpdateUserRequest(r *http.Request) (UpdateUserRequest, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return UpdateUserRequest{}, err
	}

	decode := func(data []byte) (UpdateUserRequest, error) {
		var req UpdateUserRequest
		err := json.NewDecoder(bytes.NewReader(data)).Decode(&req)
		return req, err
	}

	req, err := decode(body)
	if err == nil {
		return req, nil
	}

	repaired := birthdateValuePattern.ReplaceAllString(
		string(body),
		`"birthdate":"$1"`,
	)

	if repaired != string(body) {
		return decode([]byte(repaired))
	}

	return UpdateUserRequest{}, err
}
