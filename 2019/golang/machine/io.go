package machine

import (
	"aoc/utils"
	"fmt"
	"log"
)

var (
	StdIn  = NewReadOnlyFile(NewUserInputReader())
	StdOut = NewWriteOnlyFile(NewUserOutputWriter())
	StdErr = NewWriteOnlyFile(NewUserErrorWriter())
)

type IO struct {
	stdIn  ReaderWriter
	stdOut ReaderWriter
}

type Reader interface {
	Read() int
}

type Writer interface {
	Write(int)
}

type ReaderWriter interface {
	Read() int
	Write(int)
}

type ReadOnlyFile struct {
	reader Reader
}

type PipeFile struct {
	queue utils.Queue[int]
}

func NewPipeFile(size ...int) PipeFile {
	bufferSize := 1
	if len(size) > 0 {
		bufferSize = size[0]
	}
	return PipeFile{
		queue: utils.NewQueue[int](bufferSize),
	}
}

func (io PipeFile) Read() int {
	return utils.OrDie[int](io.queue.Dequeue())
}

func (io PipeFile) Write(value int) {
	utils.OrDie1[int](io.queue.Enqueue(value))
}

type UserInputReader struct {
	readMsg string
}

func NewUserInputReader(readMsg ...string) UserInputReader {
	var msg string
	if len(readMsg) > 0 {
		msg = readMsg[0]
	}
	return UserInputReader{msg}
}

func (r UserInputReader) Read() int {
	fmt.Print(r.readMsg)
	var userInput int
	utils.OrDie(fmt.Scanf("%d", &userInput))
	return userInput
}

type UserOutputWriter struct {
	prefix string
}

func NewUserOutputWriter(prefix ...string) UserOutputWriter {
	var msg string
	if len(prefix) > 0 {
		msg = prefix[0]
	}
	return UserOutputWriter{msg}
}

func (w UserOutputWriter) Write(value int) {
	fmt.Printf("%s%d\n", w.prefix, value)
}

type UserErrorWriter struct {
	prefix string
}

func NewUserErrorWriter(prefix ...string) UserErrorWriter {
	var msg string
	if len(prefix) > 0 {
		msg = prefix[0]
	}
	return UserErrorWriter{msg}
}

func (e UserErrorWriter) Write(value int) {
	print(fmt.Sprintf("%s%d\n", e.prefix, value))
}

func NewReadOnlyFile(reader Reader) ReadOnlyFile {
	return ReadOnlyFile{reader}
}

func (r ReadOnlyFile) Read() int {
	return r.reader.Read()
}

func (io ReadOnlyFile) Write(value int) {
	log.Fatal("ReadOnlyFile does not support Write")
}

func NewWriteOnlyFile(writer Writer) WriteOnlyFile {
	return WriteOnlyFile{writer}
}

type WriteOnlyFile struct {
	writer Writer
}

func (r WriteOnlyFile) Read() int {
	log.Fatal("WriteOnlyFile does not support Read")
	return 0
}

func (r WriteOnlyFile) Write(value int) {
	r.writer.Write(value)
}
