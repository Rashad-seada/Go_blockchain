package network

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Tick struct {
	C chan time.Time
}

func NewTick() *Tick {
	return &Tick{
		C: make(chan time.Time),
	}
}

func (t *Tick) StartTicker(tickDuration int){
	hashSeed := "Rashad"


	for {

		for i := 0; i < tickDuration; i++ {
			hash := sha256.Sum256([]byte(hashSeed))
			hashSeed = hex.EncodeToString(hash[:])
		}

		t.C <- time.Now()

	}
	
	
}
