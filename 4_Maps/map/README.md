# Maps

![maps](img\maps.png)



## Declare a map

```go
// First version:
colors := map[string]string{
	"red":   "#ff0000",
	"green": "#00ff00",
}

// Second version
var colors map[string]string

// Third version
colors := make(map[string]string)


```



## Insert values 

```go
colors["red"] = "#ff0000"
```



## Delete values

```go
delete(colors, "red")
```



## Iterating over a map

```go
func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println("Color:", color, "Hex:", hex)
	}
}
```



## Differences between Maps and Structs

![maps_vs_struct](img\map_vs_struct.png)



