package utility

type Coordinate struct {
	X int
	Y int
}

type CoordinateRange struct {
	Start *Coordinate
	End   *Coordinate
}

func (c *CoordinateRange) DetermineCoordinatesInRange(includeDiag bool) []*Coordinate {
	coordinates := make([]*Coordinate, 0)
	if c.Start.X == c.End.X {
		x := c.Start.X
		numbers := []int{c.Start.Y, c.End.Y}
		least, max := LeastAndMax(numbers)
		for y := least; y <= max; y++ {
			aCoordinate := &Coordinate{
				X: x,
				Y: y,
			}
			coordinates = append(coordinates, aCoordinate)
		}

	} else if c.Start.Y == c.End.Y {

		y := c.Start.Y
		//diff := int(math.Abs(float64(c.Start.x - c.end.x)))
		numbers := []int{c.Start.X, c.End.X}
		least, max := LeastAndMax(numbers)
		for x := least; x <= max; x++ {
			aCoordinate := &Coordinate{
				X: x,
				Y: y,
			}
			coordinates = append(coordinates, aCoordinate)
		}

	} else {
		if includeDiag {
			xNumbers := NumbersBetween(c.Start.X, c.End.X)
			OrderNumbersStartingWithAndEndingWith(xNumbers, c.Start.X, c.End.X)
			yNumbers := NumbersBetween(c.Start.Y, c.End.Y)
			OrderNumbersStartingWithAndEndingWith(yNumbers, c.Start.Y, c.End.Y)

			for j := 0; j < len(xNumbers); j++ {
				aCoordinate := &Coordinate{
					X: xNumbers[j],
					Y: yNumbers[j],
				}
				coordinates = append(coordinates, aCoordinate)
			}
		}

	}

	return coordinates
}

func PrettyPrint(coordinates []*Coordinate) {

}
