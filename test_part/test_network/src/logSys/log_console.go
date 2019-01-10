package logSys

import "fmt"

type log_console struct {

}

func (log *log_console) Debug(a ...interface{}){
	fmt.Println(a)
}