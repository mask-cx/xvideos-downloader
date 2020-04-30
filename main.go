package main

import (
	"flag"
	"fmt"
	"os"

)

var (
	_taskid   string
	_url      string
	output   string
	chanSize int
)

func init() {
	flag.StringVar(&_url, "u", "", "M3U8 url")
	flag.IntVar(&chanSize, "c", 25, "Maximum number of occurrences")
	flag.StringVar(&output, "o", "", "Save files to PREFIX/.. Output to folder")
}

func main() {
	flag.Parse()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("[error]", r)
			os.Exit(-1)
		}
	}()
	if _url == "" {
		fatal("u")
	}
	if output == "" {
		fatal("o")
	}
	if chanSize <= 0 {
		panic(" '-c' must bigger than 0")
	}
	_taskid = RandStringBytes(8)
	downloader, err := NewTask(output, _url)
	if err != nil {
		panic(err)
	}
	if err := downloader.DownloadStart(chanSize); err != nil {
		panic(err)
	}
	fmt.Println("Done!")
}

func fatal(name string) {
	panic(" '-" + name + "' is required")
}
