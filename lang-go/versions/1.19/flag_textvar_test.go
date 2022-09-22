package go119_test

import (
	"flag"
	"math/big"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFlagTextvar(t *testing.T) {
	var (
		addr     net.IP
		intVar   big.Int
		floatVar big.Float
		timeVar  time.Time
	)

	fs := flag.NewFlagSet("command", flag.PanicOnError)
	fs.TextVar(&addr, "addr", net.IPv4(127, 0, 0, 1), "set addr")
	fs.TextVar(&intVar, "int", big.NewInt(0), "set big int")
	fs.TextVar(&floatVar, "float", big.NewFloat(0), "set big float")
	fs.TextVar(&timeVar, "time", time.Now(), "set big time")

	fs.Parse([]string{
		"-addr", "8.8.8.8",
		"-int", "1267650600228229401496703205376",
		"-float", "3.141592653589793",
		"-time", "2022-08-02T12:34:56.000000Z",
	})

	assert.Equal(t, net.IPv4(8, 8, 8, 8), addr)
	assert.Equal(t, new(big.Int).Exp(big.NewInt(2), big.NewInt(100), nil), &intVar)
	assert.Equal(t, new(big.Float).SetFloat64(3.141592653589793).String(), floatVar.String())
	assert.Equal(t, time.Date(2022, 8, 2, 12, 34, 56, 0, time.UTC), timeVar)
}
