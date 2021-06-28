package main

import (
	"flag"
	"fmt"
	"github.com/nats-io/nats.go"
	"net/http"
	"time"

	"github.com/nats-io/graft"
)

func main() {
	var port string
	flag.StringVar(&port, "port", "8080", "set up port")
	flag.Parse()

	fmt.Println("hello kek")

	ci := graft.ClusterInfo{Name: "ys5zd84St3FA6AIwZYKtVx", Size: 3}

	opts := nats.GetDefaultOptions()
	rpc, err := graft.NewNatsRpc(&opts)
	if err != nil {
		fmt.Println(err)
		return
	}

	errChan := make(chan error)
	stateChangeChan := make(chan graft.StateChange)
	handler := graft.NewChanHandler(stateChangeChan, errChan)

	node, err := graft.New(ci, handler, rpc, "/tmp/graft.log")
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		for{
		fmt.Println(node.State())
		time.Sleep(time.Second * 5)
		}
	}()

	if node.State() == graft.LEADER {
		// Process as a LEADER
		fmt.Println("leader elected")
	}

	go func() {
		select {
		case sc := <-stateChangeChan:
			// Process a state change
			fmt.Println("change state to - " + sc.To.String())
		case err := <-errChan:
			// Process an error, log etc.
			fmt.Println(err)
		}
	}()

	http.HandleFunc("/", hello)
	http.ListenAndServe(":"+port, nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.UserAgent())
	w.WriteHeader(200)
	fmt.Fprintf(w, "hello kek")
}
