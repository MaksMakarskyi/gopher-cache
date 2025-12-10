package benchmark

import (
	"bufio"
	"net"
	"testing"
)

func BenchmarkLlen(b *testing.B) {
	conn, err := net.Dial("tcp", "localhost:6379")
	if err != nil {
		b.Fatal(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)

	cmd := "*6\r\n$5\r\nLPUSH\r\n$14\r\nllen_bench_key\r\n$3\r\nfoo\r\n$3\r\nbar\r\n$4\r\nfizz\r\n$4\r\nbazz\r\n"
	if _, err := conn.Write([]byte(cmd)); err != nil {
		b.Fatal(err)
	}
	if _, _, err := reader.ReadLine(); err != nil {
		b.Fatal(err)
	}

	cmd = "*2\r\n$4\r\nLLEN\r\n$14\r\nllen_bench_key\r\n"

	for b.Loop() {
		if _, err := conn.Write([]byte(cmd)); err != nil {
			b.Fatal(err)
		}

		if _, _, err := reader.ReadLine(); err != nil {
			b.Fatal(err)
		}
	}
}
