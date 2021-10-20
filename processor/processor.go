package processor

import (
	"fmt"
	"github.com/nats-io/graft"
	"github.com/nats-io/nats.go"
	"time"
)

func Job(){
	t := time.NewTicker(time.Second)
	for{
		<-t.C
		fmt.Print("Do some work  ")
	}
}

type Processor struct{
	job func()

	node            *graft.Node
	stateChangeChan chan graft.StateChange
	errChan         chan error
}

func NewProcessor(nodeName, url string, replicas int) (*Processor, error){
	proc := &Processor{job: Job}
	err := proc.initNode(nodeName, url, replicas)
	if err != nil {
		return nil,err
	}

	go proc.Start()
	go proc.listenStatusChange()
	go proc.stateReport()

	return proc, nil
}


func (p *Processor) initNode(nodeName, url string,replicas int) error{
	clusterInfo := graft.ClusterInfo{Name: nodeName, Size: replicas}
	opts := nats.GetDefaultOptions()
	opts.Url = url
	rpc, err := graft.NewNatsRpc(&opts)
	if err != nil {
		return err
	}

	p.errChan = make(chan error)
	p.stateChangeChan = make(chan graft.StateChange)
	handler := graft.NewChanHandler(p.stateChangeChan, p.errChan)
	p.node, err = graft.New(clusterInfo, handler, rpc, "/tmp/graft.log")
	if err != nil {
		return err
	}

	return nil
}

func (p *Processor) Start() {
	time.Sleep(time.Second * 5)

	switch p.node.State(){
	case graft.LEADER:
		p.job()
	case graft.CANDIDATE, graft.FOLLOWER:
		// do nothing
	}
}

func (p *Processor) ShutDown() {
	if p.node != nil {
		fmt.Println("shutdown node")
		p.node.Close()
	}
}

func (p *Processor) stateReport() {
	for{
		time.Sleep(time.Second * 5)
		fmt.Println("\n[State report] = ", p.node.State())
	}
}

func (p *Processor) listenStatusChange() {
	for {
		select {
		case sc := <-p.stateChangeChan:
			fmt.Println("\nchange state from - " + sc.From.String() + " to -" + sc.To.String())
			if sc.To == graft.LEADER{
				p.Start()
			}
		case err := <-p.errChan:
			fmt.Println(err)
		}
	}
}
