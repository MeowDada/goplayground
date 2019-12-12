package main

import (
	"fmt"
	"io"
	"os/signal"
	"syscall"
	"log"
	"os"
	"context"
	
	minio "github.com/minio/minio-go/v6"
	humanize "github.com/dustin/go-humanize"
)

type ProgressReader struct {
	rawReader  io.Reader
	readBytes  int64
	totalBytes int64
}

func NewProgressreader(reader io.Reader, totalBytes int64) ProgressReader{
	return ProgressReader{
		rawReader: reader,
		readBytes: 0,
		totalBytes: totalBytes,
	}
}

func (p *ProgressReader) Read(buf []byte) (int, error) {
	n, err := p.rawReader.Read(buf)
	p.readBytes += int64(n)
	fmt.Printf("\r%s/%s", humanize.Bytes(uint64(p.readBytes)), humanize.Bytes(uint64(p.totalBytes)))
	return n, err
}

func main() {

	var (
		address = "192.168.80.119:9000"
		accessKey = "minio"
		secretKey = "minio123"
		useSSL = false
		bucketName = "test"
		objectName = "bigbig.ctx"
		path = "/home/jack/Desktop/bigbig"
	)

	minioClient, err := minio.New(address, accessKey, secretKey, useSSL)
	if err != nil {
		log.Fatal(err)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	fi, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
		return
	}

	sigCh := make(chan os.Signal, 1)

	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	go func(){
		<-sigCh
		cancel()
	}()

	pr := NewProgressreader(f, fi.Size())

	_, err = minioClient.PutObjectWithContext(ctx, bucketName, objectName, &pr, fi.Size(), minio.PutObjectOptions{

	})
	if err != nil {
		fmt.Println(ctx.Err())
		fmt.Println(err)
	}
}