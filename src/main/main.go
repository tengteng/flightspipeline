package main

import (
	"flag"
	"fmt"
	"runtime"

	"feeder"

	"github.com/golang/glog"
)

func main() {
	defer glog.Flush()

	task := flag.String("task", "feed",
		"Specify a task. Can be feed|")
	cargo_url := flag.String("cargo_url", "cargo.eng.parsec.apple.com:2181", "Cargo Url.")
	provider := flag.String("provider", "", "Service provider.")
	sleepTime := flag.Float64("sleep_time", 0.0, "Sleep time.")

	flag.Parse()
	flag.Lookup("logtostderr").Value.Set("true")

	fmt.Printf("cargo url: [%v]\nprovider: [%v]\nsleeptime: [%vs]\n", *cargo_url, *provider, *sleepTime)

	if *task == "feed" {
		feeder.StartFeeding(*sleepTime, *cargo_url, *provider)
	} else {
		fmt.Println("Unknown feeder")
		return
	}

}
