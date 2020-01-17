package main

import "fmt"

var flats = []flat{
	{
		id:                                 1,
		name:                               "Сдам",
		stateOfTheHouse:                    "Люкс",
		price:                              150_000,
		currency:                           "TJS",
		flatRoomCount:                      3,
		town:                               "Душанбе",
		radiusFromTheLocationOfFinderMetre: 9700,
	},
	{
		id:                                 2,
		name:                               "Продам",
		stateOfTheHouse:                    "Люкс",
		price:                              230_000,
		currency:                           "TJS",
		flatRoomCount:                      2,
		town:                               "Душанбе",
		radiusFromTheLocationOfFinderMetre: 2400,
	},
	{
		id:                                 3,
		name:                               "Продам",
		stateOfTheHouse:                    "Недавно ремонтировано",
		price:                              552_900,
		currency:                           "TJS",
		flatRoomCount:                      2,
		town:                               "Душанбе",
		radiusFromTheLocationOfFinderMetre: 6560,
	},
	{
		id:                                 4,
		name:                               "Продам",
		stateOfTheHouse:                    "Есть мебели, стиральная машина, отремонтирован",
		price:                              1_550_000,
		currency:                           "TJS",
		flatRoomCount:                      5,
		town:                               "Душанбе",
		radiusFromTheLocationOfFinderMetre: 3500,
	},
	{
		id:                                 5,
		name:                               "Сдам",
		stateOfTheHouse:                    "Люкс",
		price:                              23_000,
		currency:                           "TJS",
		flatRoomCount:                      1,
		town:                               "Душанбе",
		radiusFromTheLocationOfFinderMetre: 2_500,
	},
	{
		id:                                 1,
		name:                               "Продам",
		stateOfTheHouse:                    "Люкс",
		price:                              892_000,
		currency:                           "TJS",
		flatRoomCount:                      4,
		town:                               "Худжанд",
		radiusFromTheLocationOfFinderMetre: 520_000,
	},
}

func ExampleSortByPriceAsc() {
	ascSorted := sortByPriceAsc(flats)
	fmt.Println(ascSorted)
	// Output: [{5 Сдам Люкс 23000 TJS 1 Душанбе 2500} {1 Сдам Люкс 150000 TJS 3 Душанбе 9700} {2 Продам Люкс 230000 TJS 2 Душанбе 2400} {3 Продам Недавно ремонтировано 552900 TJS 2 Душанбе 6560} {1 Продам Люкс 892000 TJS 4 Худжанд 520000} {4 Продам Есть мебели, стиральная машина, отремонтирован 1550000 TJS 5 Душанбе 3500}]
}

func ExampleSortByPriceDesc() {
	ascSorted := sortByPriceDesc(flats)
	fmt.Println(ascSorted)
	// Output: [{4 Продам Есть мебели, стиральная машина, отремонтирован 1550000 TJS 5 Душанбе 3500} {1 Продам Люкс 892000 TJS 4 Худжанд 520000} {3 Продам Недавно ремонтировано 552900 TJS 2 Душанбе 6560} {2 Продам Люкс 230000 TJS 2 Душанбе 2400} {1 Сдам Люкс 150000 TJS 3 Душанбе 9700} {5 Сдам Люкс 23000 TJS 1 Душанбе 2500}]
}

func ExampleSearchByPrice_NoResults() {
	noResult := searchByPrice(flats, 20_000)
	fmt.Println(noResult)
	// Output: []
}

func ExampleSearchByPrice_MoreResult() {
	moreResult := searchByPrice(flats, 600_000)
	fmt.Println(moreResult)
	// Output: [{1 Сдам Люкс 150000 TJS 3 Душанбе 9700} {2 Продам Люкс 230000 TJS 2 Душанбе 2400} {3 Продам Недавно ремонтировано 552900 TJS 2 Душанбе 6560} {5 Сдам Люкс 23000 TJS 1 Душанбе 2500}]
}

