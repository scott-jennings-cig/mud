package mud

import (
	"bytes"
	"encoding/binary"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// MessageType is a log message line type
type MessageType int

// Message types for log items
const (
	MESSAGESYSTEM MessageType = iota
	MESSAGECHAT
	MESSAGEACTION
	MESSAGEACTIVITY
)

// LogItem is individual chat log line
type LogItem struct {
	Message     string      `json:""`
	Author      string      `json:""`
	Timestamp   time.Time   `json:""`
	MessageType MessageType `json:""`
	Location    *Point      `json:",omit"`
}

// Point represents an (X,Y) pair in the world
type Point struct {
	X uint32
	Y uint32
}

// Add applies a vector to a point
func (p *Point) Add(v Vector) Point {
	return Point{
		X: uint32(int(p.X) + v.X),
		Y: uint32(int(p.X) + v.X)}
}

// Vector Gets the vector between two points such that v = p.Vector(q); p.Add(v) == q
func (p *Point) Vector(v Point) Vector {
	return Vector{
		X: int(v.X) - int(p.X),
		Y: int(v.Y) - int(p.Y)}
}

// Vector is for doing point-to-point comparisons
type Vector struct {
	X int
	Y int
}

// Add ccombines two vectors
func (v *Vector) Add(p Vector) Vector {
	return Vector{
		X: v.X + p.X,
		Y: v.Y + p.Y}
}

// Magnitude returns the pythagorean theorem to a vector
func (v *Vector) Magnitude() uint {
	return uint(math.Sqrt(math.Pow(float64(v.X), 2.0) + math.Pow(float64(v.Y), 2.0)))
}

// Bytes Dumps a point into a byte array
func (p *Point) Bytes() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, p)
	return buf.Bytes()
}

// PointFromBytes rehydrates a point struct
func PointFromBytes(ptBytes []byte) Point {
	var pt Point
	buf := bytes.NewBuffer(ptBytes)
	binary.Read(buf, binary.LittleEndian, &pt)
	return pt
}

// Direction is a cardinal direction
type Direction int

// Cardinal directions
const (
	DIRECTIONNORTH Direction = iota
	DIRECTIONEAST
	DIRECTIONSOUTH
	DIRECTIONWEST
)

// ToTheRight gives the direction to the right of the current one
func ToTheRight(d Direction) Direction {
	switch d {
	case DIRECTIONNORTH:
		return DIRECTIONEAST
	case DIRECTIONEAST:
		return DIRECTIONSOUTH
	case DIRECTIONSOUTH:
		return DIRECTIONEAST
	case DIRECTIONWEST:
		return DIRECTIONNORTH
	}

	return DIRECTIONNORTH
}

// ToTheLeft gives the direction to the rigleftt of the current one
func ToTheLeft(d Direction) Direction {
	switch d {
	case DIRECTIONNORTH:
		return DIRECTIONWEST
	case DIRECTIONWEST:
		return DIRECTIONSOUTH
	case DIRECTIONSOUTH:
		return DIRECTIONEAST
	case DIRECTIONEAST:
		return DIRECTIONNORTH
	}

	return DIRECTIONNORTH
}

// VectorForDirection maps directions to a distance vector
var VectorForDirection map[Direction]Vector

// DirectionForVector maps vectors to directions
var DirectionForVector map[Vector]Direction

// LoadResources loads data for the game
func LoadResources() {
	loadCreatureTypes("./bestiary.json")
	loadItemTypes("./items.json")
	loadTerrainTypes("./terrain.json")
}

// MakeTransitionFunction helps build Markov chains.
func MakeTransitionFunction(name string, transitionList []string) (func() string, []string) {
	total := 0

	type transitionName struct {
		name   string
		weight int
	}

	transitionInternalList := make([]transitionName, 0)
	returnTransitionList := make([]string, 0)

	for _, transition := range transitionList {
		splitString := strings.SplitN(transition, ":", 2)
		weightString := "1"
		returnTransitionList = append(returnTransitionList, splitString[0])

		if (len(splitString)) > 1 {
			weightString = splitString[1]
		}

		weight, err := strconv.Atoi(weightString)

		if err != nil {
			panic(err)
		}

		transitionInternalList = append(transitionInternalList, transitionName{name: splitString[0], weight: weight})
		total += weight
	}

	return func() string {
		if transitionInternalList != nil && len(transitionInternalList) != 0 {
			weight := 0
			countTo := rand.Int() % total

			for _, item := range transitionInternalList {
				weight += item.weight

				if weight > countTo {
					return item.name
				}
			}
		}
		return ""
	}, returnTransitionList
}

func init() {
	VectorForDirection = map[Direction]Vector{
		DIRECTIONNORTH: Vector{X: 0, Y: -1},
		DIRECTIONEAST:  Vector{X: -1, Y: 0},
		DIRECTIONSOUTH: Vector{X: 0, Y: 1},
		DIRECTIONWEST:  Vector{X: 1, Y: 0}}
	DirectionForVector = make(map[Vector]Direction)
	for k, v := range VectorForDirection {
		DirectionForVector[v] = k
	}
}
