package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"net/http"

	"github.com/nats-io/graft"
)

func main() {
	fmt.Println("hello kek")

	ci := graft.ClusterInfo{Name: "health_manager", Size: 3}

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

	if node.State() == graft.LEADER {
		// Process as a LEADER
	}

	select {
	case sc := <-stateChangeChan:
		// Process a state change
		fmt.Println(sc.To.String())
	case err := <-errChan:
		// Process an error, log etc.
		fmt.Println(err)
	}

	http.HandleFunc("/", hello)
	http.ListenAndServe(":8090", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.UserAgent())
	w.WriteHeader(200)
	fmt.Fprintf(w, "hello kek")
}
