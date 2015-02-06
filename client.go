package editor
import "fmt"
import "os"
import "os/exec"
import "os/signal"
import "syscall"
import "net/rpc"

type Client struct{
}

func(cl *Client) Start() error{ 
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	var b []byte = make([]byte ,1)
	client, err := rpc.DialHTTP("tcp","127.0.0.1:1234")
	if err != nil{
		fmt.Println("dail err : ", err)
		return err
	}
	go func(){
		var reply int
		for{
			os.Stdin.Read(b)
			go client.Call("EchoServer.Echo", b, &reply)
		}
	}()
	<-c
	err = exec.Command("stty", "-F", "/dev/tty", "sane").Run()
	if err != nil{
		fmt.Println("sane command failed", err)
	}
	return nil
}
