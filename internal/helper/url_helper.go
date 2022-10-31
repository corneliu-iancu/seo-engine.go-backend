package helper

import (
	"fmt"
	"net/url"
	"strings"
)

func GetURIAsSlice(u *url.URL) []string {
	pathParams := []string{u.Host}

	p := strings.Split(u.Path, string('/')) // @todo: validate if the last path param is "/"
	// fmt.Println("[DEBUG][GetURIAsSlice]", len(p[len(p)-1]))
	if len(p[len(p)-1]) == 0 { // empty parameter at the end due to "/" as last character in URL.
		fmt.Println("[DEBUG][GetURIAsSlice] Warning! Input URI ends in a '/'. Removing")
		pathParams = append(pathParams, p[1:len(p)-1]...)
	} else {
		pathParams = append(pathParams, p[1:]...)
	}

	return pathParams
}
