package main

import (
	"fmt"
	"html/template"
	"net/http"
	"reflect"
	"strconv"
)

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("input.html")
		t.Execute(w, nil)
	}
	lat1 := r.FormValue("latitude")
	lon1 := r.FormValue("longitude")
	pin1 := r.FormValue("pincode")

	lat, _ := strconv.ParseFloat(lat1, 7)
	lon, _ := strconv.ParseFloat(lon1, 7)
	pin, _ := strconv.ParseInt(pin1, 10, 24)

	fmt.Println("lat : ", lat)
	fmt.Println("lon : ", lon, reflect.TypeOf(lon))
	fmt.Println("pin : ", pin)

	ne_lat := 24.6541753
	ne_lon := 73.7612527

	sw_lat := 24.4635105
	sw_lon := 73.6436533

	if ne_lat > lat && ne_lon > lon && sw_lat < lat && sw_lon < lon && pin == 313001 {

		fmt.Println("correct")
		fmt.Fprintf(w, "correct")
	} else {
		fmt.Println("incorrect")
		fmt.Fprintf(w, "incorrect")
	}
}

func main() {

	http.HandleFunc("/", login)

	http.ListenAndServe(":8080", nil)
}
