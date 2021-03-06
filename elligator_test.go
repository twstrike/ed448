package ed448

import (
	. "gopkg.in/check.v1"
)

func (s *Ed448Suite) Test_PointFromNonUniformHash(c *C) {
	ser := [56]byte{
		0x0d, 0xc3, 0x50, 0x83, 0xbd, 0x0d, 0x55, 0xd3,
		0xdd, 0x8c, 0x50, 0x28, 0x1c, 0x32, 0x77, 0x7e,
		0x62, 0x03, 0x82, 0xa3, 0x18, 0x72, 0xc1, 0x7b,
		0xc2, 0x69, 0x11, 0xb7, 0x34, 0xbd, 0x2b, 0xc4,
		0x26, 0x3b, 0x82, 0xc2, 0x12, 0x3c, 0x2a, 0x3c,
		0xe7, 0x49, 0x3c, 0x21, 0xfb, 0xb2, 0xb5, 0xe2,
		0x41, 0xa8, 0xce, 0xa6, 0xb7, 0xf1, 0xc1, 0x3f,
	}

	exp := &twExtendedPoint{
		&bigNumber{
			0x333bf6, 0x348f59f, 0xfa8f51, 0x1a45b6f,
			0xc96ff10, 0xd7f98a3, 0xd0c0b22, 0x6a3b696,
			0x832df79, 0x2824845, 0xd78421b, 0x748a71,
			0x5d67a95, 0x566135d, 0xd4ac41a, 0x3ace179,
		},
		&bigNumber{
			0x4650aac, 0xb984b48, 0x4b3deb4, 0x1f7c40,
			0x5766713, 0x938ae1c, 0x14e14aa, 0x23088a8,
			0x2a3d97a, 0xfc23e90, 0xd9b4f23, 0x9a3bef4,
			0xdf20d8e, 0xcd9cc1d, 0xb76a0d7, 0x48ad86d,
		},
		&bigNumber{
			0xcb1854c, 0xf29cb25, 0x3159292, 0x2101752,
			0xbc095c5, 0x13ccc0a, 0x7058e53, 0x7b77b37,
			0x5316c77, 0xaa930cd, 0x2a179e8, 0x4a4fca6,
			0xe91a389, 0x1593de7, 0xff691f7, 0x623b4c1,
		},

		&bigNumber{
			0x2470f89, 0x72c68a3, 0xa7701ab, 0x41b2e5a,
			0x5958957, 0x22ca902, 0xf5b72a5, 0x7ddc067,
			0xe724bec, 0x67350a2, 0x842d3e1, 0x413c1d8,
			0xdb8a266, 0xf2472a6, 0xb630473, 0x8468933,
		},
	}
	p := pointFromNonUniformHash(ser)
	c.Assert(p.x, DeepEquals, exp.x)
	c.Assert(p.y, DeepEquals, exp.y)
	c.Assert(p.z, DeepEquals, exp.z)
	c.Assert(p.t, DeepEquals, exp.t)
}

