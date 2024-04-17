package main

import "C"
import (
	"awesomeProject/modelLoader"
	"fmt"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

//helper function to load string model

func getModelAsString() string {
	data, err := os.ReadFile("model.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	return string(data)
}

// helper function to get rss as bytes
func GetRssMB() string {
	// Read memory statistics from /proc/self/statm
	data, err := os.ReadFile("/proc/self/statm")
	if err != nil {
		fmt.Println("Error reading /proc/self/statm:", err)
		return "0"
	}

	// Extract resident memory size (in pages)
	fields := strings.Fields(string(data))
	if len(fields) < 2 {
		fmt.Println("Unexpected format of /proc/self/statm")
		return "0"
	}
	rssPages, err := strconv.ParseUint(fields[1], 10, 64)
	if err != nil {
		fmt.Println("Error parsing resident memory size:", err)
		return "0"
	}

	// Convert pages to bytes (assuming 4 KB page size)
	rssBytes := rssPages * 4096
	rssBytes /= 1024 * 1024

	return strconv.FormatUint(rssBytes, 10)
}

var modelString string

func init() {
	modelString = getModelAsString()
	debug.FreeOSMemory()
	println()
	time.Sleep(10 * time.Second)
}

func loadLearner() {
	modelLoader.Load(modelString)
	time.Sleep(20 * time.Second)
	debug.FreeOSMemory()
	println("C_API Model Loaded: " + GetRssMB())
}

func freeLearner() {
	modelLoader.ReleaseMemory()
	time.Sleep(20 * time.Second)
	debug.FreeOSMemory()
	println("Model Released: " + GetRssMB())
}

func main() {
	for i := 1; i <= 20; i++ {
		println("Initial: " + GetRssMB())
		loadLearner()
		time.Sleep(20 * time.Second)
		freeLearner()
		time.Sleep(20 * time.Second)
		println()
	}
}
