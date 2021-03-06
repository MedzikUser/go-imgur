package test

import (
	"net/http"
	"os"
	"testing"
)

var imgDeleteHashs []string
var imgID string

func TestUploadImageFromURL(t *testing.T) {
	key := os.Getenv("IMGUR_CLIENT_ID")

	client := createClient(new(http.Client), key, "")

	i, _, err := client.UploadImageFromURL("https://golang.org/doc/gopher/fiveyears.jpg", "")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	imgDeleteHashs = append(imgDeleteHashs, i.Data.Deletehash)
}

func TestUploadImageFromURL2(t *testing.T) {
	key := os.Getenv("IMGUR_CLIENT_ID")

	client := createClient(new(http.Client), key, "")

	i, _, err := client.UploadImageFromURL("https://golang.org/doc/gopher/fiveyears.jpg", "")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	imgID = i.Data.ID
}

func TestUploadImageFromFile(t *testing.T) {
	key := os.Getenv("IMGUR_CLIENT_ID")

	client := createClient(new(http.Client), key, "")

	i, _, err := client.UploadImageFromFile("test_data/fiveyears.jpg", "")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	imgDeleteHashs = append(imgDeleteHashs, i.Data.Deletehash)
}

func TestGetImageInfo(t *testing.T) {
	key := os.Getenv("IMGUR_CLIENT_ID")

	client := createClient(new(http.Client), key, "")

	_, _, err := client.GetImageInfo(imgID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestDeleteImagesAuthed(t *testing.T) {
	client, err := generateAccessToken()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	_, _, err = client.DeleteImageAuthed(imgID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestDeleteImagesUnAuthed(t *testing.T) {
	key := os.Getenv("IMGUR_CLIENT_ID")

	client := createClient(new(http.Client), key, "")

	for _, v := range imgDeleteHashs {
		_, _, err := client.DeleteImageUnAuthed(v)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
	}
}
