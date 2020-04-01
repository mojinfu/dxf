package entity

import (
	"github.com/mojinfu/dxf/format"
	"github.com/mojinfu/point"
)

// LwPolyline represents LWPOLYLINE Entity.
type LwPolyline struct {
	*entity
	Num      int // 90
	Closed   bool
	Vertices [][]float64
	Bulges   []float64
}

// IsEntity is for Entity interface.
func (l *LwPolyline) IsEntity() bool {
	return true
}

// NewLwPolyline creates a new LwPolyline.
func NewLwPolyline(size int) *LwPolyline {
	vs := make([][]float64, size)
	bulges := make([]float64, size)
	for i := 0; i < size; i++ {
		vs[i] = make([]float64, 2)
	}
	l := &LwPolyline{
		entity:   NewEntity(LWPOLYLINE),
		Num:      size,
		Closed:   false,
		Vertices: vs,
		Bulges:   bulges,
	}
	return l
}

// Format writes data to formatter.
func (l *LwPolyline) Format(f format.Formatter) {
	l.entity.Format(f)
	f.WriteString(100, "AcDbPolyline")
	f.WriteInt(90, l.Num)
	if l.Closed {
		f.WriteInt(70, 1)
	} else {
		f.WriteInt(70, 0)
	}
	for i := 0; i < l.Num; i++ {
		for j := 0; j < 2; j++ {
			f.WriteFloat((j+1)*10, l.Vertices[i][j])
		}
	}
}

// String outputs data using default formatter.
func (l *LwPolyline) String() string {
	f := format.NewASCII()
	return l.FormatString(f)
}

// FormatString outputs data using given formatter.
func (l *LwPolyline) FormatString(f format.Formatter) string {
	l.Format(f)
	return f.Output()
}

// Close closes LwPolyline.
func (l *LwPolyline) Close() {
	l.Closed = true
}

func (l *LwPolyline) BBox() ([]float64, []float64) {
	mins := make([]float64, 3)
	maxs := make([]float64, 3)
	for _, p := range l.Vertices {
		for i := 0; i < len(p); i++ {
			if p[i] < mins[i] {
				mins[i] = p[i]
			}
			if p[i] > maxs[i] {
				maxs[i] = p[i]
			}
		}
	}
	return mins, maxs
}
func (l *LwPolyline) Poly() []*point.Point {
	poly := []*point.Point{}
	for i := range l.Vertices {
		//log.Println(len(l.Vertices[i]))
		if i == len(l.Vertices)-1 {
			//nest.Bulge2Arc(&point.Point{X: 0, Y: 0}, &point.Point{X: 100, Y: 0}, -0.4143)
		} else {
			pointA := &point.Point{
				X: l.Vertices[i][0],
				Y: l.Vertices[i][1],
			}
			pointB := &point.Point{
				X: l.Vertices[i+1][0],
				Y: l.Vertices[i+1][1],
			}
			poly = append(poly, point.Bulge2Arc(pointA, pointB, l.Bulges[i])...)
		}
	}
	return poly
}
