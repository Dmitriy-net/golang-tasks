//go:build !solution

package hotelbusiness

import (
	"sort"
)

type Guest struct {
	CheckInDate  int
	CheckOutDate int
}

type Load struct {
	StartDate  int
	GuestCount int
}

func ComputeLoad(guests []Guest) []Load {
	changes := make(map[int]int)

	for _, guest := range guests {
		changes[guest.CheckInDate] += 1
		changes[guest.CheckOutDate] -= 1
	}

	var dates []int
	for date := range changes {
		dates = append(dates, date)
	}
	sort.Ints(dates)

	var loads []Load
	currentGuests := 0

	for _, date := range dates {
		currentGuests += changes[date]
		if len(loads) == 0 || loads[len(loads)-1].GuestCount != currentGuests {
			loads = append(loads, Load{StartDate: date, GuestCount: currentGuests})
		}
	}

	return loads
}
