package main

import (
	"fmt"
	"log"
	"net/http"
)
/*When the web application is served directly to the internet, we
can get the client IP address like this*/
//func getClientIpAddr(req *http.Request) string {
// 	return req.RemoteAddr
//}

/*
In reality, most web applications are served behind proxies and load balancers;
we can detect client IP addresses by retrieving the value from
X-Forwarded-For headers. Combining with fallback to the default
client IP address from request, we have this little function that
does the job
*/
func getClientIpAddr(req *http.Request) string {
	clientIp := req.Header.Get("X-FORWARDED-FOR")
	if clientIp != "" {
		return clientIp
	}
	return req.RemoteAddr
}

func main() {
	// Define your handler function
	handler := func(w http.ResponseWriter, r *http.Request) {
		clientIp := getClientIpAddr(r)
		// Write response
		fmt.Fprintf(w, "Hello, %s @%s!", r.URL.Path[1:], clientIp)
	}

	// Register handler function for route "/"
	http.HandleFunc("/", handler)

	// Start server
	log.Fatal(http.ListenAndServe(":5000", nil))
}
