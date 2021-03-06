package routine

import(
	"testing"
	"fmt"
    "sync"
    "reflect"
    "go.uber.org/goleak"
)

// TestChannelSuccess will create a success test for channel function
func TestChannel(t *testing.T){
	var wg sync.WaitGroup
	var tests = []struct {
        want int
    }{
		{ want : 0 },
    }

    for _, tt := range tests {
        testname := fmt.Sprintf("Test1")
        t.Run(testname, func(t *testing.T) {
            ans := Channel(&wg)
            if ans != tt.want {
                fmt.Println("got", reflect.TypeOf(ans), "want", reflect.TypeOf(tt.want))
                t.Errorf("got %d want %d", ans, tt.want)
            }
        })
    }
}

func BenchmarkChannel(b *testing.B){
    var wg sync.WaitGroup
    for i := 0; i < b.N; i++ {
        Channel(&wg)
    }
}

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}