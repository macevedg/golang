package main

import (
	//"database/sql"
	"fmt"

	"github.com/tkuchiki/go-timezone"
	//"strconv"
)

func main() {
	//i, err := strconv.Atoi("lololololololololol")
	//num := sql.NullInt64{Int64: int64(i), Valid: err == nil}

	//fmt.Printf("number: %v %d", num, i)

	location2 := "America/Sao_Paulo"
	location3 := "America/New_York"
	//location, _ := time.LoadLocation(location2)

	//zone, _ := time.Now().In(location).Zone()

	tz := timezone.New()

	sao, _ := tz.GetTimezoneAbbreviation(location2, false)
	ny, _ := tz.GetTimezoneAbbreviation(location3, false)

	//fmt.Println(abbr)
	// Get the time zone name
	//zoneName, err := monday.ParseInLocation
	//ratezone.Name(location, currentTime)

	fmt.Println(location2, sao)
	fmt.Println(location3, ny)
}

func main11() {
	/*
		words := make([]string,8)
		words[0] = "hello world"
		words[1] = "hello world"

		fmt.Println(len(words))


		var i int8 = 120

		i +=10

		fmt.Println(i)
	*/

	go fmt.Println("1")
	fmt.Println("2")

	ch := make(chan int)

	close(ch)

	val := <-ch

	fmt.Println(val)

	var stocks map[string]float64

	price := stocks["MSFT"]

	fmt.Println(price)
}
