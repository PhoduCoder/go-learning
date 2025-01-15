package main
import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
)

func main() {

	type ExitStatus struct{
		exitCode int
		Message string
	}

    sigs := make(chan os.Signal, 1) //Creates a channel that takes OS signals and is buffered channel with 1 capacity
    //done := make(chan struct{}) //Create another channel that takes empty struct
	done := make(chan ExitStatus)
	//Trap the signal that we are interested in 
	signal.Notify(sigs,syscall.SIGINT)
	
    go func() {
    for {
        s := <-sigs //Receiving the data in the channel
        switch s {
            case syscall.SIGINT:
                fmt.Println()
                fmt.Println("My process has been interrupted. Someone might of pressed CTRL-C")
				fmt.Println("Some clean up is occuring")
				//done <-struct{}{}
                done <- ExitStatus{ exitCode: 100, Message: "Ctrl-C pressed"}
            }
        }
    }()
    fmt.Println("Program is blocked until a signal is caught")
	//done <- struct{}{}
	status := <-done 
    fmt.Printf("Out of here because of exit Code %d\n", status.exitCode)
}