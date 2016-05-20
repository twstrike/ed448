package ed448

import (
	"bytes"
	"flag"
	"io"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type Ed448Suite struct{}

var _ = Suite(&Ed448Suite{})
var slow = flag.Bool("slow", false, "Include slow tests.")

func (s *Ed448Suite) TestMarshalAndUnmarshal(c *C) {
	curve := newBigintsCurve()
	x, y := gx, gy

	marshaled := marshal(curve, x, y)

	uncompressedForm := uint8(4)
	c.Assert(marshaled[0], Equals, uncompressedForm)

	ux, uy := unmarshal(curve, marshaled)

	c.Assert(x, DeepEquals, ux)
	c.Assert(y, DeepEquals, uy)
}

func (s *Ed448Suite) TestKeyGenerationWithBigintsArithmetic(c *C) {
	if !*slow {
		c.Skip("Slow test. Add -slow flag to run it.")
	}

	goldilocks := NewBigintsGoldilocks()

	priv, pub, err := goldilocks.GenerateKey(getReader())

	c.Assert(err, Equals, nil)
	px, py := unmarshal(newBigintsCurve(), pub)
	c.Assert(priv[1] > 0, Equals, true)
	c.Assert(bisCurve.isOnCurve(px, py), Equals, true)
}

func getReader() io.Reader {
	len := (rho.BitLen() + 7)
	zeroes := make([]byte, len)
	return bytes.NewReader(zeroes)
}
