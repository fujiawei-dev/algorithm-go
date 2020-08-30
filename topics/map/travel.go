package main

import "fmt"

func main() {
	tickets := map[int]int{1: 2, 3: 4, 2: 3}
	//tickets := map[int]int{2: 1, 4: 3, 3: 2}
	PrintTravel(tickets)
}

func RevTickets(tickets map[int]int) map[int]int {
	revTickets := make(map[int]int)
	for k, v := range tickets {
		revTickets[v] = k
	}
	return revTickets
}

func PrintTravel(tickets map[int]int) {
	revTickets := RevTickets(tickets)
	var next int
	for k := range tickets {
		_, ok := revTickets[k]
		if !ok {
			next = k
			break
		}
	}
	fmt.Print(next, " ")
	for range tickets {
		next = tickets[next]
		fmt.Print(next, " ")
	}
}
