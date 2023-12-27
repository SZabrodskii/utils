package main

import "fmt"

const bufferMaxSize = 1024

type MaxSizeExceededError struct {
	desiredLen int
}

func (e *MaxSizeExceededError) Error() string {
	return fmt.Sprintf("buffer max size exceeded: %d > %d", e.desiredLen, bufferMaxSize)
}

type EndOfBufferError struct{}

func (eob *EndOfBufferError) Error() string {
	return "end of buffer"
}

type ByteBuffer struct {
	buffer []byte
	offset int
}

type ByteReader interface {
	ReadByte() (byte, error)
}

type ByteWriter interface {
	WriteByte(c byte) error
}

func (b *ByteBuffer) WriteByte(c byte) error {
	if len(b.buffer) >= bufferMaxSize {
		return &MaxSizeExceededError{}
	}
	b.buffer = append(b.buffer, c)

	return nil
}
func (b *ByteBuffer) ReadByte() (byte, error) {
	if b.offset >= len(b.buffer) {
		return 0, &EndOfBufferError{}
	}
	c := b.buffer[b.offset]
	b.offset++
	return c, nil
}
func (b *ByteBuffer) Write(p []byte) (int, error) {
	if len(b.buffer)+len(p) > bufferMaxSize {
		return 0, &MaxSizeExceededError{desiredLen: len(b.buffer) + len(p)}
	}
	b.buffer = append(b.buffer, p...)
	return len(p), nil
}

func (b *ByteBuffer) Read(p []byte) (int, error) {
	if b.offset >= len(b.buffer) {
		return 0, new(EndOfBufferError)
	}

	n := copy(p, b.buffer[b.offset:])
	b.offset += n
	return n, nil
}

func main() {
	buf := &ByteBuffer{}
	err := buf.WriteByte('a')
	if err != nil {
		if _, ok := err.(*MaxSizeExceededError); ok {
			fmt.Println("Maximum size exceeded error occurred")
		}

	}

	for {
		_, err := buf.ReadByte()
		if err != nil {
			if _, ok := err.(*EndOfBufferError); ok {
				fmt.Println("End of buffer error occurred")
				break
			}
		}
	}

}
