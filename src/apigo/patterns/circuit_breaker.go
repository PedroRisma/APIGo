package patterns

import (
	"github.com/sparrc/go-ping"
	"time"
)

type CircuitBreaker struct {
	state			string
	timeout			time.Time
	allowErrors		int
	errors			int
	chState 		chan string
	chTimeout 		chan time.Time

}

func IniCircuit(cb CircuitBreaker, errors int) {

	cb.state = "CLOSE"
	cb.allowErrors = errors
	cb.timeout = time.Now()
	cb.chState <- cb.state
	cb.chTimeout <- cb.timeout

	go func(cs chan string, ct chan time.Time) {

		for {
			select {
				case st := <- cs:
					if st == "HALF-OPEN"{
						TestConnection(cb)
					}else {
						if st == "OPEN"{
							cb.chTimeout <- <- time.After(time.Millisecond * 5)
						}else {
							cb.errors = 0
						}
					}
				case <- ct:
					cb.state = "HALF-OPEN"
					cb.chState <- cb.state
			}
		}

	}(cb.chState, cb.chTimeout)

}

func ConnectionEnabled(cb CircuitBreaker) bool {

	if cb.state == "CLOSE" {
		return true
	}else {
		return false
	}

}

func InformError (cb CircuitBreaker) {

	cb.errors++
	if cb.errors - cb.allowErrors == 0 {
		cb.state = "OPEN"
		cb.chState <- cb.state
	}

}

func TestConnection(cb CircuitBreaker, urls ... string) {

	var c [len(urls)] chan bool
	i := 0
	for _, in := range urls {
		go func(url string){
			for {
				c[i] <- <- testResponse(url)
			}
		}(in)
	}

	for k := 0; k< i; k++ {
		if !<- c[k] {
			cb.state = "OPEN"
			cb.chState <- cb.state
			return
		}
	}

	cb.state = "CLOSE"
	cb.chState <- cb.state
}

func testResponse(url string) (chan bool){

	response := make(chan bool)
	response <- true
	pinger, err := ping.NewPinger(url)
	if err != nil {
		panic(err)
	}
	pinger.Count = 1
	pinger.Run()
	pinger.OnFinish = func(statistics *ping.Statistics) {
		if  statistics.PacketsSent - statistics.PacketsRecv != 0{
			response <- false
		}
	}

	return response
}
