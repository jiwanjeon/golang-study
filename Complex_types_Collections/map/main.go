package main

import (
	"fmt"
	"sort"
)

// - powerful & ingenious(정교한) & versatile(다재다능한)
// - collection of unordered pairs of key-value
// - fast lookups and values that can retrieve, update or delete with the help of keys
// - A.K.A hash map, hash table, unordered map, dictionary, associative array
// - the type of keys and type of values must be the same type

func main() {
	states := make(map[string]string) //map[key_type]value_type or map[key_type]value_type, initial_capacity
	fmt.Println(states)

	states["WA"] = "Washington"
	states["OR"] = "Oregon"
	states["CA"] = "California"
	fmt.Println(states)

	california := states["CA"]
	fmt.Println(california)

	//How to Delete ==> delete(map_name, key)
	delete(states, "OR")

	//How to add key-value in the map ==> map_name[key] = value
	states["NY"] = "New York"
	fmt.Println(states)

	for k, v := range states {
		fmt.Printf("%v: %v", k, v)
	}

	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	fmt.Println(keys)
	sort.Strings(keys)
	fmt.Println(keys)

	for i := range keys {
		fmt.Println(states[keys[i]])
	}

	//Modifying map :  As we know that maps are of reference type. So, when we assign an existing map to a new variable,
	//both the maps still refer to the same underlying data structure. So, when we update one map it will reflect in another map.
	// Creating and initializing a map
	m_a_p := map[int]string{
		90: "Dog",
		91: "Cat",
		92: "Cow",
		93: "Bird",
		94: "Rabbit",
	}
	fmt.Println("Original map: ", m_a_p)

	// Assigned the map into a new variable
	new_map := m_a_p

	// Perform modification in new_map
	new_map[96] = "Parrot"
	new_map[98] = "Pig"

	// Display after modification
	fmt.Println("New map: ", new_map)
	fmt.Println("\nModification done in old map:\n", m_a_p)
}
