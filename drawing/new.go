package drawing

import (
	"math"

	"github.com/mojinfu/dxf/entity"
)

func (d *Drawing) ReplaceCircle(circle *entity.Circle, x, y float64) (*entity.Circle, error) {
	return d.Circle(circle.CurrentCoord()[0]+x, circle.CurrentCoord()[1]+y, circle.CurrentCoord()[2], circle.Radius)
}

func (d *Drawing) ReplaceArc(arc *entity.Arc, x, y float64, degrees int) (*entity.Arc, error) {
	new := []float64{}
	for i := range arc.Angle {
		new = append(new, arc.Angle[i]+float64(degrees))
	}
	return d.Arc(arc.CurrentCoord()[0]+x, arc.CurrentCoord()[1]+y, arc.CurrentCoord()[2], arc.Radius, new[0], new[1])
}

func (d *Drawing) ReplaceLwPolyline(lwPoly *entity.LwPolyline, x, y float64, degrees int) (*entity.LwPolyline, error) {
	newP := [][]float64{}
	newB := []float64{}
	for i := range lwPoly.Vertices {
		newP = append(newP, translatePlanePoint(x, y, rotationPlanePoint(degrees, lwPoly.Vertices[i])))
	}
	for i := range lwPoly.Bulges {
		newB = append(newB, lwPoly.Bulges[i])
	}
	return d.LwPolylineWithBulges(lwPoly.Closed, newP, newB)
}

func (d *Drawing) ReplacePolyline(poly *entity.Polyline, x, y float64, degrees int) (*entity.Polyline, error) {
	new := [][]float64{}
	for i := range poly.Vertices {
		new = append(new, translatePlanePoint(x, y, rotationPlanePoint(degrees, poly.Vertices[i].Coord)))
	}
	newpoly, err := d.Polyline(false, new...)
	if err != nil {
		return nil, err
	}
	newpoly.Flag = poly.Flag
	return newpoly, nil
}

func translatePlanePoint(x, y float64, point []float64) []float64 {
	var x1 = point[0] + x
	var y1 = point[1] + y
	var z1 = point[2]
	return []float64{x1, y1, z1}

}
func rotationPlanePoint(degrees int, point []float64) []float64 {
	angle := float64(degrees) * math.Pi / 180
	var x1 = point[0]*math.Cos(angle) - point[1]*math.Sin(angle)
	var y1 = point[0]*math.Sin(angle) + point[1]*math.Cos(angle)
	var z1 = point[2]
	return []float64{x1, y1, z1}
}
