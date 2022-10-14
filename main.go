package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("from: %s\n", r.RemoteAddr)
		fmt.Printf("path: %s\n", r.URL)
		fmt.Printf("user agent: %s\n", r.UserAgent())
		fmt.Printf("method: %s\n", r.Method)
		fmt.Printf("content type: %s\n", r.Header.Get("Content-Type"))

		for k, v := range r.URL.Query() {
			var x struct{}
			if err := json.Unmarshal([]byte(v[0]), &x); err != nil {
				fmt.Printf("Parameter %s is: %s\n", k, v)
			} else {
				dst := &bytes.Buffer{}
				err := json.Indent(dst, []byte(v[0]), "", "  ")
				if err != nil {
					fmt.Printf("Parameter %s is: %s\n", k, v)
				} else {
					fmt.Printf("Parameter %s is: \n", k)
					fmt.Println(dst.String())
				}
			}
		}

		var y struct{}
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Body cannot be read")
		}
		if err := json.Unmarshal(bodyBytes, &y); err != nil {
			fmt.Println("body is:")
			fmt.Println(string(bodyBytes))
		} else {
			dst := &bytes.Buffer{}
			err := json.Indent(dst, bodyBytes, "", "  ")
			if err != nil {
				fmt.Println("body is:")
				fmt.Println(string(bodyBytes))
			} else {
				fmt.Println("body is:")
				fmt.Println(dst.String())
			}

		}

		fmt.Println("------------")

		w.WriteHeader(http.StatusOK)
	}
}

func main() {
	server := &http.Server{
		Addr:    os.Args[1],
		Handler: handle(),
	}

	fmt.Printf("Listening on %s", server.Addr)
	_ = server.ListenAndServe()
}
