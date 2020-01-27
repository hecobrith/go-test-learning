package main
// this package is in charge of dealing with most of the servers and stuff on the internet
import "net/http"
import "log"
import "encoding/json"

func main()  {
	// the http handle function is a server and recive three parameters, first is the path
	// then is the callback o the handle
	// functions are first class citizens so we can pass a function declaration just inside the callback without declare
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		// these will map all the query parameters
		names := r.URL.Query()["name"]
		var name string
		if len(names) == 1 {
			name = names[0]
		}
		// to create a json we first need to create a map of string to string
		m := map[string]string{"name":name}
		// the  we need to encode this and we pass a writer interface to the inizialisation of these so we can encode
		enc := json.NewEncoder(w)
		enc.Encode(m)
		// we are writeing into a writer interface and writeers always  work with array of bytes
		// so in input out put you will wlays be working with a collection of bytes (SLICES)
		w.Write([]byte("Hello" + name))
	})

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}
}