package main

import (
	"encoding/json"
	"log"
	"fmt"
)

type Movie struct{
	Title string
	Year  int `json:"released"`
	Color bool `json:color,omitempty`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	// ...
}

func main() {
	{
		data, err := json.Marshal(movies)
		if err != nil{
			log.Fatalf("JSON marshaling failed: %s", err)
		}

		fmt.Printf("%s\n",data)
	}

	{
		data, err := json.MarshalIndent(movies, "","    ")
		if err != nil{
			log.Fatalf("JSON marshaling failed: %s", err)
		}
		fmt.Printf("%s\n",data)

		//-----------------------------------------------------
		var titles []struct{Title string}
		if err := json.Unmarshal(data,&titles); err != nil{
			log.Fatal("JSON unmarshaling failed: %s", err)
		}
		fmt.Println(titles)
		fmt.Printf("%q \n",titles)

		//----------------------------------------------------
		var years []struct{Year  int `json:"released"`}
		if err := json.Unmarshal(data,&years); err != nil{
			log.Fatal("JSON unmarshaling failed: %s", err)
		}
		fmt.Println(years)
		fmt.Printf("%v \n",years)


		//----------------------------------------------------
		var actors []struct{Actors []string}
		if err := json.Unmarshal(data,&actors); err != nil{
			log.Fatal("JSON unmarshaling failed: %s", err)
		}
		fmt.Println(actors)
		fmt.Printf("%q \n",actors)

		//----------------------------------------------------
		str := "abcd"
		fmt.Println(str)
		fmt.Printf("%q",str)
	}
}
