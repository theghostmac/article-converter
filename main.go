package main

import "flag"

const (
	version = "v0_27122022"
	tempDir = "article-converter"
)

func main() {
	flagVersion := flag.Bool("v", false, version)
	flagLink := flag.String("link", "", "Link to the article for download")
	flagType := flag.String("type", "epub", "Type of file. Available: epub (default), \n mobi")
	flagTitle := flag.String("title", "", "The article title is the default title, but you can change it using this flag")
}
