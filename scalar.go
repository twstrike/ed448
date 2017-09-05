package ed448

import "errors"

// Scalar is a interface of a Ed448 scalar
type Scalar interface {
	Equals(a Scalar) bool
	Copy() Scalar
	Add(a, b Scalar)
	Sub(a, b Scalar)
	Mul(a, b Scalar)
	Halve(a Scalar)
	Encode() []byte
	BarretDecode(src []byte) error
}

type scalar [scalarWords]word

func (s *scalar) montgomeryMultiply(x, y *scalar) {
	out := &scalar{}
	carry := word(0x00)

	for i := 0; i < scalarWords; i++ {
		chain := dword(0x00)
		for j := 0; j < scalarWords; j++ {
			chain += dword(x[i])*dword(y[j]) + dword(out[j])
			out[j] = word(chain)
			chain >>= wordBits
		}

		saved := word(chain)
		multiplicand := out[0] * montgomeryFactor
		chain = dword(0x00)

		for j := 0; j < scalarWords; j++ {
			chain += dword(multiplicand)*dword(ScalarQ[j]) + dword(out[j])
			if j > 0 {
				out[j-1] = word(chain)
			}
			chain >>= wordBits
		}
		chain += dword(saved) + dword(carry)
		out[scalarWords-1] = word(chain)
		carry = word(chain >> wordBits)
	}

	out.subExtra(out, ScalarQ, carry)
	copy(s[:], out[:])
}

func (s *scalar) equals(x *scalar) bool {
	diff := word(0x00)
	for i := uintZero; i < scalarWords; i++ {
		diff |= s[i] ^ x[i]
	}
	return word(((dword(diff))-1)>>wordBits) == decafTrue
}

func (s *scalar) copy() *scalar {
	out := &scalar{}
	copy(out[:], s[:])
	return out
}

func (s *scalar) set(w word) {
	s[0] = w
}

func (s *scalar) subExtra(minuend *scalar, subtrahend *scalar, carry word) {
	out := &scalar{}
	var chain sdword

	for i := uintZero; i < scalarWords; i++ {
		chain += sdword(minuend[i]) - sdword(subtrahend[i])
		out[i] = word(chain)
		chain >>= wordBits
	}

	borrow := chain + sdword(carry)
	chain = 0

	for i := uintZero; i < scalarWords; i++ {
		chain += sdword(out[i]) + (sdword(ScalarQ[i]) & borrow)
		out[i] = word(chain)
		chain >>= wordBits
	}
	copy(s[:], out[:])
}

func (s *scalar) add(a, b *scalar) {
	out := &scalar{}
	var chain dword

	for i := uintZero; i < scalarWords; i++ {
		chain += dword(a[i]) + dword(b[i])
		out[i] = word(chain)
		chain >>= wordBits
	}
	out.subExtra(out, ScalarQ, word(chain))
	copy(s[:], out[:])
}

func (s *scalar) sub(x, y *scalar) {
	noExtra := word(0x00)
	s.subExtra(x, y, noExtra)
}

func (s *scalar) mul(x, y *scalar) {
	s.montgomeryMultiply(x, y)
	s.montgomeryMultiply(s, scalarR2)
}

func (s *scalar) halve(a *scalar) {
	mask := -(a[0] & 1)
	var chain dword
	var i uint

	for i = 0; i < scalarWords; i++ {
		chain += dword(a[i]) + dword(ScalarQ[i]&mask)
		s[i] = word(chain)
		chain >>= wordBits
	}
	for i = 0; i < scalarWords-1; i++ {
		s[i] = s[i]>>1 | s[i+1]<<(wordBits-1)
	}

	s[i] = s[i]>>1 | word(chain<<(wordBits-1))
}

// Serializes an array of words into an array of bytes (little-endian)
func (s *scalar) serialize(dst []byte) error {
	wordBytes := wordBits / 8
	if len(dst) < fieldBytes {
		return errors.New("dst length smaller than fieldBytes")
	}

	for i := 0; i*wordBytes < fieldBytes; i++ {
		for j := 0; j < wordBytes; j++ {
			b := s[i] >> uint(8*j)
			dst[wordBytes*i+j] = byte(b)
		}
	}
	return nil
}

