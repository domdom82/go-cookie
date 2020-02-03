package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {

		cookie := &http.Cookie{
			Name:     "JSESSIONID",
			Value:    "abc123",
			Secure:   true,
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		}

		http.SetCookie(w, cookie)
		w.Write([]byte(fmt.Sprintf("set cookie: \n%s\n", cookie)))
	})

	http.ListenAndServe(":8080", mux)

}
