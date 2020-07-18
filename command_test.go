package nsq

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

func BenchmarkCommand(b *testing.B) {
	//b.StopTimer()
	//data := make([]byte, 2048)
	//cmd := Publish("test", data)
	//var buf bytes.Buffer
	//b.StartTimer()
	//
	//for i := 0; i < b.N; i++ {
	//	cmd.WriteTo(&buf)
	//}

	for i := 0; i < b.N; i++ {
		s := strconv.Itoa(rand.Int())
		fmt.Println(s)
	}

}
