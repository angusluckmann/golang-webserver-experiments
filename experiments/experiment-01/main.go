/* ---------- FILE DESCRIPTION ----------
   Projectname: experiment-01
   Purpose: A first example of the
            webserver capabilites of
            the golang standard
            library.
   Note: If the program execution
         fails, change the webserver
         port.
   Author: Angus Luckmann
   License: Apache License 2.0
-------------------------------------- */

package main

import (
	"fmt"
	"net/http"
)

// Event Handling Function ("index.html")
func handle_events_index_html(par_writer http.ResponseWriter, par_request *http.Request) {
	fmt.Fprintf(par_writer, "<h1>index.html</h1><p>Some text</p>")
}

// Event Handling Function ("test.html")
func handle_events_test_html(par_writer http.ResponseWriter, par_request *http.Request) {
	fmt.Fprintf(par_writer, "<h1>test.html</h1><p>Some text</p>")
}

// Main Function
func main() {
	// Prepare for webserver start
	fmt.Println("Preparing the webserver start process ...")
	http.HandleFunc("/index.html", handle_events_index_html)
	http.HandleFunc("/test.html", handle_events_test_html)

	// Start the webserver
	fmt.Println("Starting the webserver ...")
	http.ListenAndServe(":8080", nil) // 8080 = Port of webserver, nil = ???
}
