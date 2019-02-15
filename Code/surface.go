// Surface computes an SVG rendering of a 3-D surface function.
// 
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320               // canvas size in pixels
	cells		  = 100					   // number of grid cells
	xyrange		  = 30.0				   // axis ranges (-xyrange..+xyrange)
	xyscale		  = width / 2 / xyrange    // pixels per x or y unit
	zscale		  = height * 0.4 		   // pixels per z unit
	angle		  = math.Pi / 6 		   // angle of x, y axes (=30°)
)

var sin