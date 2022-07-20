package gg

import "golang.org/x/image/draw"

// ScaleStyle determines the way image pixels are interpolated when scaled.
// See
// 	https://pkg.go.dev/golang.org/x/image/draw
// for the corresponding interpolators.
type ScaleStyle int

const (
	// BiLinear is the tent kernel. It is slow, but usually gives high quality
	// results.
	BiLinear ScaleStyle = iota

	// ApproxBiLinear is a mixture of the nearest neighbor and bi-linear
	// interpolators. It is fast, but usually gives medium quality results.
	//
	// It implements bi-linear interpolation when upscaling and a bi-linear
	// blend of the 4 nearest neighbor pixels when downscaling. This yields
	// nicer quality than nearest neighbor interpolation when upscaling, but
	// the time taken is independent of the number of source pixels, unlike the
	// bi-linear interpolator. When downscaling a large image, the performance
	// difference can be significant.
	ApproxBiLinear

	// NearestNeighbor is the nearest neighbor interpolator. It is very fast,
	// but usually gives very low quality results. When scaling up, the result
	// will look 'blocky'.
	NearestNeighbor

	// CatmullRom is the Catmull-Rom kernel. It is very slow, but usually gives
	// very high quality results.
	//
	// It is an instance of the more general cubic BC-spline kernel with parameters
	// B=0 and C=0.5. See Mitchell and Netravali, "Reconstruction Filters in
	// Computer Graphics", Computer Graphics, Vol. 22, No. 4, pp. 221-228.
	CatmullRom
)

func (s ScaleStyle) transformer() draw.Interpolator {
	switch s {
	case BiLinear:
		return draw.BiLinear
	case ApproxBiLinear:
		return draw.ApproxBiLinear
	case NearestNeighbor:
		return draw.NearestNeighbor
	case CatmullRom:
		return draw.CatmullRom
	}
	return draw.BiLinear // BiLinear by default.
}

func (dc *Context) SetScaleStyle(s ScaleStyle) {
	dc.scaleStyle = s
}

func (dc *Context) SetScaleBiLinear() {
	dc.SetScaleStyle(BiLinear)
}

func (dc *Context) SetScaleApproxBiLinear() {
	dc.SetScaleStyle(ApproxBiLinear)
}

func (dc *Context) SetScaleNearestNeighbor() {
	dc.SetScaleStyle(NearestNeighbor)
}

func (dc *Context) SetScaleCatmullRom() {
	dc.SetScaleStyle(CatmullRom)
}
