package smooch

import (
	"sort"
	"time"
	"fmt"
	"strings"
)

type Unit struct {
	Size int64
	Name string
}

func (unit Unit) ComparableTo(x int64) bool {
	if (unit.Size / x == 0) {
		return false
	}
	return true
}

type Scale []Unit

func ScaleOf(units ...Unit) Scale {
	sort.Slice(units, func(i, j int) bool {
		return units[i].Size > units[j].Size
	})
	return Scale(units)
}

var TimeScale = ScaleOf(Scale{
	{int64(time.Hour*24*365), "year"},
	{int64(time.Hour*24*30),  "month"},
	{int64(time.Hour*24*7),   "week"},
	{int64(time.Hour*24),     "day"},
	{int64(time.Hour),        "hour"},
	{int64(time.Minute),      "minute"},
	{int64(time.Second),      "second"},
}...)

func (scale Scale) Format(x, precision int64, concat bool) string {
	// sort the units in descending order
	scale = ScaleOf(scale...)
	// skip units incomparable to the input until we hit the last on the scale
	for len(scale) > 1 && scale[0].ComparableTo(x) {
		scale = scale[1:]
	}
	if !concat {
		exp := float64(x) / float64(scale[0].Size)
		return fmt.Sprintf("%.1f %ss", exp, scale[0].Name)
	}
	var B strings.Builder
	last := func(i int) bool {
		return i == len(scale) - 1 || !scale[i+1].ComparableTo(precision)
	}
	for i := range scale {
		if i > 0 {
			B.WriteString(", ")
			if last(i) {
				B.WriteString("and ")
			}
		}
		remainder := x % scale[i].Size
		if !last(i) {
			exp := x / scale[i].Size
			B.WriteString(fmt.Sprintf("%d %s", exp, scale[i].Name))
			if exp != 1 {
				B.WriteByte('s')
			}
		} else {
			exp := float64(x) / float64(scale[i].Size)
			B.WriteString(fmt.Sprintf("%.1f %ss", exp, scale[i].Name))
			break
		}
		x = remainder
	}
	return B.String()
}
