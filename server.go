package editor
import "net/rpc"
import "fmt"
import "net"
import  "net/http"
import "bytes"
type EchoServer struct{
}

func(e *EchoServer) Echo(args *EchoArgs, reply *EchoReply) error{
	//backpsace 
	if bytes.Equal(*args, []byte{ 127}) {
		fmt.Print("\b\033[K")
		return nil
	}
	fmt.Print(string(*args))
	return nil
}

func StartServer() error{
	echo := new(EchoServer)
	rpc.Register(echo)
	rpc.HandleHTTP()
	l,e := net.Listen("tcp", ":1234")
	if e != nil{
		fmt.Println("listen error : ", e)
	}
	http.Serve(l, nil)
	return nil
}
