package labrpc

import (
	"sync"
	"strconv"
	"testing"
	"runtime"
)

type JunkArgs struct{
	X int
}

type JunkReply struct{
	X string
}

type JunkServer struct{
	mu sync.Mutex
	log1 []string
	log2 []int
}

func (js *JunkServer) Handler1(args string,reply *int){
	js.mu.Lock()
	defer js.mu.Unlock()
	js.log1=append(js.log1,args)
	*reply,_=strconv.Atoi(args)
}

func TestBasic(t *testing.T)  {
	runtime.GOMAXPROCS(4)

	rn:=MakeNetwork()

	e:=rn.MakeEnd("end1-99")

	js:=&JunkServer{}
	svc:=MakeService(js)

	rs:=MakeServer()

	rs.AddService(svc)
	rn.AddServer("server99",rs)

	rn.Connect("end1-99","server99")
	rn.Enable("end1-99",true)

	{
		reply:=0
		e.Call("JunkServer.Handler1","9099",&reply)
		if reply!=9099{
			t.Fatalf("wrong reply from Handler1")
		}
	}
}
