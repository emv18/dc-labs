package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
)

//This method is a struct type point
type Point struct {
	X, Y float64
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//generatePoints array
func generatePoints(s string) ([]Point, error) {

	points := []Point{}

	s = strings.Replace(s, "(", "", -1)
	s = strings.Replace(s, ")", "", -1)
	vals := strings.Split(s, ",")
	if len(vals) < 2 {
		return []Point{}, fmt.Errorf("Point [%v] was not well defined", s)
	}

	var x, y float64

	for idx, val := range vals {

		if idx%2 == 0 {
			x, _ = strconv.ParseFloat(val, 64)
		} else {
			y, _ = strconv.ParseFloat(val, 64)
			points = append(points, Point{x, y})
		}
	}
	return points, nil
}

// Distance gets the distance between points
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// getArea gets the area inside from a given shape
func getArea(points []Point) float64 {
	vertices := len(points)
	sum1 := 0.0
	sum2 := 0.0

	for i := 0; i < vertices-1; i++ {
		sum1 += points[i].X * points[i+1].Y
		sum2 += points[i].Y * points[i+1].X
	}

	sum1 += points[vertices-1].X * points[0].Y
	sum2 += points[0].X * points[vertices-1].Y

	area := math.Abs((sum1 - sum2) / 2.0)
	return area
}

// getPerimeter gets the perimeter from a given array of connected points
func getPerimeter(points []Point) float64 {
	perimeter := 0.0

	for i := 0; i < len(points); i++ {
		if i == len(points)-1 {
			perimeter += Distance(points[i], points[0])
		} else {
			perimeter += Distance(points[i], points[i+1])
		}
	}

	return perimeter
}

// handler handles the web request and reponds it
func handler(w http.ResponseWriter, r *http.Request) {

	var vertices []Point
	for k, v := range r.URL.Query() {
		if k == "vertices" {
			points, err := generatePoints(v[0])
			if err != nil {
				fmt.Fprintf(w, fmt.Sprintf("error: %v", err))
				return
			}
			vertices = points
			break
		}
	}

	// Results gathering
	area := getArea(vertices)
	perimeter := getPerimeter(vertices)

	// Logging in the server side
	log.Printf("Received vertices array: %v", vertices)

	// Response construction
	response := fmt.Sprintf("Welcome to my Remote Shapes Analyzer\n")
	response += fmt.Sprintf(" - Your figure has : [%v] vertices\n", len(vertices))
	if len(vertices) < 3 {
		response += fmt.Sprintf("ERROR - Your shape is not compliying with the minimum number of vertices.")
	} else {
		response += fmt.Sprintf(" - Vertices        : %v\n", vertices)
		response += fmt.Sprintf(" - Perimeter       : %v\n", perimeter)
		response += fmt.Sprintf(" - Area            : %v\n", area)
	}
	// Send response to client
	fmt.Fprintf(w, response)
}
