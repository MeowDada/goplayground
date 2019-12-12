.PHONY: all clean

GO_EXTENSION:=.go

READ_FILE_PROG:=read-file
READ_FILE_SRCS:=read-file

CHAINING_READER_PROG:=chaining-reader
CHAINING_READER_SRCS:=chaining-reader

all: read-file chaining-reader

read-file:
	go build -o $(READ_FILE_PROG) $(addsuffix $(GO_EXTENSION), $(READ_FILE_SRCS))

chaining-reader:
	go build -o $(CHAINING_READER_PROG) $(addsuffix $(GO_EXTENSION), $(CHAINING_READER_SRCS))

clean:
	rm $(READ_FILE_PROG)
	rm $(CHAINING_READER_PROG)