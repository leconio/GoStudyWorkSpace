package _map

import "fmt"

func init() {
	colors := map[string]string{}
	colors["red"] = "#f00"
	colors["white"] = "#fff"
	fmt.Println(colors)
	value, exists := colors["red"]
	if exists {
		fmt.Println(value)
	}

	testMap1()
	testMap2()
}

func testMap1() {
	colors := map[string]string{"AliceBlue": "#f0f8ff", "Coral": "#ff7F50", "DarkGray": "#a9a9a9", "ForestGreen": "#228b22",}
	for key, value := range colors {
		fmt.Println(key,value)
	}
}

func testMap2() {
	colors := map[string]string{"AliceBl|ue": "#f0f8ff", "Coral": "#ff7F50", "DarkGray": "#a9a9a9", "ForestGreen": "#228b22",}
	delete(colors,"Coral")
	fmt.Println(colors)
}