var elligatorTestVectorsInput = [][]byte{
	{0x2d, 0x86, 0xa1, 0x42, 0x33, 0x8d, 0xe2, 0x74, 0x80, 0x63, 0x54, 0xc4, 0x3e, 0x29,
		0xaf, 0x70, 0x5a, 0xa9, 0xa1, 0x89, 0x3e, 0x6f, 0xd3, 0xee, 0x2e, 0x95, 0x22, 0xc9,
		0xce, 0xb4, 0x0b, 0xe2, 0x44, 0x1b, 0xac, 0x8a, 0x4f, 0x78, 0x06, 0x43, 0x43, 0x89,
		0x25, 0xd7, 0x91, 0x46, 0x98, 0x8b, 0x1c, 0xa1, 0x12, 0xda, 0x71, 0x4d, 0xe9, 0x2a},
	{0xee, 0x79, 0x8e, 0xe0, 0x86, 0xde, 0x1f, 0x5a, 0x57, 0xa2, 0xca, 0x28, 0xdb, 0x84,
		0x51, 0xd3, 0x06, 0xcb, 0xb9, 0xee, 0x22, 0x27, 0xc4, 0x97, 0xf4, 0xa6, 0x7a, 0x69,
		0x06, 0xd7, 0xeb, 0xbc, 0x7a, 0xa8, 0x5f, 0x94, 0x6f, 0xf9, 0xdf, 0xf7, 0x9e, 0x1b,
		0x7e, 0x88, 0xd9, 0x7e, 0x3a, 0xd4, 0xa4, 0xe0, 0xa1, 0x20, 0x32, 0x32, 0x3a, 0xb7},
	{0x2e, 0x1b, 0x10, 0x93, 0xb3, 0x47, 0x75, 0x97, 0x66, 0x46, 0x49, 0xb0, 0xb7, 0xc6,
		0xac, 0x1f, 0x9b, 0xb7, 0x5d, 0xd9, 0xfd, 0xb5, 0x08, 0x96, 0xcb, 0xaa, 0x06, 0x15,
		0xc5, 0x25, 0x43, 0x6d, 0x62, 0x54, 0xec, 0x13, 0xd9, 0x19, 0x0e, 0xa4, 0x25, 0xe5,
		0xba, 0x80, 0xee, 0xfc, 0x25, 0x9b, 0xcd, 0x1e, 0x2a, 0x5a, 0xf0, 0x0e, 0x8a, 0x9e},
	{0x8a, 0x59, 0x3f, 0xb9, 0x9c, 0x04, 0xb5, 0xc0, 0x50, 0xc9, 0x0d, 0xc7, 0x90, 0x93,
		0x65, 0x89, 0x41, 0x5b, 0x9b, 0xd6, 0x78, 0x3d, 0x9e, 0x92, 0x5e, 0x63, 0x4b, 0x87,
		0x81, 0x4f, 0xd1, 0xda, 0x2a, 0x36, 0xcd, 0x80, 0x45, 0xbb, 0x6c, 0x36, 0xd6, 0x7e,
		0xb8, 0x2c, 0x17, 0x84, 0x01, 0x35, 0x6b, 0xe8, 0x40, 0x42, 0x9c, 0x78, 0x0c, 0x70},
	{0x80, 0xb2, 0x5f, 0xfc, 0xfb, 0xd8, 0x0b, 0x83, 0xa7, 0x86, 0xc1, 0x07, 0x6d, 0x3a,
		0x23, 0xd8, 0x50, 0x49, 0xfd, 0x4c, 0x51, 0x91, 0x92, 0xa7, 0xd1, 0xe8, 0x52, 0x38,
		0x93, 0x6e, 0x1c, 0x09, 0x22, 0x15, 0xc8, 0x0b, 0x2d, 0x9d, 0xd1, 0x3d, 0x88, 0x49,
		0x82, 0x9d, 0x7f, 0x6a, 0x38, 0x2a, 0x5a, 0xce, 0x05, 0x16, 0x6e, 0x4b, 0x08, 0x5b},
	{0xc1, 0x15, 0x77, 0x32, 0xc6, 0xd2, 0xba, 0xf4, 0x48, 0x88, 0x7a, 0x1c, 0x4a, 0x2a,
		0x90, 0xb4, 0x0b, 0x07, 0x84, 0x0f, 0xf9, 0x62, 0xda, 0x1f, 0x71, 0x91, 0x05, 0x8c,
		0xb9, 0x37, 0xdf, 0xe5, 0xce, 0xb2, 0x5e, 0x34, 0x4e, 0x33, 0xfc, 0x9d, 0xf0, 0xc6,
		0x8e, 0x99, 0xcb, 0x35, 0x07, 0xaa, 0xfe, 0xb9, 0xa6, 0xc9, 0x66, 0x75, 0xbb, 0xf1},
	{0xa5, 0x50, 0x98, 0x77, 0xa2, 0xbb, 0xe8, 0x0d, 0x07, 0xc2, 0x3b, 0x26, 0x46, 0x73,
		0x85, 0xf9, 0x7c, 0x16, 0xbe, 0x48, 0x82, 0x40, 0x0f, 0x31, 0x80, 0x0e, 0x15, 0xdd,
		0x43, 0x9e, 0x52, 0x34, 0x43, 0xcf, 0x94, 0x68, 0x88, 0x59, 0xb7, 0x62, 0x64, 0x3d,
		0x64, 0xbe, 0xda, 0x91, 0xf7, 0x50, 0xac, 0x6e, 0x00, 0x16, 0xaf, 0xaf, 0xd3, 0x09},
	{0xbd, 0x9b, 0xe4, 0xe9, 0x20, 0x93, 0xcf, 0x24, 0x40, 0x79, 0xa6, 0xff, 0x63, 0xad,
		0x01, 0xe1, 0x9c, 0xae, 0x6d, 0x80, 0x65, 0xed, 0x83, 0xbb, 0x05, 0x2e, 0x14, 0xe2,
		0x39, 0x04, 0x8e, 0x3b, 0x8a, 0xeb, 0x90, 0xe9, 0x35, 0xbe, 0xbe, 0x29, 0x24, 0x1e,
		0x34, 0x4d, 0xc9, 0x0d, 0x31, 0xd0, 0x4e, 0x99, 0xd6, 0xa1, 0xad, 0xca, 0x8b, 0x38},
	{0xc3, 0x5e, 0xfb, 0xe1, 0xab, 0xee, 0x01, 0xf9, 0xe4, 0x5e, 0x03, 0x84, 0xfa, 0x2f,
		0x94, 0x3a, 0x6e, 0x8f, 0x56, 0x11, 0x86, 0x4b, 0x55, 0x5f, 0x18, 0x6c, 0x7c, 0xf8,
		0xe3, 0x4c, 0xc6, 0x27, 0xcb, 0xa5, 0x85, 0xfb, 0xcf, 0xc4, 0x26, 0x84, 0xeb, 0x30,
		0xbe, 0x62, 0x23, 0x5c, 0x1e, 0x10, 0xe8, 0x82, 0xca, 0x42, 0x19, 0xa8, 0xc4, 0x85},
	{0x3c, 0x30, 0x0f, 0xed, 0xd9, 0x86, 0x6f, 0x6a, 0xfa, 0xbc, 0x14, 0x3e, 0x1f, 0x73,
		0x0a, 0xf6, 0xea, 0xda, 0xc0, 0x20, 0x7e, 0x00, 0x88, 0x88, 0xb6, 0xeb, 0x79, 0xa2,
		0xf7, 0xe6, 0xe6, 0x7e, 0xd0, 0x1e, 0x71, 0xaf, 0x64, 0x77, 0x7b, 0x90, 0xbf, 0x61,
		0x0a, 0x5e, 0x36, 0xca, 0xd0, 0xcd, 0x88, 0xef, 0x88, 0x3a, 0x9b, 0x6a, 0xb8, 0x13},
	{0x11, 0xf8, 0x2f, 0x21, 0xe4, 0x61, 0x64, 0x36, 0xe6, 0x9e, 0xd8, 0xe3, 0x57, 0x03,
		0xcc, 0xcd, 0x1f, 0x65, 0xaa, 0x75, 0xf0, 0x7e, 0x8a, 0xfa, 0xa3, 0x35, 0x29, 0xcc,
		0x22, 0x58, 0xeb, 0x2b, 0x0f, 0xb1, 0x82, 0x71, 0x0f, 0xfc, 0x67, 0xd1, 0xe0, 0xd0,
		0xde, 0x37, 0x3d, 0x4f, 0xd2, 0xd5, 0xb1, 0x7b, 0x58, 0xb3, 0xc7, 0xd4, 0x73, 0x12},
	{0x3d, 0xbd, 0xcf, 0x91, 0xe8, 0x35, 0xa8, 0x30, 0xfd, 0x8a, 0xf9, 0xc6, 0x9d, 0xc1,
		0x30, 0x66, 0xdf, 0x1e, 0x24, 0x44, 0x8b, 0x91, 0x78, 0xa0, 0x99, 0xbb, 0x07, 0x57,
		0x3e, 0xfe, 0xc4, 0x8e, 0xab, 0x2c, 0x11, 0x9b, 0xcb, 0xbb, 0x82, 0x8d, 0x20, 0xc1,
		0x64, 0x7d, 0x42, 0x31, 0xdf, 0xeb, 0x9b, 0xd0, 0x86, 0xf2, 0x6d, 0xb7, 0x7e, 0x71},
	{0xac, 0x8b, 0xf3, 0x02, 0x0a, 0x1c, 0x73, 0x3a, 0x59, 0x10, 0x92, 0xb6, 0x7a, 0x32,
		0x23, 0xca, 0x2f, 0xab, 0x64, 0x53, 0xd2, 0x25, 0xba, 0x83, 0x2e, 0x34, 0xd0, 0xc4,
		0xbf, 0xca, 0x95, 0x2b, 0xe3, 0x2d, 0x39, 0x76, 0xca, 0x73, 0x8c, 0x5a, 0xb3, 0xdd,
		0xc9, 0xc7, 0x62, 0x70, 0x78, 0x41, 0x83, 0x72, 0xdb, 0x77, 0x0f, 0x17, 0xb5, 0x5c},
	{0xf6, 0xc5, 0x5d, 0x6b, 0x46, 0x97, 0xd6, 0xf8, 0x3d, 0x6e, 0xcc, 0xc4, 0xdb, 0x2f,
		0x72, 0xf8, 0xf2, 0xf6, 0x7e, 0x75, 0x24, 0xff, 0x91, 0xd6, 0xf6, 0xc8, 0xa7, 0x56,
		0xab, 0x03, 0x96, 0x7a, 0x64, 0x89, 0x42, 0x71, 0xe7, 0x1e, 0x71, 0xd8, 0x95, 0x72,
		0x9f, 0x06, 0x31, 0xfd, 0x7c, 0x0d, 0xe1, 0xc2, 0x73, 0xc0, 0x90, 0x92, 0x43, 0x23},
	{0x9b, 0x30, 0x03, 0x76, 0xa4, 0xb9, 0x5e, 0xa2, 0x02, 0x4b, 0xdb, 0xd9, 0x7a, 0x96,
		0x93, 0xc3, 0xf6, 0x0a, 0xe0, 0xbb, 0xdb, 0xda, 0xfc, 0x47, 0x09, 0x27, 0x8b, 0x65,
		0x34, 0xc8, 0xa2, 0xd5, 0xff, 0x9b, 0xb3, 0xd2, 0x10, 0xd9, 0x49, 0xd5, 0xbf, 0x09,
		0x49, 0x19, 0xb9, 0x0d, 0x2f, 0x0f, 0xf9, 0x82, 0xed, 0x92, 0x79, 0x95, 0xdc, 0x60},
	{0x90, 0x95, 0x7d, 0x59, 0x78, 0x10, 0xb2, 0x7b, 0x84, 0x9a, 0x69, 0x1f, 0x5d, 0x27,
		0xd5, 0x48, 0x96, 0x3d, 0x35, 0x4a, 0xe9, 0xe2, 0x9a, 0xd5, 0x9a, 0x23, 0x0a, 0x15,
		0x5a, 0xaa, 0x6f, 0xe7, 0xc5, 0x4c, 0x82, 0xd5, 0x08, 0x14, 0xd8, 0xfd, 0xcd, 0x2d,
		0x3b, 0xb1, 0xe5, 0x53, 0xa8, 0x41, 0xf9, 0x71, 0xd7, 0x24, 0xa4, 0x64, 0x7a, 0xba},
}

