package main

import "testing"

func TestDistanceBetween(t *testing.T) {
	travelDistance := DistanceBetween(ProcessOrbitMap("COM)B&B)C&C)D&D)E&E)F&B)G&G)H&D)I&E)J&J)K&K)L&K)YOU&I)SAN"), "YOU", "SAN")

	assertEqualsInts(t, travelDistance, 4)
}

func TestCalculateTotalOrbits(t *testing.T) {
	totalOrbits := CalculateTotalOrbits(ProcessOrbitMap("COM)B&B)C&C)D"))

	assertEqualsInts(t, 6, totalOrbits)

	totalOrbits = CalculateTotalOrbits(ProcessOrbitMap("COM)B&B)C&C)D&D)E&E)F&B)G&G)H&D)I&E)J&J)K&K)L"))

	assertEqualsInts(t, 42, totalOrbits)
}

func TestOrbiterString(t *testing.T) {
	root := orbiter{"COM", "", nil}
	child := orbiter{"B", "COM", &root}

	assertEquals(t, "B -> parent: COM", child.String())
	assertEquals(t, "COM -> root", root.String())
}

func TestProcessOrbitMap(t *testing.T) {
	orbiters := ProcessOrbitMap("COM)B&B)C&C)D")

	if _, ok := orbiters["B"]; !ok {
		t.Errorf("orbiters did not have key B")
	}

	if _, ok := orbiters["C"]; !ok {
		t.Errorf("orbiters did not have key C")
	}

	if _, ok := orbiters["D"]; !ok {
		t.Errorf("orbiters did not have key D")
	}
}

func TestCreateOrbiter(t *testing.T) {
	orbiters := make(map[string]*orbiter)

	CreateOrbiter("COM)B", orbiters)

	anOrbiter, ok := orbiters["B"]

	if !ok {
		t.Errorf("orbiters did not have key B")
	}

	assertEquals(t, "B", anOrbiter.name)
	assertEquals(t, "COM", anOrbiter.parentName)
}

func assertEqualsInts(t *testing.T, expected int, actual int) {
	if expected != actual {
		t.Errorf("got: %d wanted: %d", actual, expected)
	}
}

func assertEquals(t *testing.T, expected string, actual string) {
	if expected != actual {
		t.Errorf("got: %s wanted: %s", actual, expected)
	}
}
