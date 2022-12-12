// main.go
// go-cors is a fast & lightweight golang reverse proxy used to bypass those annoying cors errors
/* go-cors is limited to only GET requests with no headers  */
package main

import (
	"net/http"
	"io"
	"net/url"
)

//vars
var PORT = ":4001"

func main() {
	//main route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "url")

		//grab header value
    	requestURL := r.Header.Get("url")

		//validate url
		_, err := url.ParseRequestURI(requestURL)
		if err == nil {
			//send request & return results
			res , err := http.Get(requestURL)
			if(err == nil) {
				body, _ := io.ReadAll(res.Body)
				w.Write(body) //return back result
			} else { 
				//request error
				w.Write([]byte("error sending request"))
			} 
		} else {
			w.Write([]byte("invalid url"))
		}
	})
	http.ListenAndServe(PORT, nil)
}