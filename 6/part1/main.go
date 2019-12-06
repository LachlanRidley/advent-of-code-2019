package main

import (
	"fmt"
	"os"
	"strings"
)

type orbiter struct {
	name       string
	parentName string
	parent     *orbiter
}

func (o orbiter) String() string {
	if o.parent == nil {
		return fmt.Sprintf("%s -> root", o.name)
	} else {
		return fmt.Sprintf("%s -> parent: %s", o.name, o.parent.name)
	}
}

func (o orbiter) TraverseOrbits() int {
	if o.parent == nil {
		return 0
	} else {
		return o.parent.TraverseOrbits() + 1
	}
}

func CalculateTotalOrbits(orbiters map[string]*orbiter) int {
	total := 0
	for _, v := range orbiters {
		total += v.TraverseOrbits()
	}

	return total
}

func ProcessOrbitMap(orbitMap string) map[string]*orbiter {
	rootOrbiter := orbiter{
		"COM", "", nil,
	}

	orbiters := map[string]*orbiter{
		"COM": &rootOrbiter,
	}

	orbitMapEntries := strings.Split(orbitMap, "&")

	for _, orbitMapEntry := range orbitMapEntries {
		CreateOrbiter(orbitMapEntry, orbiters)
	}

	// link up to parents
	for _, v := range orbiters {
		if v.name == "COM" {
			continue
		}

		v.parent = orbiters[v.parentName]
	}

	return orbiters
}

func CreateOrbiter(orbitMapEntry string, orbiters map[string]*orbiter) {
	parts := strings.Split(orbitMapEntry, ")")

	newOrbiter := orbiter{parts[1], parts[0], nil}

	orbiters[newOrbiter.name] = &newOrbiter
}

func FindNameOfCommonAncestor(orbiters map[string]*orbiter, o1Name string, o2Name string) string {
	var o1List []string

	o1Pointer := orbiters[o1Name]

	for o1Pointer.parent != nil {
		o1List = append(o1List, o1Pointer.name)
		o1Pointer = o1Pointer.parent
	}

	o1List = reverse(o1List)

	var o2List []string

	o2Pointer := orbiters[o2Name]

	for o2Pointer.parent != nil {
		o2List = append(o2List, o2Pointer.name)
		o2Pointer = o2Pointer.parent
	}

	o2List = reverse(o2List)

	prevPoint := o1List[0]
	for i := 1; o1List[i] == o2List[i]; i++ {
		prevPoint = o1List[i]
	}

	return prevPoint
}

func DistanceBetween(orbiters map[string]*orbiter, o1Name string, o2Name string) int {
	o1Dist := orbiters[o1Name].TraverseOrbits() - 1
	o2Dist := orbiters[o2Name].TraverseOrbits() - 1

	ancestorName := FindNameOfCommonAncestor(orbiters, o1Name, o2Name)
	ancestorDistance := orbiters[ancestorName].TraverseOrbits()

	return o1Dist + o2Dist - 2*ancestorDistance
}

func reverse(numbers []string) []string {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

func main() {
	travelDistance := DistanceBetween(ProcessOrbitMap(os.Args[1]), "YOU", "SAN")

	fmt.Printf("%d\n", travelDistance)
}
