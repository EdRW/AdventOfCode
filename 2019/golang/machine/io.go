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
	StdIn  Reader
	StdOut Writer
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

type Pipe struct {
	queue utils.Queue[int]
}

func NewPipe(size ...int) Pipe {
	bufferSize := 1
	if len(size) > 0 {
		bufferSize = size[0]
	}
	return Pipe{
		queue: utils.NewQueue[int](bufferSize),
	}
}

func (io Pipe) Read() int {
	return utils.OrDie[int](io.queue.Dequeue())
}

func (io Pipe) Write(value int) {
	utils.OrDie1[int](io.queue.Enqueue(value))
}

func (io Pipe) Len() int {
	return len(io.queue)
}

func (io Pipe) Flush() []int {
	out := make([]int, 0)
	if io.Len() > 0 {
		out = append(out, io.Read())
	}
	return out
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

// A ReaderWriter that panics when Write() func is called
type ReadOnlyFile struct {
	reader Reader
}

// Adds a Write() func to a Reader that panics when called
func NewReadOnlyFile(reader Reader) ReadOnlyFile {
	return ReadOnlyFile{reader}
}

func (r ReadOnlyFile) Read() int {
	return r.reader.Read()
}

func (io ReadOnlyFile) Write(value int) {
	log.Fatal("ReadOnlyFile does not support Write")
}

// A ReaderWriter that panics when Read() func is called
type WriteOnlyFile struct {
	writer Writer
}

// Adds a Read() func to a Writer that panics when called
func NewWriteOnlyFile(writer Writer) WriteOnlyFile {
	return WriteOnlyFile{writer}
}

func (r WriteOnlyFile) Read() int {
	log.Fatal("WriteOnlyFile does not support Read")
	return 0
}

func (r WriteOnlyFile) Write(value int) {
	r.writer.Write(value)
}

// writes to the multiple writers when `Write(value)` is called
type MultiWriter struct {
	writers []Writer
}

// Creates a MultiWriter that writes to the multiple
// writers provided when `Write(value)` is called
func NewMultiWriter(writers ...Writer) MultiWriter {
	return MultiWriter{writers}
}

func (m MultiWriter) Write(value int) {
	for _, writer := range m.writers {
		writer.Write(value)
	}
}
