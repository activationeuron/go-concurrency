package example

import (
	"fmt"
	"time"
)

func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Microsecond * 1)
	}
}

func main() {
	// this call is in main routine
	fun("direct call")

	// we dont get any output because the main go runtine finished its executations

	go fun("go rutine 1")

}
