package imgur

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"
)

// Delete Image from Ingur using Client-ID
//	info, status, err := DeleteImageUnAuthed("deleteHash")
func (client *Client) DeleteImageUnAuthed(hash string) (*ImageInfoWithoutData, int, error) {
	URL := "https://api.imgur.com/3/image/" + hash

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	err := writer.Close()
	if err != nil {
		return nil, -1, err
	}

	req, err := http.NewRequest("DELETE", URL, payload)
	if err != nil {
		return nil, -1, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Client-ID %s", client.Imgur.ClientID))
	req.Header.Set("Content-Type", writer.FormDataContentType())

	res, err := client.HTTPClient.Do(req)
	if err != nil {
		return nil, -1, err
	}
	defer res.Body.Close()

	if res.StatusCode >= 300 {
		return nil, res.StatusCode, errors.New("Imgur Failed with Status: " + strconv.Itoa(res.StatusCode))
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, -1, err
	}

	dec := json.NewDecoder(bytes.NewReader(body))

	var i ImageInfoWithoutData

	err = dec.Decode(&i)
	if err != nil {
		return nil, -1, err
	}

	if !i.Success {
		return nil, i.Status, errors.New("Imgur Failed with Status: " + strconv.Itoa(i.Status))
	}

	return &i, i.Status, nil
}

// Delete Image from Ingur using Bearer
//	info, status, err := DeleteImageAuthed("abc")
func (client *Client) DeleteImageAuthed(id string) (*ImageInfoWithoutData, int, error) {
	URL := "https://api.imgur.com/3/image/" + id

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	err := writer.Close()
	if err != nil {
		return nil, -1, err
	}

	req, err := http.NewRequest("DELETE", URL, payload)
	if err != nil {
		return nil, -1, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.Imgur.AccessToken))
	req.Header.Set("Content-Type", writer.FormDataContentType())

	res, err := client.HTTPClient.Do(req)
	if err != nil {
		return nil, -1, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, -1, err
	}

	dec := json.NewDecoder(bytes.NewReader(body))

	var i ImageInfoWithoutData

	err = dec.Decode(&i)
	if err != nil {
		return nil, -1, err
	}

	if !i.Success {
		return nil, i.Status, errors.New("Imgur Failed with Status: " + strconv.Itoa(i.Status))
	}

	return &i, i.Status, nil
}