var elligatorTestVectorsOutput = [][]byte{
	{0xa6, 0x99, 0x3b, 0x5a, 0x6c, 0xbb, 0x40, 0x71, 0x6e, 0xb2, 0xaf, 0xa1, 0x53, 0x05,
		0x27, 0x75, 0xd2, 0x55, 0xff, 0x2f, 0x64, 0x4e, 0x2f, 0x91, 0x32, 0xb4, 0x04, 0xfc,
		0x80, 0x68, 0x08, 0x09, 0x40, 0x43, 0xf7, 0xa2, 0xe4, 0x7c, 0x0a, 0xd9, 0x27, 0x2f,
		0x53, 0x33, 0x2d, 0x21, 0xf4, 0x07, 0x70, 0xd6, 0x60, 0xa8, 0xf1, 0xf1, 0xed, 0x23},
	{0xde, 0x6a, 0x92, 0x82, 0xee, 0x9f, 0x8f, 0xa9, 0xb0, 0x2c, 0xa9, 0x5e, 0xd4, 0xbf,
		0x7f, 0x87, 0xb7, 0x1f, 0xc3, 0x64, 0xbc, 0x75, 0xd5, 0x71, 0xf2, 0xe9, 0xa7, 0x07,
		0xf7, 0x16, 0x66, 0xb2, 0xdf, 0x06, 0x55, 0xf2, 0x00, 0x2e, 0x1c, 0x84, 0x23, 0x9e,
		0xed, 0x70, 0xde, 0xd8, 0xa6, 0x92, 0xaf, 0x39, 0x52, 0x03, 0x38, 0xc7, 0xc9, 0xef},
	{0x02, 0x51, 0x0b, 0x4c, 0x16, 0xa7, 0x01, 0xa1, 0x68, 0x82, 0xb5, 0x1e, 0xc5, 0xd1,
		0x4e, 0x25, 0x18, 0x5b, 0x7a, 0x8c, 0xd3, 0x12, 0xc3, 0xcf, 0xc0, 0x7c, 0x11, 0x00,
		0x40, 0xd0, 0x01, 0xad, 0x59, 0x0a, 0xd7, 0x2d, 0xc3, 0x07, 0x74, 0xd8, 0x2b, 0x1a,
		0x91, 0xb9, 0xe3, 0x6c, 0x42, 0x3e, 0x93, 0x7d, 0x26, 0x4b, 0x2d, 0x99, 0xd6, 0xb6},
	{0x9c, 0x64, 0x7b, 0x77, 0x1c, 0x28, 0x82, 0x64, 0xe8, 0x0f, 0xc8, 0x11, 0x4c, 0x58,
		0xdb, 0x46, 0xe8, 0xf0, 0x66, 0x6c, 0x10, 0xd7, 0xf5, 0x6b, 0xa8, 0x56, 0xae, 0x67,
		0x09, 0x2a, 0xa8, 0x8c, 0x42, 0x16, 0x65, 0x2e, 0x6a, 0x12, 0x9c, 0x1b, 0x40, 0x90,
		0xca, 0xab, 0xe3, 0x9a, 0xfd, 0x35, 0x2b, 0xe4, 0xdc, 0x40, 0x99, 0x81, 0x9c, 0x59},
	{0x06, 0xe9, 0x16, 0x29, 0xce, 0x93, 0x48, 0x6a, 0xd3, 0xa7, 0xe7, 0x29, 0xf0, 0x1c,
		0x4d, 0x29, 0x4a, 0x4b, 0xde, 0xef, 0xaf, 0x48, 0x32, 0x04, 0xc1, 0x67, 0xdf, 0xe8,
		0xf0, 0xc9, 0xd2, 0x32, 0x50, 0x6f, 0xa5, 0x21, 0xf5, 0x30, 0x0e, 0x19, 0xa0, 0x00,
		0x43, 0x24, 0x50, 0x8b, 0x39, 0x0a, 0x6f, 0x25, 0x81, 0x4f, 0xc8, 0x68, 0x3a, 0xa4},
	{0x68, 0x11, 0x77, 0xb0, 0x76, 0xc9, 0xe5, 0x53, 0xc7, 0xe5, 0x7a, 0x22, 0xe7, 0x59,
		0x05, 0x96, 0xe3, 0x48, 0x2d, 0xe2, 0x3f, 0x28, 0x55, 0xa8, 0xaf, 0x82, 0xcc, 0x51,
		0x6c, 0x52, 0xa9, 0x37, 0x35, 0xed, 0x3d, 0xde, 0x91, 0xb8, 0x21, 0x0b, 0xad, 0x64,
		0xb1, 0x7d, 0x0c, 0x1d, 0x7c, 0x14, 0xcc, 0xc1, 0x52, 0x6c, 0xc4, 0x14, 0x0f, 0x11},
	{0x68, 0x05, 0x63, 0x1c, 0x06, 0xf6, 0xd0, 0xb5, 0xcc, 0xf7, 0x1f, 0xea, 0x2e, 0x4c,
		0xdf, 0x3e, 0xa3, 0x10, 0x4a, 0x44, 0xa8, 0x21, 0x20, 0x5a, 0x25, 0x01, 0x4c, 0x9a,
		0x17, 0xac, 0x43, 0x33, 0xbb, 0xf6, 0xbb, 0x28, 0x9b, 0x42, 0x57, 0xcc, 0xd7, 0xf7,
		0xbb, 0x11, 0xe5, 0xc4, 0xdd, 0xd8, 0x6d, 0xa9, 0x53, 0x19, 0xdc, 0x47, 0x04, 0x4d},
	{0x4c, 0x0e, 0x89, 0x30, 0xee, 0x39, 0xf2, 0xa7, 0x43, 0xd1, 0x79, 0x74, 0x5b, 0x4c,
		0x94, 0x0f, 0xf5, 0x8f, 0x53, 0x99, 0x57, 0x32, 0x31, 0x3d, 0x7e, 0xe7, 0x8c, 0xa2,
		0xde, 0xca, 0x42, 0xa4, 0x8f, 0x00, 0x40, 0xc7, 0x9a, 0x7e, 0xd5, 0x47, 0x00, 0x0b,
		0x20, 0x8b, 0x95, 0x94, 0xce, 0xc4, 0xe3, 0xe9, 0xdf, 0x5c, 0x01, 0x38, 0xb8, 0xaa},
	{0x48, 0xc3, 0x3a, 0x47, 0x66, 0x05, 0xfe, 0x0f, 0xbb, 0x33, 0xd3, 0x7b, 0x67, 0x2a,
		0xac, 0x14, 0xd7, 0xc6, 0x2b, 0x84, 0x56, 0xd2, 0x77, 0x60, 0x8f, 0xc2, 0x90, 0x6d,
		0x03, 0x87, 0x1d, 0x39, 0x59, 0xdd, 0x4a, 0x4c, 0xaf, 0xab, 0xe7, 0xc2, 0x5b, 0x6f,
		0x59, 0xc9, 0xa9, 0xd1, 0x7c, 0x72, 0x4d, 0x97, 0x55, 0x52, 0x98, 0xc9, 0xdf, 0x3f},
	{0x0a, 0x0c, 0x08, 0x9d, 0x50, 0x5d, 0x30, 0xd1, 0xce, 0x91, 0xcf, 0x36, 0x96, 0xca,
		0x76, 0x10, 0xa4, 0xe5, 0x4a, 0xf6, 0xf6, 0x05, 0xcd, 0x68, 0xff, 0x30, 0x3c, 0xb5,
		0x0b, 0xbd, 0xba, 0xb9, 0x90, 0x36, 0x51, 0xed, 0x6b, 0xdc, 0x35, 0xf2, 0xa8, 0x0b,
		0xc7, 0x64, 0xe3, 0x50, 0xf8, 0xa2, 0x3f, 0x70, 0x03, 0xdc, 0xd3, 0xaa, 0x36, 0x4f},
	{0x56, 0x21, 0x3f, 0x80, 0x39, 0x79, 0xce, 0x00, 0x33, 0xa2, 0xaa, 0x9b, 0xcc, 0xb8,
		0x51, 0x3b, 0x82, 0x0b, 0x15, 0x52, 0xe8, 0x14, 0x75, 0x86, 0x4a, 0x48, 0xfe, 0x60,
		0xe9, 0x22, 0x73, 0xa8, 0xf2, 0xe5, 0x7a, 0x77, 0xb8, 0x1a, 0xf1, 0x74, 0x6e, 0x42,
		0xe6, 0x47, 0xcc, 0xc6, 0xfa, 0x54, 0xe0, 0xd0, 0x7c, 0xdd, 0x33, 0x76, 0xc2, 0x39},
	{0xf4, 0x8f, 0xa8, 0x82, 0xb5, 0x2f, 0x79, 0xf1, 0x8f, 0x33, 0xac, 0xfc, 0x23, 0x71,
		0x5e, 0x8f, 0x3e, 0x6c, 0xcf, 0x8e, 0xa8, 0x7a, 0x3f, 0xc0, 0x71, 0xcd, 0xb1, 0xeb,
		0xd2, 0x96, 0xf2, 0x9e, 0x83, 0x15, 0x78, 0xa9, 0x21, 0x29, 0x1d, 0x3c, 0x80, 0x13,
		0x52, 0x59, 0x45, 0x96, 0xa1, 0x7d, 0x27, 0x68, 0xe2, 0xc2, 0x86, 0x32, 0x13, 0x7d},
	{0xaa, 0x3b, 0x6c, 0x33, 0xc2, 0x7a, 0x5a, 0x25, 0xf9, 0x45, 0x20, 0x30, 0x56, 0x73,
		0x32, 0xe1, 0x70, 0x5b, 0xdf, 0x72, 0x45, 0xef, 0xd8, 0x98, 0x60, 0x2c, 0xcf, 0x79,
		0x93, 0x4c, 0xa7, 0x40, 0xed, 0x8a, 0x12, 0xc7, 0xee, 0x82, 0x1e, 0x99, 0x22, 0x52,
		0x1a, 0xb8, 0xbf, 0xca, 0x3a, 0x1d, 0xb9, 0x16, 0xe4, 0x66, 0x78, 0xc5, 0x1f, 0x81},
	{0xba, 0x1c, 0xfd, 0xca, 0x84, 0x4f, 0x16, 0x71, 0x6a, 0x77, 0xba, 0x74, 0x7a, 0x1f,
		0x46, 0xd2, 0x9f, 0xfa, 0x90, 0x3a, 0x74, 0xe5, 0xf2, 0x14, 0xfb, 0xef, 0x06, 0x67,
		0x67, 0x7d, 0xcf, 0x9b, 0xb0, 0x2a, 0xf7, 0xe3, 0x4d, 0x27, 0x02, 0xea, 0xdb, 0xbe,
		0x80, 0xeb, 0xcf, 0x94, 0x4c, 0x2a, 0x54, 0x2a, 0x98, 0x35, 0x59, 0xd9, 0x24, 0x8a},
	{0x50, 0xdf, 0xb7, 0xe7, 0x92, 0x92, 0xf3, 0xb0, 0x4e, 0x0d, 0x5c, 0x73, 0x8a, 0xf2,
		0xba, 0xc6, 0xda, 0xdf, 0x00, 0xe5, 0x37, 0x7b, 0xbf, 0xc1, 0xe7, 0x13, 0xe1, 0xda,
		0x5f, 0xa1, 0xa3, 0xc2, 0xfd, 0x4b, 0x10, 0x81, 0x0d, 0x99, 0xcf, 0x8f, 0xca, 0x91,
		0x37, 0x3e, 0x47, 0x8a, 0x84, 0xab, 0xcd, 0x65, 0xdf, 0xf9, 0x27, 0x3c, 0x13, 0xf1},
	{0xe4, 0xe1, 0xa4, 0x8d, 0x1d, 0x72, 0xe2, 0x72, 0x3b, 0x09, 0x09, 0xf9, 0x7f, 0xcd,
		0x57, 0x0d, 0xdf, 0x8c, 0xdc, 0x47, 0xdf, 0x6d, 0xfa, 0x6a, 0x8d, 0x67, 0x45, 0x4f,
		0x6b, 0x44, 0x6d, 0xbf, 0xf3, 0x41, 0x1c, 0x57, 0x1c, 0xf0, 0x77, 0x14, 0x06, 0xf6,
		0x8c, 0xb9, 0xa3, 0x40, 0x34, 0x70, 0xd6, 0x36, 0xe5, 0xa6, 0xce, 0x1b, 0x84, 0xcc},
}

