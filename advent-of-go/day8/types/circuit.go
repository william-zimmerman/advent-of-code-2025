package types

import (
	"iter"
	"maps"
	"slices"

	"github.com/google/uuid"
	"github.com/samber/lo/it"
)

type JunctionBox struct {
	X, Y, Z int
}

type JunctionBoxPair struct {
	Box1, Box2 JunctionBox
	Distance   float64
}

type set[T comparable] map[T]struct{}

type Circuit set[JunctionBox]

func (c Circuit) add(boxes ...JunctionBox) {
	for _, box := range boxes {
		c[box] = struct{}{}
	}
}

func (c Circuit) Len() int {
	return it.Length(maps.Keys(c))
}

func (c Circuit) all() iter.Seq[JunctionBox] {
	return maps.Keys(c)
}

type circuitMap struct {
	circuitIdByJunctionBox map[JunctionBox]uuid.UUID
	circuitById            map[uuid.UUID]Circuit
}

func InitCircuitMap(boxes ...JunctionBox) circuitMap {
	circuitMap := circuitMap{map[JunctionBox]uuid.UUID{}, map[uuid.UUID]Circuit{}}

	for _, box := range boxes {
		circuitId := uuid.New()
		circuitMap.circuitIdByJunctionBox[box] = circuitId
		circuit := Circuit{}
		circuit.add(box)
		circuitMap.circuitById[circuitId] = circuit
	}

	return circuitMap
}

func (c *circuitMap) Connect(b1, b2 JunctionBox) {
	box1CircuitId, box1BelongsToCircuit := c.circuitIdByJunctionBox[b1]
	box2CircuitId, box2BelongsToCircuit := c.circuitIdByJunctionBox[b2]

	if !box1BelongsToCircuit || !box2BelongsToCircuit {
		panic("expecting both junction boxes to belong to circuit map")
	}

	if box1CircuitId == box2CircuitId {
		return
	}

	for box := range c.circuitById[box2CircuitId].all() {
		c.circuitIdByJunctionBox[box] = box1CircuitId
		c.circuitById[box1CircuitId].add(box)
		delete(c.circuitById, box2CircuitId)
	}
}

func (c circuitMap) Circuits() []Circuit {
	return slices.Collect(maps.Values(c.circuitById))
}

func (c circuitMap) UniqueCircuitCount() int {
	return it.Length(maps.Keys(c.circuitById))
}
