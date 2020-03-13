package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const lastWord = "Go!"
const countDownStart = 3

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep(){
	c.sleep(c.duration)
}

//Countdown count from 3 to 1
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countDownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, lastWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1* time.Second,time.Sleep}
	Countdown(os.Stdout,sleeper)
}
