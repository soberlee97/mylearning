package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var mapStudents = make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		mapStudents[key] = value
	}
	fmt.Println(mapStudents)

	var keys = make([]string, 0, 200)
	for key := range mapStudents {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	fmt.Println(keys)
	for _, key := range keys {
		fmt.Println(key, mapStudents[key])
	}

}
