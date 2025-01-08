package main

import (
	"fmt"
	"github.com/asseco-voice/go-docx"
	"io"
	"os"
)

func Fumiama() {
	d := docx.NewA4()
	p := d.AddParagraph()
	p.AddLink("link", "b")

	f, err := os.Create("testdata/test_out.docx")
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = d.WriteTo(f)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func Open() {
	file, err := os.Open("testdata/123.docx")
	if err != nil {
		return
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		return
	}
	size := fileInfo.Size()
	reader := io.ReaderAt(file)
	doc, err := docx.Parse(reader, size)
	fmt.Println(doc)
}

func main() {
	Open()
}
