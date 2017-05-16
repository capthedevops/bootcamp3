package main
import (
	"fmt"
	"net/http"
	"os"
	"github.com/newrelic/go-agent"
	"log"
)

func handler(w http.ResponseWriter, r *http.Request) {
	h, _ := os.Hostname()
	fmt.Fprintf(w, "Hi there, I'm served from %s!", h)
}

func main() {
	config := newrelic.NewConfig("Bootcamp3", "8435c85f5d3212c0d85a3059a0928a34b55ca3b5")
	app, err := newrelic.NewApplication(config)

	http.HandleFunc(newrelic.WrapHandleFunc(app, "/users", handler))

	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}

}
