package main

import (
	"flag"
	"log"
	"net/http"
)

// Define a string constant containing the HTML for the webpage. This consists of a <h1>
// header tag, and some JavaScript which calls our POST /v1/tokens/authentication
// endpoint and writes the response body to inside the <div id="output"></div> tag.
const html = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
</head>
<body>
    <section class="contact-form">
      <h1>Preflight CORS</h1>
      <p>Use this handy contact form to get in touch with me.</p>
      
      <form>
        <div class="input-group">
          <label for="email">Email Address</label>
          <input id="email" name="email" type="email"/>
        </div>
        
        <div class="input-group">
          <label for="password">Password</label>
          <input id="password" name="password" type="text"/>
        </div>
        
        <button type="submit">Send It!</button>
      </form>
    </section>

    <div class="results">
      <h2>Form Data</h2>
      <pre></pre>
    </div>
    <script>
      function handleFormSubmit(event) {
      event.preventDefault();

      const data = new FormData(event.target);

      const formJSON = Object.fromEntries(data.entries());

      const results = document.querySelector(".results pre");
      results.innerText = JSON.stringify(formJSON, null, 2);
      
      fetch("http://localhost:4000/v1/tokens/authentication", {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify(formJSON, null, 2)
      }).then(
        function (response) {
          response.text().then(function (text) {
            document.querySelector(".results pre").innerText = text;
          });
        },
        function (err) {
          document.querySelector(".results pre").innerText = err;
        }
      );
    }

    const form = document.querySelector(".contact-form");
    form.addEventListener("submit", handleFormSubmit);

    </script>
</body>
</html>`

func main() {
	addr := flag.String("addr", ":9000", "Server address")
	flag.Parse()

	log.Printf("starting server on %s", *addr)

	err := http.ListenAndServe(*addr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(html))
	}))
	log.Fatal(err)
}