func (s *Ed448Suite) Test_elligator_testVectors(c *C) {
	for ix, v := range elligatorTestVectorsInput {
		var inp [56]byte
		var exp serialized
		q := &twExtendedPoint{
			x: &bigNumber{},
			y: &bigNumber{},
			z: &bigNumber{},
			t: &bigNumber{},
		}
		copy(exp[:], elligatorTestVectorsOutput[ix])
		copy(inp[:], v)
		decafDecode(q, exp, false)
		p := pointFromNonUniformHash(inp)
		c.Assert(p.Equals(q), Equals, true)
	}
}

func (s *Ed448Suite) Test_PointFromUniformHash(c *C) {
	ser := [112]byte{
		0x15, 0x1d, 0xb9, 0x33, 0x08, 0xd8, 0xe2, 0xeb,
		0x97, 0x7c, 0x97, 0xf3, 0x16, 0xcf, 0x43, 0xd3,
		0x0b, 0x45, 0x67, 0x2b, 0x83, 0x17, 0xd7, 0x1c,
		0x39, 0xe8, 0x8a, 0xfc, 0x5f, 0x81, 0xa1, 0x4f,
		0x0d, 0x35, 0x86, 0xf0, 0xc3, 0xd9, 0x14, 0xa2,
		0x22, 0xa7, 0x49, 0x6c, 0x37, 0x87, 0x9c, 0x5c,
		0xd7, 0x8b, 0xd5, 0x8a, 0x4a, 0x03, 0x98, 0x7a,
		0xcc, 0x94, 0x46, 0xcb, 0xa2, 0xd6, 0x4b, 0xe0,
		0x09, 0xbe, 0xca, 0x3c, 0xbf, 0x3e, 0xab, 0xc2,
		0x7f, 0x57, 0xa8, 0x09, 0x1b, 0x02, 0x29, 0xae,
		0xc0, 0x8b, 0x1e, 0x70, 0x35, 0xcc, 0x7a, 0x04,
		0x37, 0x32, 0x96, 0x47, 0x38, 0x7e, 0xc7, 0xfe,
		0x43, 0x6f, 0x71, 0xc0, 0x0c, 0x29, 0xf3, 0x59,
		0x17, 0x3f, 0x5d, 0x49, 0x06, 0x8b, 0x03, 0xd9,
	}

	exp := &twExtendedPoint{
		&bigNumber{
			0x0733a0e6, 0x0565dc88, 0x05b952cf, 0x0ac980c7,
			0x06f8790a, 0x068b9aa4, 0x03ffb7ac, 0x0f9498f8,
			0x0c9f684f, 0x05a4009e, 0x0d3b77fa, 0x05d415d3,
			0x0c1b1986, 0x08e8f930, 0x0e71fded, 0x0c4e739,
		},
		&bigNumber{
			0x69f3a21, 0x68162b, 0xde028ac, 0xe630df1,
			0x11624e6, 0xb906e9d, 0x32bf83b, 0x6a2ecad,
			0xdcd4912, 0x3c6cb80, 0x42fea3e, 0x1e3abf9,
			0xa1cfe30, 0x7e2ef28, 0xde0fd5d, 0xab1f02f,
		},
		&bigNumber{
			0xefd11da, 0xf0d00fe, 0x2c0aa2, 0x7dad43d,
			0xc6d9a72, 0x36e95fc, 0x7581bda, 0xc15e0e6,
			0xfcdb211, 0xadccc30, 0x5ad0c9b, 0x9944eba,
			0x259e9b0, 0x80ed3be, 0xa45f8f4, 0xa45fef7,
		},

		&bigNumber{
			0x1303ca8, 0x508047a, 0x51aca8e, 0xb65f114,
			0xb5ed74b, 0x340d7df, 0x76e6f27, 0x251cd2a,
			0xf0bd360, 0xa047eb2, 0x50c4ad0, 0x76addbe,
			0x194e63a, 0xc910898, 0x265ddac, 0x8d02176,
		},
	}
	p := pointFromUniformHash(ser)
	c.Assert(p.x, DeepEquals, exp.x)
	c.Assert(p.y, DeepEquals, exp.y)
	c.Assert(p.z, DeepEquals, exp.z)
	c.Assert(p.t, DeepEquals, exp.t)
}

