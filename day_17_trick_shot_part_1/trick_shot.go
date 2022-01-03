package main

import (
	"fmt"
	"math"
)

// Position x(t):
// if vx - (t-1) > 0 i.e. t <= vx then
//   x(t) = vx + vx - 1 + ... + vx - (t-1) = t * vx - t * (t - 1) / 2
// if t >= vx + 1 then
//   x(t) = vx + vx - 1 + ... + 1 = vx * (vx + 1) / 2
// ----
// We need vx * (vx + 1) / 2 >= x0, i.e.
//   1/2 [vx^2 + vx] = 1/2 [(vx + 1/2)^2 - 1/4] >= x0
//   (vx + 1/2)^2 >= 2 x0 + 1/4
//   vx >= sqrt(2 x0 + 1/4) - 1/2
// ----
// x(t) = t * vx - t^2 / 2 + t / 2 = -1/2 [t^2 - (2 vx + 1) * t]
// Solving x(y) == x0 is as follows:
// x(t) = -1/2 [(t - (vx + 1/2))^2  - (vx + 1/2)^2] = x0
// (t - (vx + 1/2))^2  - (vx + 1/2)^2 = -2 x0
// (t - (vx + 1/2))^2 = (vx + 1/2)^2 - 2 x0
// If t <= vx then t - (vx + 1/2) < 0, so:
//   t = vx + 1/2 - sqrt[(vx + 1/2)^2 - 2 x0]
//
// Position y(t) = vy + vy - 1 + ... + vy - (t-1)
// The maximum height = vy + vy - 1 + ... + 1 = vy * (vy + 1) / 2
// Solving y(t) = y0 yields:
//   t = vy + 1/2 + sqrt[(vy + 1/2)^2 - 2 y0]
// Because we always have: t > vy + 1/2
//
// Constraints:
//   sqrt(2 x1 + 1/4) - 1/2 <= vx <= x2
//   y1 <= vy <= -y1

func sqr(x float64) float64 {
	return x * x
}

func timeX(vx, x1, x2 float64) (int, int) {
	if vx*(vx+1)/2 < x1 {
		return -1, -1
	}

	d2 := sqr(vx+0.5) - 2*x1
	t1 := vx + 0.5 - math.Sqrt(d2)
	t2 := 1.0e9
	if vx*(vx+1)/2 > x2 {
		d2 = sqr(vx+0.5) - 2*x2
		t2 = vx + 0.5 - math.Sqrt(d2)
	}
	return int(math.Ceil(t1)), int(math.Floor(t2))
}

func timeY(vy, y1, y2 float64) (int, int) {
	t1 := vy + 0.5 + math.Sqrt(sqr(vy+0.5)-2*y2)
	t2 := vy + 0.5 + math.Sqrt(sqr(vy+0.5)-2*y1)
	return int(math.Ceil(t1)), int(math.Floor(t2))
}

func main() {
	// x1, x2 := 20, 30
	// y1, y2 := -10, -5
	x1, x2 := 70, 125
	y1, y2 := -159, -121

	fx1, fx2 := float64(x1), float64(x2)
	fy1, fy2 := float64(y1), float64(y2)
	h := 0
	x0 := int(math.Ceil(math.Sqrt(2*fx1+0.25) - 0.5))
	for vy := -y1; vy >= y1; vy-- {
		f := false
		for vx := x0; vx <= x2; vx++ {
			xt1, xt2 := timeX(float64(vx), fx1, fx2)
			if xt1 == -1 || xt1 > xt2 {
				continue
			}

			yt1, yt2 := timeY(float64(vy), fy1, fy2)
			if yt1 > yt2 {
				continue
			}

			if xt2 >= yt1 && yt2 >= xt1 {
				h = vy * (vy + 1) / 2
				f = true
				break
			}
		}
		if f {
			break
		}
	}

	fmt.Println(h)
}
