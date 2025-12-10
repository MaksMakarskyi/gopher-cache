package benchmark

import (
	"bufio"
	"net"
	"testing"
)

func BenchmarkGet(b *testing.B) {
	conn, err := net.Dial("tcp", "localhost:6379")
	if err != nil {
		b.Fatal(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)

	cmd := "*3\r\n$3\r\nSET\r\n$13\r\nget_bench_key\r\n$3\r\nbar\r\n"
	if _, err := conn.Write([]byte(cmd)); err != nil {
		b.Fatal(err)
	}
	if _, err := reader.ReadString('\n'); err != nil {
		b.Fatal(err)
	}

	cmd = "*2\r\n$3\r\nGET\r\n$13\r\nget_bench_key\r\n"

	for b.Loop() {
		if _, err := conn.Write([]byte(cmd)); err != nil {
			b.Fatal(err)
		}

		if _, _, err := reader.ReadLine(); err != nil {
			b.Fatal(err)
		}
		if _, _, err := reader.ReadLine(); err != nil {
			b.Fatal(err)
		}
	}
}