func (s *scalar) decodeShort(b []byte, size uint) {
	k := uint(0)
	for i := uint(0); i < scalarLimbs; i++ {
		out := word(0)
		for j := uint(0); j < 4 && k < size; j, k = j+1, k+1 {
			out |= (word(b[k])) << (8 * j)
		}
		s[i] = out
	}
}

func (s *scalar) decode(b []byte) word {
	s.decodeShort(b, scalarBytes)

	accum := sdword(0x00)
	for i := 0; i < 14; i++ {
		accum += sdword(s[i]) - sdword(ScalarQ[i])
		accum >>= wordBits
	}

	s.mul(s, &scalar{0x01})

	return word(accum)
}

// HACKY: either the param or the return
func decodeLong(s *scalar, b []byte) *scalar {
	y := &scalar{}
	bLen := len(b)
	size := bLen - (bLen % fieldBytes)

	if bLen == 0 {
		s = scalarZero.copy()
		return s
	}

	if size == bLen {
		size -= fieldBytes
	}
	s.decodeShort(b[size:], uint(bLen-size))

	if bLen == scalarBytes {
		s.mul(s, &scalar{0x01})
		return s
	}

	for size == bLen-(bLen%fieldBytes) {
		size -= scalarBytes
		s.montgomeryMultiply(s, scalarR2)
		y.decode(b[size:])
		s.add(s, y)
	}

	return s.copy()
}

//Exported methods

// NewScalar returns a Scalar in Ed448 with decaf
func NewScalar(in ...[]byte) Scalar {
	if len(in) > 1 {
		panic("too many arguments to function call")
	}

	if in == nil {
		return &scalar{}
	}

	out := &scalar{}

	bytes := in[0][:]
	return decodeLong(out, bytes)
}

// Equals compares two scalars. Returns true if they are the same; false, otherwise.
func (s *scalar) Equals(x Scalar) bool {
	return s.equals(x.(*scalar))
}

// Copy copies scalars.
func (s *scalar) Copy() Scalar {
	out := &scalar{}
	copy(out[:], s[:])
	return out
}

// Add adds two scalars. The scalars may use the same memory.
func (s *scalar) Add(x, y Scalar) {
	s.add(x.(*scalar), y.(*scalar))
}

// Sub subtracts two scalars. The scalars may use the same memory.
func (s *scalar) Sub(x, y Scalar) {
	noExtra := word(0)
	s.subExtra(x.(*scalar), y.(*scalar), noExtra)
}

// Mul multiplies two scalars. The scalars may use the same memory.
func (s *scalar) Mul(x, y Scalar) {
	s.montgomeryMultiply(x.(*scalar), y.(*scalar))
	s.montgomeryMultiply(s, scalarR2)
}

// Halve halfs a scalar. The scalars may used the same memory.
func (s *scalar) Halve(x Scalar) {
	s.halve(x.(*scalar))
}

// Encode serializes a scalar to wire format.
func (s *scalar) Encode() []byte {
	dst := make([]byte, fieldBytes)
	s.serialize(dst)
	return dst
}

// Decode reads a scalar from wire format or from bytes and reduces mod scalar prime.
// XXX: this will reduce with barret, change name and receiver
func (s *scalar) BarretDecode(src []byte) error {
	if len(src) < fieldBytes {
		return errors.New("ed448: cannot decode a scalar from a byte array with a length unequal to 56")
	}
	barrettDeserializeAndReduce(s[:], src, &curvePrimeOrder)
	return nil
}

// Decode reads a scalar from wire format or from bytes and reduces mod scalar prime.
// XXX: make scalar part of signature as it will generate confusion otherwise
func Decode(x Scalar, src []byte) Scalar {
	return decodeLong(x.(*scalar), src)
}
