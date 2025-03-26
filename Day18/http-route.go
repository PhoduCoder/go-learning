package main

import (
	"log"
	"net/http"
)

type hello struct{}

// Implement the serveHTTP interface
// ResponseWriter is to write any thing to response
// http.Request helps to understand the request
func (h hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := "<h1>Hello World</h1>"
	w.Write([]byte(msg))

}

func main() {
	http.HandleFunc("/more", func(w http.ResponseWriter, r *http.Request) {
		moreText := `<table>
					<caption>
					Front-end web developer course 2025
				</caption>
				<thead>
					<tr>
					<th scope="col">Person</th>
					<th scope="col">Most interest in</th>
					<th scope="col">Age</th>
					</tr>
				</thead>
				<tbody>
					<tr>
					<th scope="row">Chris</th>
					<td>HTML tables</td>
					<td>22</td>
					</tr>
					<tr>
					<th scope="row">Dennis</th>
					<td>Web accessibility</td>
					<td>45</td>
					</tr>
					<tr>
					<th scope="row">Sarah</th>
					<td>JavaScript frameworks</td>
					<td>29</td>
					</tr>
					<tr>
					<th scope="row">Karen</th>
					<td>Web performance</td>
					<td>36</td>
					</tr>
				</tbody>
				<tfoot>
					<tr>
					<th scope="row" colspan="2">Average age</th>
					<td>33</td>
					</tr>
				</tfoot>
				</table>`

		w.Write([]byte(moreText))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
