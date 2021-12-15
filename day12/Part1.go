package main

import (
	"adventcodingchallenge_2021/utility"
	"strings"
)

type Part1 struct {
	data map[string]*Cave
}

func Parse(data []string) map[string]*Cave {

	aMap := make(map[string]*Cave, 0)

	for _, row := range data {

		strings := strings.Split(row, "-")

		connections := make([]*Cave, 0)
		for _, aString := range strings {
			var aCave *Cave
			var ok bool
			if aCave, ok = aMap[aString]; !ok {
				size := 0
				if utility.IsUpper(aString) {
					size = 1
				}
				aCave = &Cave{
					size:        size,
					description: aString,
					connections: make([]*Cave, 0),
				}
				aMap[aCave.description] = aCave
			}
			connections = append(connections, aCave)

		}
		caveA := connections[0]
		caveb := connections[1]
		caveA.addConnectionTo(caveb)

	}
	return aMap
}

func (alg *Part1) Process(data []string) (error, interface{}) {
	alg.data = Parse(data)

	caveSystem := &CaveSystem{
		mapOfCaves: alg.data,
		pathTree:   nil,
	}
	caveSystem.findPaths()
	return nil, 0
}

func (c *Cave) isBig() bool {
	return c.size == 1
}

func (c *Cave) isSmall() bool {
	return c.size == 0
}
func (c *Cave) isAlreadyConnected(aTarget *Cave) bool {
	found := false
	for _, aCave := range c.connections {
		found = aCave.description == aTarget.description
		if found {
			break
		}
	}
	return found
}

func (c *Cave) addConnectionTo(target *Cave) {
	if !c.isAlreadyConnected(target) {
		c.connections = append(c.connections, target)
		//target.addConnectionTo(c)
	}

}

type Cave struct {
	size        int
	description string
	connections []*Cave
}

type CaveSystem struct {
	mapOfCaves map[string]*Cave
	pathTree   *Path
}

func (c *CaveSystem) findPaths() {
	start := c.mapOfCaves["start"]
	//end := c.mapOfCaves["end"]

	c.pathTree = start.buildTree()

	if c.pathTree != nil {
		stack := utility.NewSimpleStringStack()
		addToStack(c.pathTree.headOfPath, c.pathTree, stack)

	}
}

func addToStack(aCave *Cave, path *Path, stack *utility.SimpleStringStack) {

	var previousCave string
	if stack.HasMore() {
		previousCave = stack.Peek()
	}

	stack.Push(aCave.description)
	if len(path.paths) == 0 {

		stack.Push(previousCave)
	} else {

	}

	for _, aPath := range path.paths {
		addToStack(aPath.headOfPath, aPath, stack)
	}

}

type Path struct {
	headOfPath *Cave
	paths      []*Path
}

func (p *Path) addPathToCave(aPath *Path, aCave *Cave) {
	p.paths = append(p.paths, aPath)
}

func (c *Cave) buildTree() *Path {

	aPath := &Path{paths: make([]*Path, 0), headOfPath: c}

	for _, aCave := range c.connections {
		anotherPath := aCave.buildTree()
		aPath.addPathToCave(anotherPath, aCave)
	}

	return aPath
}
