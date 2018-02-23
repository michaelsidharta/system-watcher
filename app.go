package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

func main() {
	// startTime := time.Now()
	// var currentTime time.Time
	//var lastCPUPercentage, currentCPUPercentage float64
	for {
		//currentTime = time.Now()
		cpuPercentage, _ := cpu.Percent(100*time.Millisecond, false)
		currentCPUPercentage := cpuPercentage[0]
		fmt.Println(currentCPUPercentage)
	}
}
