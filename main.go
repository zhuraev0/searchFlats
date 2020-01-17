package main

import "sort"

type flat struct {
	id                                 int64
	name                               string
	stateOfTheHouse                    string
	price                              int64
	currency                           string
	flatRoomCount                      int64
	town                               string
	radiusFromTheLocationOfFinderMetre int64
}

func sortByPriceAsc(flats []flat) []flat {
	result := make([]flat, len(flats))
	copy(result, flats)
	sort.Slice(result, func(i, j int) bool {
		return result[i].price < result[j].price
	})
	return result
}

func sortByPriceDesc(flats []flat) []flat {
	result := make([]flat, len(flats))
	copy(result, flats)
	sort.Slice(result, func(i, j int) bool {
		return result[i].price > result[j].price
	})
	return result
}

func searchByPrice(flats []flat, limit int64) []flat {
	result := make([]flat, 0)
	for _, flat := range flats {
		if flat.price <= limit {
			result = append(result, flat)
		}
	}
	return result
}

func searchByMaxPrice(flats []flat, maxPrice int64) []flat {
	result := make([]flat, 0)
	for _, flat := range flats {
		if flat.price <= maxPrice {
			result = append(result, flat)
		}
	}
	return result
}

func searchFlatsByTown(flats []flat, town string) []flat {
	result := make([]flat, 0)
	for _, flat := range flats {
		if flat.town == town {
			result = append(result, flat)
		}
	}
	return result
}

func searchFlatsByTowns(flats []flat, towns []string) []flat {
	result := make([]flat, 0)
	for _, flat := range flats {
		for _, town := range towns {
			if flat.town == town {
				result = append(result, flat)
				continue
			}
		}
	}
	return result
}

func searchFlatsOnRadius2500Metres(flats []flat, radiusMetre int64) []flat {
	result := make([]flat, 0)
	for _, flat := range flats {
		if flat.radiusFromTheLocationOfFinderMetre <= radiusMetre {
			result = append(result, flat)
		}
	}
	return result
}

func searchFlatsByStateOfTheHouse(flats []flat, stateOfTheHouse string) []flat {
	result := make([]flat, 0)
	for _, flat := range flats {
		if flat.stateOfTheHouse == stateOfTheHouse {
			result = append(result, flat)
		}
	}
	return result
}

func searchByMinAndMaxPrice(flats []flat, min int64, max int64) []flat {
	result := make([]flat, 0)
	for _, flat := range flats {
		if flat.price >= min && flat.price <= max {
			result = append(result, flat)
		}
	}
	return result
}

func searchAllFlatsByFlatRoomCount(flats []flat, FlatRoomCount int64) []flat {
	result := make([]flat, 0)
	for _, flat := range flats {
		if flat.flatRoomCount == FlatRoomCount {
			result = append(result, flat)
		}
	}
	return result
}
func main() {

}