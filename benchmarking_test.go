package ed448

import (
	"math/big"

	. "gopkg.in/check.v1"
)

func (s *Ed448Suite) BenchmarkBigintsAddition(c *C) {
	curve := newBigintsCurve()
	c.ResetTimer()
	x, y := gx, gy
	for i := 0; i < c.N; i++ {
		rx, ry := curve.add(x, y, x, y)
		x, y = rx.(*big.Int), ry.(*big.Int)
	}
}

func (s *Ed448Suite) BenchmarkRadixAddition(c *C) {
	curve := newRadixCurve()
	c.ResetTimer()
	x, y := gx, gy
	for i := 0; i < c.N; i++ {
		rx, ry := curve.add(x, y, x, y)
		x, y = rx.(*big.Int), ry.(*big.Int)
	}
}

func (s *Ed448Suite) BenchmarkDoubling(c *C) {
	curve := newBigintsCurve()
	c.ResetTimer()
	x, y := gx, gy
	for i := 0; i < c.N; i++ {
		rx, ry := curve.double(x, y)
		x, y = rx.(*big.Int), ry.(*big.Int)
	}
}

func (s *Ed448Suite) BenchmarkMultiplication(c *C) {
	curve := newBigintsCurve()
	c.ResetTimer()
	x, y := gx, gy
	for i := 0; i < c.N; i++ {
		rx, ry := curve.multiply(x, y, []byte{0x03})
		x, y = rx.(*big.Int), ry.(*big.Int)
	}
}

func (s *Ed448Suite) BenchmarkRadixPointIsOnCurve(c *C) {
	gx := mustDeserialize(serialized{
		0x9f, 0x93, 0xed, 0x0a, 0x84, 0xde, 0xf0,
		0xc7, 0xa0, 0x4b, 0x3f, 0x03, 0x70, 0xc1,
		0x96, 0x3d, 0xc6, 0x94, 0x2d, 0x93, 0xf3,
		0xaa, 0x7e, 0x14, 0x96, 0xfa, 0xec, 0x9c,
		0x70, 0xd0, 0x59, 0x3c, 0x5c, 0x06, 0x5f,
		0x24, 0x33, 0xf7, 0xad, 0x26, 0x6a, 0x3a,
		0x45, 0x98, 0x60, 0xf4, 0xaf, 0x4f, 0x1b,
		0xff, 0x92, 0x26, 0xea, 0xa0, 0x7e, 0x29,
	})
	gy := mustDeserialize(serialized{0x13})

	curve := newRadixCurve()

	c.ResetTimer()
	for i := 0; i < c.N; i++ {
		curve.isOnCurve(&Affine{gx, gy})
	}
}

func (s *Ed448Suite) BenchmarkRadixPointDouble(c *C) {
	gx := serialized{
		0x9f, 0x93, 0xed, 0x0a, 0x84, 0xde, 0xf0,
		0xc7, 0xa0, 0x4b, 0x3f, 0x03, 0x70, 0xc1,
		0x96, 0x3d, 0xc6, 0x94, 0x2d, 0x93, 0xf3,
		0xaa, 0x7e, 0x14, 0x96, 0xfa, 0xec, 0x9c,
		0x70, 0xd0, 0x59, 0x3c, 0x5c, 0x06, 0x5f,
		0x24, 0x33, 0xf7, 0xad, 0x26, 0x6a, 0x3a,
		0x45, 0x98, 0x60, 0xf4, 0xaf, 0x4f, 0x1b,
		0xff, 0x92, 0x26, 0xea, 0xa0, 0x7e, 0x29,
	}
	gy := serialized{0x13}

	p, err := NewPoint(gx, gy)
	c.Assert(err, IsNil)

	c.ResetTimer()
	for i := 0; i < c.N; i++ {
		p = p.Double()
	}

	curve := newRadixCurve()
	c.Assert(curve.isOnCurve(p), Equals, true)
}

func (s *Ed448Suite) BenchmarkRadixPointUnifiedAddition(c *C) {
	gx := serialized{
		0x9f, 0x93, 0xed, 0x0a, 0x84, 0xde, 0xf0,
		0xc7, 0xa0, 0x4b, 0x3f, 0x03, 0x70, 0xc1,
		0x96, 0x3d, 0xc6, 0x94, 0x2d, 0x93, 0xf3,
		0xaa, 0x7e, 0x14, 0x96, 0xfa, 0xec, 0x9c,
		0x70, 0xd0, 0x59, 0x3c, 0x5c, 0x06, 0x5f,
		0x24, 0x33, 0xf7, 0xad, 0x26, 0x6a, 0x3a,
		0x45, 0x98, 0x60, 0xf4, 0xaf, 0x4f, 0x1b,
		0xff, 0x92, 0x26, 0xea, 0xa0, 0x7e, 0x29,
	}
	gy := serialized{0x13}

	basePoint, err := NewPoint(gx, gy)
	c.Assert(err, IsNil)

	c.ResetTimer()
	p := basePoint
	for i := 0; i < c.N; i++ {
		p = p.Add(p)
	}

	curve := newRadixCurve()
	c.Assert(curve.isOnCurve(p), Equals, true)
}
