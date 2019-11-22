package main

import (
	"time"
	"fmt"
	"log"
	"os"
	"io"

	minio "github.com/minio/minio-go/v6"
	humanize "github.com/dustin/go-humanize"
)

type ProgressReader struct {
	rawReader io.Reader
	total     int64
	readBytes int64
}

func NewProgressReader(filename string) *ProgressReader {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	fi, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &ProgressReader{
		rawReader: f,
		total:     fi.Size(),
		readBytes: 0,
	}
}

func (p *ProgressReader) Read(buf []byte) (int, error) {
	nr, err := p.rawReader.Read(buf)
	p.readBytes += int64(nr)
	return nr ,err
}

func (p *ProgressReader) Progress() {
	progress := p.readBytes * 100 / p.total
	fmt.Printf("\r%d%% (%s/%s)", progress, humanize.Bytes(uint64(p.readBytes)), humanize.Bytes(uint64(p.total)))
}

func main() {

	var (
		address = "192.168.80.119:9000"
		accessKey = "minio"
		secretKey = "minio123"
		ssl = false
	)

	clnt, err := minio.New(address, accessKey, secretKey, ssl)
	if err != nil {
		log.Fatal(err)
	}

	var (
		bucket = "test"
		object = "something.tmp"
		filepath = "/home/jack/Desktop/dummy_collections/90/dummy_90_500MB.tmp"
	)

	pr := NewProgressReader(filepath)

	opts := minio.PutObjectOptions{
		Progress: pr,
	}

	go func() {
		for {
			time.Sleep(time.Duration(200)*time.Millisecond)
			pr.Progress()
		}
	}()

	clnt.FPutObject(bucket, object, filepath, opts)
}