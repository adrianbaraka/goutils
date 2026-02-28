package coords

import "fmt"

type Coordinate struct {
	Row, Column int
}

// Returns a new coordinate type
func NewCoordinate(row, column int) Coordinate {
	return Coordinate{Row: row, Column: column}
}

type Two_Darray[T any] struct {
	Length int
	Columns int
	FlatArray [] T
}

// Returns a new 2-D array with elements of type T and initialize all items to the zero value of T
func NewTwo_Darray [T any] (length, columns int) *Two_Darray[T]{
	var zero T
	ar := make([]T, 0, length)
	// initialize all values to zero value
	for range length {
		ar = append(ar, zero)
	}

	return &Two_Darray[T]{
		Length: length,
		Columns: columns,
		FlatArray: ar,
	}
}

// Get the index of the given coordinate in the flat array
func (a *Two_Darray[T]) GetIndex(c Coordinate) (int, error){
	// rows start at 0
	numRows := a.NumRows()
	if c.Row > numRows {
		return 0, fmt.Errorf("Invalid row %v in array with %v rows", c.Row, numRows)
	}
	if c.Column > a.Columns {
		return 0, fmt.Errorf("Invalid column %v in array with %v columns.", c.Column, a.Columns)
	}
	return (c.Row * a.Columns ) + c.Column, nil
}

// Set a state of a given coordinate takes type T values
func (a *Two_Darray[T]) Set(state T, c Coordinate) error {

	idx, err := a.GetIndex(c)
	if err != nil {
		return err
	}

	a.FlatArray[idx] = state
	return nil
}

// Get the value at a given coordinate 
func (a *Two_Darray[T]) GetVal(c Coordinate) (T, error){
	idx, err := a.GetIndex(c)
	var zero T // zero value of T
	if err != nil {
		return zero, err
	}
	return a.FlatArray[idx], nil
}

// Visualize the 2d array
func (a *Two_Darray[T]) Visual() {
	fmt.Println()
	for i := 0; i < a.Length; i += a.Columns {
		fmt.Printf("\t")
		for j:= i; j < i + a.Columns; j++ {
			fmt.Printf("%v", a.FlatArray[j])
		}
		fmt.Println()
	}
	fmt.Println()
}
// Get the number of rows in the 2D-grid
func (a *Two_Darray[T]) NumRows() int {
	return a.Length / a.Columns
}

// Returns a slice of all the valid 8 neighbours of the given coordinate.
// 	If invalid coordinates are given it will simply return an empty slice.
func (a *Two_Darray[T]) GetNeighbours(c Coordinate) ([]Coordinate) {
	neighbours := [] Coordinate {}
	n := [...] int { -1, 0, 1}

	for _ ,rd := range n {
		for _,cd := range n {
			if rd == 0 && cd == 0 {
				continue // element itself skip over it
			}

			r := c.Row + rd
			c := c.Column + cd

			if r >= 0 && r < a.NumRows() && c >= 0 && c < a.Columns {
				ord := Coordinate{r,c}
				neighbours = append(neighbours, ord)
			}
		}
	}
	return neighbours
}

// Returns a slice of the valid 4 neighbours of c not including diagnals. 
func (a *Two_Darray[T]) GetNeighbours4(c Coordinate) [] Coordinate {
	neighbours := [] Coordinate {}
	//n := [...] int { -1, 0, 1}

	n := make(map [string] [2] int)

	n["up"] = [...] int {-1, 0}
	n["down"] = [...] int {1, 0}
	n["left"] = [...] int {0, -1}
	n["right"] = [...] int {0, 1}


	for _, val := range n {
		ro := c.Row + val[0]
		co := c.Column + val[1]

		if ro >= 0 && co >= 0 && ro < a.NumRows() && co < a.Columns {
			ord := Coordinate{ro,co}
			neighbours = append(neighbours, ord)
		}
	}
	//fmt.Println(neighbours)
	return neighbours
}

// return a map of neighbours with keys the name
func (a *Two_Darray[T]) GetNeighboursMap(c Coordinate) (map[ string ] Coordinate) {
	neigbours := make(map [string] Coordinate)

	positions := [...] string {"upLeft", "up", "upRight", "left", "right", "downLeft", "down", "downRight"}
	n := [...] int { -1, 0, 1}
	var i int
	for _ ,rd := range n {
		for _,cd := range n {
			if rd == 0 && cd == 0 {
				continue // element itself skip over it
			}

			r := c.Row + rd
			c := c.Column + cd

			if r >= 0 && r < a.NumRows() && c >= 0 && c < a.Columns {
				ord := Coordinate{r,c}
				neigbours[positions[i]] = ord
			}
			i ++
		}
	}
	return neigbours
}