func (s *Ed448Suite) Test_InvertElligatorNonUniform(c *C) {
	p := &twExtendedPoint{
		&bigNumber{
			0x0fffffff, 0x0fffffff, 0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff, 0x0fffffff, 0x0fffffff,
			0x0ffffffe, 0x0fffffff, 0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff, 0x0fffffff, 0x0fffffff,
		},
		&bigNumber{
			0x00000000, 0x01000000, 0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff, 0x0fffffff, 0x0fffffff,
			0x0ffffffe, 0x0fffffff, 0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff, 0x0fffffff, 0x0fffffff,
		},
		&bigNumber{
			0x0ffffffe, 0x0fffffff, 0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff, 0x0fffffff, 0x0fffffff,
			0x0ffffffe, 0x0fffffff, 0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff, 0x0fffffff, 0x0fffffff,
		},

		&bigNumber{
			0x0fffffff, 0x0fffffff, 0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff, 0x0fffffff, 0x0fffffff,
			0x0ffffffe, 0x0fffffff, 0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff, 0x0fffffff, 0x0fffffff,
		},
	}

	exp := [56]byte{
		0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xfe, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	}

	ser, ok := invertElligatorNonUniform(p, word(0x01))
	c.Assert(ser, DeepEquals, exp)
	c.Assert(ok, Equals, true)

	q := &twExtendedPoint{
		&bigNumber{
			0x07a9207e, 0x016130d5, 0x010497c8, 0x075f03fb,
			0x0b50d815, 0x0b185d70, 0x0029d45a, 0x01d4d572,
			0x04939594, 0x092be6fb, 0x01b851dd, 0x080ba46f,
			0x00e58dd4, 0x02a256f7, 0x0a30d0a0, 0x095ec3cd,
		},
		&bigNumber{
			0x00923a40, 0x002eee8c, 0x0d19ff94, 0x04093df5,
			0x0f158f1e, 0x01d57e7a, 0x0a387f85, 0x0ea1c535,
			0x036af3c6, 0x098713a4, 0x047d691d, 0x044101c5,
			0x0d217bc4, 0x07a787b7, 0x06589af3, 0x0cbdfd81,
		},
		&bigNumber{
			0x038a5f41, 0x04f8ce5e, 0x02012801, 0x0d839806,
			0x0630513e, 0x0c4f0e13, 0x030358c0, 0x0074703a,
			0x0aa04899, 0x0d4f4768, 0x05c7127d, 0x0ef4c099,
			0x05e36d5b, 0x07f1e299, 0x01e079a0, 0x05909da2,
		},

		&bigNumber{
			0x0fdc305f, 0x04c698f8, 0x022c3e10, 0x0f1538b6,
			0x0ebf17a1, 0x0a31b24a, 0x03821060, 0x0e7cf360,
			0x0d8bc54d, 0x08d025e7, 0x019756d3, 0x0a7be4de,
			0x067dd8b9, 0x00c19582, 0x0044f64e, 0x096cd541,
		},
	}

	exp2 := [56]byte{
		0xbf, 0x4f, 0xb4, 0x45, 0x84, 0x59, 0x5c, 0x3c,
		0x4e, 0x04, 0x44, 0x87, 0xdb, 0x8d, 0xf6, 0x1d,
		0xe9, 0x0c, 0x04, 0x3c, 0x9f, 0x2d, 0x92, 0x09,
		0x5d, 0xe4, 0xb6, 0xb7, 0x0a, 0x76, 0x25, 0x38,
		0x92, 0x4c, 0xef, 0x12, 0xaf, 0x27, 0x3b, 0x00,
		0xee, 0xd7, 0xdb, 0x0c, 0x1c, 0xed, 0xa7, 0x0a,
		0x32, 0xf7, 0x52, 0x4a, 0x7e, 0x79, 0xa9, 0xd9,
	}

	ser2, ok2 := invertElligatorNonUniform(q, word(0x07))
	c.Assert(ser2, DeepEquals, exp2)
	c.Assert(ok2, Equals, true)
}

