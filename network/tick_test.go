package network

import (
	"fmt"
	"testing"
	"time"
)

func TestTick(t *testing.T) {
	tick := Tick{
		C: make(chan time.Time),
	}

	go tick.StartTicker(1000000)

	for t := range tick.C {
		fmt.Println("t",t)
	}

}
