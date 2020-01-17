package ed448

import (
	"math/big"
	"sync"
)

// CurveParams contains the parameters of an elliptic curve and also provides
// a generic, non-constant time implementation of Curve.
// This are the Montgomery params
type CurveParams struct {
	P       *big.Int // the order of the underlying field
	N       *big.Int // the order of the base point
	A       *big.Int // the constant of the curve equation
	Gu, Gv  *big.Int // (x,y) of the base point
	BitSize int      // the size of the underlying field
	Name    string   // the canonical name of the curve
}

// A GoldilocksCurve represents Goldilocks curve or edwards448.
// See https://www.hyperelliptic.org/EFD/g1p/auto-shortw.html
type GoldilocksCurve interface {
	// Params returns the parameters for the curve.
	Params() *CurveParams
	// Params returns the edwards parameters for the curve.
	EdwardsParams() *CurveParams
	// IsOnCurve reports whether the given (x,y) lies on the curve.
	IsOnCurve(x, y *big.Int) bool
	// Add returns the sum of (x1,y1) and (x2,y2)
	Add(x1, y1, x2, y2 *big.Int) (x, y *big.Int)
	// Double returns 2*(x,y)
	Double(x1, y1 *big.Int) (x, y *big.Int)
	// ScalarMult returns k*(Bx,By) where k is a number in big-endian form.
	ScalarMult(x1, y1 *big.Int, k []byte) (x, y *big.Int)
	// ScalarBaseMult returns k*G, where G is the base point of the group
	// and k is an integer in big-endian form.
	ScalarBaseMult(k []byte) (x, y *big.Int)
}

// Params returns the parameters for the curve.
func (curve *CurveParams) Params() *CurveParams {
	return curve
}

var initonce sync.Once
var curve448 *CurveParams
var ed448 *CurveParams

func initAll() {
	initCurve448()
	//inited448()
}

func initCurve448() {
	// See https://safecurves.cr.yp.to/field.html and https://tools.ietf.org/html/rfc7748#section-4.2
	curve448 = &CurveParams{Name: "curve-448"}
	curve448.P, _ = new(big.Int).SetString("726838724295606890549323807888004534353641360687318060281490199180612328166730772686396383698676545930088884461843637361053498018365439", 10)
	curve448.N, _ = new(big.Int).SetString("181709681073901722637330951972001133588410340171829515070372549795146003961539585716195755291692375963310293709091662304773755859649779", 10)
	curve448.A, _ = new(big.Int).SetString("156326", 10)
	curve448.Gu, _ = new(big.Int).SetString("5", 10)
	curve448.Gv, _ = new(big.Int).SetString("3617de4a96262c6f5d9e98bf9292dc29f8f41dbd289a147ce9da3113b5f0b8c00a60b1ce1d7e819d7a431d7c90ea0e5f", 10)
	curve448.BitSize = 448
}