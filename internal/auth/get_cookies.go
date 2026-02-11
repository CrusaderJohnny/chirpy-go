package auth

import (
	"errors"
	"net/http"
)

func GetCookies(w http.ResponseWriter, r *http.Request) ([]*http.Cookie, error) {
	cookiesJar := r.Cookies()
	if len(cookiesJar) == 0 {
		return []*http.Cookie{}, errors.New("no cookies found")
	}
	return cookiesJar, nil
}