func (s *Ed448Suite) Test_InvertElligatorUniform(c *C) {
	p := &twExtendedPoint{
		&bigNumber{
			0x0fc18ee8, 0x06e52384, 0x090f51e9, 0x02090a94,
			0x0314d5f3, 0x0813d5d0, 0x01d8ea67, 0x0d01dcbe,
			0x0930c74f, 0x051b42c9, 0x04d2a414, 0x058c7d3b,
			0x0858ed59, 0x06cac045, 0x0d8a80e7, 0x021a5a8e,
		},
		&bigNumber{
			0x04232323, 0x09b24e49, 0x04384c1a, 0x001bf8e4,
			0x0c09a485, 0x00bf20b5, 0x0906679f, 0x0b88a665,
			0x073d389f, 0x053fc152, 0x0edab941, 0x08dd83c8,
			0x054e0716, 0x0eb04549, 0x01bb1c4e, 0x0b870d06,
		},
		&bigNumber{
			0x0f075470, 0x08535015, 0x05b5b1bb, 0x048e5d42,
			0x0f8c5a7c, 0x0050ecdf, 0x0ebdc995, 0x0ccb0cf2,
			0x0d048eda, 0x0180ed00, 0x08436d25, 0x04c69104,
			0x07e469aa, 0x0f7ea8ea, 0x09d25f90, 0x0f1960b9,
		},

		&bigNumber{
			0x08b93b9a, 0x0fd7d592, 0x0db0cc77, 0x02fdd1c2,
			0x0dcd7ad4, 0x0b2aea7e, 0x01f71af2, 0x086a9d4b,
			0x092bbef3, 0x0a7ac451, 0x0eb40b5d, 0x04560fbd,
			0x057107d7, 0x01204ba7, 0x0c1740f4, 0x0a7b08e0,
		},
	}

	src := [112]byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x24, 0xcc, 0x70, 0x7e, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	exp := [112]byte{
		0x7f, 0x48, 0x75, 0xc5, 0x6d, 0xc0, 0xb5, 0x3d,
		0xfd, 0x10, 0xfa, 0x38, 0x80, 0x16, 0x9c, 0x9c,
		0x06, 0xc0, 0xa7, 0xdb, 0x75, 0x9f, 0xa9, 0x8e,
		0x02, 0x65, 0xd9, 0x8d, 0x10, 0x74, 0x12, 0x88,
		0xef, 0x66, 0x6a, 0x4c, 0x05, 0xe9, 0x64, 0xb6,
		0xcf, 0xb3, 0x53, 0x30, 0xba, 0x92, 0x3f, 0x08,
		0x27, 0xa0, 0x83, 0x68, 0x19, 0xe9, 0x5b, 0xaf,
		0x24, 0xcc, 0x70, 0x7e, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	ser, ok := invertElligatorUniform(src, p, word(0x06))
	c.Assert(ser, DeepEquals, exp)
	c.Assert(ok, Equals, true)
}
