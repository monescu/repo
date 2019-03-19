package homework

//Point represents a coordonate for a circle
type Point struct {
	X int
	Y int
}

//FindSmallestCircle find the minimal radius
func FindSmallestCircle(points []Point) int {

	centre := points[0]
	points = points[1:]

	radius := 0

	for !areAllPointsInCircle(points, centre, radius) {
		radius++
	}

	return radius
}

//FindSmallestCircleMatrix solves the problem using a matrix of distances
func FindSmallestCircleMatrix(points []Point) (int, Point) {

	distanceMatrix := make([][]int, len(points))

	for index := 0; index < len(points); index++ {
		distanceMatrix[index] = make([]int, len(points))
	}

	for i := 0; i < len(points); i++ {
		for j := 0; j < i; j++ {
			if i == j {
				distanceMatrix[i][j] = 0
			} else {
				distanceMatrix[i][j] = calculateDistanceSquared(points[i], points[j])
			}
		}
	}

	maxSlice := []int{}

	for i := 0; i < len(points); i++ {
		maxSlice = append(maxSlice, maxOfArray(distanceMatrix[i]))
	}

	smallest, position := minAndPositionOfArray(maxSlice)

	return smallest, points[position]
}

//-------------------PRIVATES--------------------------

func calculateDistanceSquared(centre Point, another Point) int {
	return (another.X-centre.X)*(another.X-centre.X) + (another.Y-centre.Y)*(another.Y-centre.Y)
}

func areAllPointsInCircle(points []Point, centre Point, radius int) bool {
	for _, point := range points {
		if calculateDistanceSquared(centre, point) > radius*radius {
			return false
		}
	}

	return true
}

func maxOfArray(array []int) int {
	max := array[0]
	for _, value := range array {
		if max < value {
			max = value
		}

	}
	return max
}

func minAndPositionOfArray(array []int) (int, int) {
	min := array[0]
	pos := 0
	for index, value := range array {
		if min > value {
			min = value
			pos = index
		}
	}
	return min, pos
}

//-----------------------------------------------------
