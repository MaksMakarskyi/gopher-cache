package benchmark

import (
	"bufio"
	"net"
	"testing"
)

func BenchmarkSadd(b *testing.B) {
	conn, err := net.Dial("tcp", "localhost:6379")
	if err != nil {
		b.Fatal(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)

	cmd := "*3\r\n$4\r\nSADD\r\n$14\r\nsadd_bench_key\r\n$3\r\nbar\r\n"

	for b.Loop() {
		if _, err := conn.Write([]byte(cmd)); err != nil {
			b.Fatal(err)
		}

		if _, _, err := reader.ReadLine(); err != nil {
			b.Fatal(err)
		}
	}
}