func ExampleSearchByMaxPrice() {
	maxPrice := searchByMaxPrice(flats, 800_000)
	fmt.Println(maxPrice)
	// Output: [{1 Сдам Люкс 150000 TJS 3 Душанбе 9700} {2 Продам Люкс 230000 TJS 2 Душанбе 2400} {3 Продам Недавно ремонтировано 552900 TJS 2 Душанбе 6560} {5 Сдам Люкс 23000 TJS 1 Душанбе 2500}]
}

func ExampleSearchByTown_NoResults() {
	noResult := searchFlatsByTown(flats, "Кайраккум")
	fmt.Println(noResult)
	// Output: []
}

func ExampleSearchByDistrict_TwoOrMoreResults() {
	result := searchFlatsByTown(flats, "Душанбе")
	fmt.Println(result)
	// Output: [{1 Сдам Люкс 150000 TJS 3 Душанбе 9700} {2 Продам Люкс 230000 TJS 2 Душанбе 2400} {3 Продам Недавно ремонтировано 552900 TJS 2 Душанбе 6560} {4 Продам Есть мебели, стиральная машина, отремонтирован 1550000 TJS 5 Душанбе 3500} {5 Сдам Люкс 23000 TJS 1 Душанбе 2500}]
}

func ExampleSearchByDistricts_NoResults() {
	result := searchFlatsByTowns(flats, []string{"Худжанд", "Душанбе"})
	fmt.Println(result)
	// Output: [{1 Сдам Люкс 150000 TJS 3 Душанбе 9700} {2 Продам Люкс 230000 TJS 2 Душанбе 2400} {3 Продам Недавно ремонтировано 552900 TJS 2 Душанбе 6560} {4 Продам Есть мебели, стиральная машина, отремонтирован 1550000 TJS 5 Душанбе 3500} {5 Сдам Люкс 23000 TJS 1 Душанбе 2500} {1 Продам Люкс 892000 TJS 4 Худжанд 520000}]
}

func ExampleSearchFlatsOnRadius2500Metres() {
	radiusMetre := searchFlatsOnRadius2500Metres(flats, 3_500)
	fmt.Println(radiusMetre)
	// Output: [{2 Продам Люкс 230000 TJS 2 Душанбе 2400} {4 Продам Есть мебели, стиральная машина, отремонтирован 1550000 TJS 5 Душанбе 3500} {5 Сдам Люкс 23000 TJS 1 Душанбе 2500}]
}

func ExampleSearchFlatsByStateOfTheHouse() {
	stateOfFlat := searchFlatsByStateOfTheHouse(flats, "Люкс")
	fmt.Println(stateOfFlat)
	// Output: [{1 Сдам Люкс 150000 TJS 3 Душанбе 9700} {2 Продам Люкс 230000 TJS 2 Душанбе 2400} {5 Сдам Люкс 23000 TJS 1 Душанбе 2500} {1 Продам Люкс 892000 TJS 4 Худжанд 520000}]
}

func ExampleSearchAllFlatsByFlatRoomCount() {
	FlatRoomCount := searchAllFlatsByFlatRoomCount(flats, 2)
	fmt.Println(FlatRoomCount)
	// Output: [{2 Продам Люкс 230000 TJS 2 Душанбе 2400} {3 Продам Недавно ремонтировано 552900 TJS 2 Душанбе 6560}]
}

func ExampleSearchByMinAndMaxPrice_OneResult() {
	result := searchByMinAndMaxPrice(flats, 60_000, 900_500)
	fmt.Println(result)
	// Output: [{1 Сдам Люкс 150000 TJS 3 Душанбе 9700} {2 Продам Люкс 230000 TJS 2 Душанбе 2400} {3 Продам Недавно ремонтировано 552900 TJS 2 Душанбе 6560} {1 Продам Люкс 892000 TJS 4 Худжанд 520000}]
}
