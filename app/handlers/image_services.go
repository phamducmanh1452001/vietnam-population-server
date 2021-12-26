package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"vietnam-population-server/app/router"
)

var magicTable = map[string]string{
	"\xff\xd8\xff":      "image/jpeg",
	"\x89PNG\r\n\x1a\n": "image/png",
	"GIF87a":            "image/gif",
	"GIF89a":            "image/gif",
}

func mimeFromIncipit(incipit []byte) string {
	incipitStr := string(incipit)
	for magic, mime := range magicTable {
		if strings.HasPrefix(incipitStr, magic) {
			return mime
		}
	}
	return ""
}

func UploadImage(w *router.ResponseWriter, r *http.Request) {
	header := w.Writer().Header()
	header.Add("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	r.ParseMultipartForm(10 << 16)
	file, handler, err := r.FormFile("image")

	defer file.Close()

	if err != nil {
		respondError(w, http.StatusForbidden, "Error Retrieving the File: "+err.Error())
		return
	}

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	if err != nil {
		fmt.Println(err)
	}

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	mime := mimeFromIncipit(fileBytes)

	if mime == "" {
		respondError(w, http.StatusNotFound, "Error extension of image file")
		return
	}

	tempFile, _ := ioutil.TempFile("images", "avatar*.png")
	avatarName := tempFile.Name()
	defer tempFile.Close()

	tempFile.Write(fileBytes)

	respondJSON(w, 200, map[string]string{
		"name": avatarName,
		"mime": mime,
	})

}

func DownloadImage(w *router.ResponseWriter, r *http.Request) {
	imageName, err := getParam(r, "name")
	if err != nil {
		respondError(w, http.StatusBadRequest, "URL Param is missing")
		return
	}
	fileBytes, err := ioutil.ReadFile("images/" + imageName)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	header := w.Writer().Header()
	header.Add("Access-Control-Allow-Origin", "*")
	header.Set("Content-Type", "image/jpeg")
	w.WriteHeader(http.StatusOK)
	w.Write(fileBytes)
}
