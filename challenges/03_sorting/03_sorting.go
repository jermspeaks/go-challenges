package main

import (
	"fmt"
	"sort"
)

// Flight - a struct that
// contains information about flights
type Flight struct {
	Origin      string
	Destination string
	Price       int
}

type ByPrice []Flight

func (a ByPrice) Len() int           { return len(a) }
func (a ByPrice) Less(i, j int) bool { return a[i].Price < a[j].Price }
func (a ByPrice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// SortByPrice sorts flights from lower to highest
func SortByPrice(flights []Flight) []Flight {
	sort.Sort(ByPrice(flights))
	return flights
}

func printFlights(flights []Flight) {
	for _, flight := range flights {
		fmt.Printf("Origin: %s, Destination: %s, Price: %d", flight.Origin, flight.Destination, flight.Price)
	}
}

func main() {
	flights := []Flight{
		Flight{Price: 30},
		Flight{Price: 20},
		Flight{Price: 50},
		Flight{Price: 1000},
	}

	sortedList := SortByPrice(flights)
	printFlights(sortedList)
}
