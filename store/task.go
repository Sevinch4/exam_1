package store

import (
	"fmt"
	"math"
)

type Store struct {
	Data []Costumer `json:"costumer"`
}
type Costumer struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Cash      int    `json:"cash"`
	Baskets   Basket `json:"basket"`
}
type Basket struct {
	ID       string `json:"id"`
	Products []Product
	Total    int `json:"total"`
}
type Product struct {
	ID       string `json:"id"`
	Category string `json:"category"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}

//task-1 barcha customerlarni umumiy mablagini va qancha summa harid qilganini hisonlang va korsating
/*
	umumiy mablag: 150,000 +120,00+...=470,000
	qancha summa harid qilgan : 59,000 + 42,000 + 85,000 = 186,000
*/
func (s *Store) Task1() {
	fmt.Println("-----------------------------------task-1--------------------------------------------------")
	costumers := s.Data
	//fmt.Println(costumers)
	countMap := make(map[string]int)
	basketMap := make(map[string]int)
	sumUmumiy := 0
	sumHarid := 0

	for _, c := range costumers {
		countMap[c.ID] = c.Cash
		basketMap[c.Baskets.ID] = c.Baskets.Total
	}
	for _, s := range costumers {
		sumUmumiy += countMap[s.ID]
		sumHarid += basketMap[s.Baskets.ID]
	}
	fmt.Printf("Umumiy mablag: %d\nQancha summa harid qilgan: %d\n", sumUmumiy, sumHarid)
}

//task-2 eng kop pul sarflagan customerni aniqlang va ko'rsating
/*
	eng kop pul sarf etgan mijoz: Micheal Jordan(200,000 sum)
*/
func (s *Store) Task2() {
	fmt.Println("-----------------------------------task 2--------------------------------------------------")
	clients := s.Data
	max := 0
	costumerMap := make(map[string]int)
	costumerName := make(map[string]string)
	costumerLastName := make(map[string]string)

	for _, costumer := range clients {
		costumerMap[costumer.ID] = costumer.Baskets.Total
		costumerName[costumer.ID] = costumer.FirstName
		costumerLastName[costumer.ID] = costumer.LastName
	}
	for _, c := range clients {
		if costumerMap[c.ID] > max {
			max = costumerMap[c.ID]
		}
	}
	for c := range costumerMap {
		if max == costumerMap[c] {
			fmt.Printf("Eng ko'p pul sarf etgan mijoz: %s %s (%d sum)\n", costumerName[c], costumerLastName[c], max)
		}
	}
}

// task-3 eng qimmat mahsulot steak(30,000 sum)
func (s *Store) Task3() {
	fmt.Println("-----------------------------------task-3--------------------------------------------------")
	data := s.Data
	maxProduct := 0
	productName := make(map[string]string)
	countMap := make(map[string]int)

	for _, c := range data {
		for _, b := range c.Baskets.Products {
			countMap[b.ID] = b.Price
			productName[b.ID] = b.Name
		}
	}
	for _, p := range data {
		for _, b := range p.Baskets.Products {
			if countMap[b.ID] > maxProduct {
				maxProduct = countMap[b.ID]
			}
		}
	}
	for c := range countMap {
		if maxProduct == countMap[c] {
			fmt.Printf("Eng qimmat mahsulot: %s (%d sum)\n", productName[c], maxProduct)
		}
	}
}

// task4 barcha productlar un ortacha narxni hisonlang va korsating -> (12+4+5)/3 = ....
func (s *Store) Task4() {
	fmt.Println("-----------------------------------task-4--------------------------------------------------")
	data := s.Data
	counter := 0
	sum := 0
	priceMap := make(map[string]int)

	for _, c := range data {
		for _, p := range c.Baskets.Products {
			priceMap[p.ID] = p.Price
		}
	}
	for _, c := range data {
		for _, p := range c.Baskets.Products {
			sum += priceMap[p.ID]
			counter++
		}
	}
	fmt.Println("Ortacha narx: ", float32(sum)/float32(counter))

}

// task5 eng arzon savdo qilgan customerni aniqlang va korsating
func (s *Store) Task5() {
	fmt.Println("-----------------------------------task-5--------------------------------------------------")
	clients := s.Data
	costumerMap := make(map[string]int)
	costumerName := make(map[string]string)
	costumerLastName := make(map[string]string)

	for _, costumer := range clients {
		costumerMap[costumer.ID] = costumer.Baskets.Total
		costumerName[costumer.ID] = costumer.FirstName
		costumerLastName[costumer.ID] = costumer.LastName
	}
	min := math.MaxInt32
	for _, c := range clients {
		if costumerMap[c.ID] < min {
			min = costumerMap[c.ID]
		}
	}
	for c := range costumerMap {
		if min == costumerMap[c] {
			fmt.Printf("Eng kam pul sarf etgan mijoz: %s %s(%d sum)\n", costumerName[c], costumerLastName[c], min)
		}
	}
}

// task6 end ko'p sotilgan productlar categoriyasini aniqlang va chiqaring
func (s *Store) Task6() {
	fmt.Println("-----------------------------------task-6--------------------------------------------------")
	data := s.Data
	category := make(map[string]string)
	quantity := make(map[string]int)
	maxSale := 0

	for _, c := range data {
		for _, p := range c.Baskets.Products {
			category[p.ID] = p.Category
			quantity[p.ID] = p.Quantity
		}
	}
	for _, c := range data {
		for _, p := range c.Baskets.Products {
			if quantity[p.ID] > maxSale {
				maxSale = quantity[p.ID]
			}
		}
	}
	for p := range quantity {
		if maxSale == quantity[p] {
			fmt.Println("Eng ko'p sotilgan categoriya:", category[p])
		}
	}

}

//task7 eng kop va eng kam sotilgan productlar nomini chiqarish
/*
	eng kop sotilgan mahsulot nomi: carrot (micheal jordan - 5ta)
	eng kam sotilgan mahsulot nomi: apple (dwayne johnson - 1ta)
*/
func (s *Store) Task7() {
	fmt.Println("-----------------------------------task-7--------------------------------------------------")
	data := s.Data

	quantity := make(map[string]int)

	min := math.MaxInt32
	max := math.MinInt32

	clientName := make(map[string]string)
	clientLastName := make(map[string]string)
	productName := make(map[string]string)

	for _, d := range data {
		for _, b := range d.Baskets.Products {
			quantity[b.ID] = b.Quantity
			productName[b.ID] = b.Name
			clientName[b.ID] = d.FirstName
			clientLastName[b.ID] = d.LastName
		}
	}
	for _, d := range data {
		for _, b := range d.Baskets.Products {
			if max < quantity[b.ID] {
				max = quantity[b.ID]
			}
			if min > quantity[b.ID] {
				min = quantity[b.ID]
			}
		}
	}
	for i := range quantity {
		if max == quantity[i] {
			fmt.Printf("Eng kop sotilgan mahsulot nomi:%s (%s %s - %dta)\n", productName[i], clientName[i], clientLastName[i], max)
		}
	}
	for i := range quantity {
		count := 0
		if min == quantity[i] {
			count++
			if count == 1 {
				fmt.Printf("Eng kam sotilgan mahsulot nomi:%s (%s %s - %dta)\n", productName[i], clientName[i], clientLastName[i], min)
				break
			}
		}

	}

}

// task8 har bir savdo ucun ortacha mahsulot miqdorini hisoblang
func (s *Store) Task8() {
	fmt.Println("-----------------------------------task-8--------------------------------------------------")
	data := s.Data
	counter := 0
	sum := 0
	priceMap := make(map[string]int)
	for _, c := range data {
		for _, p := range c.Baskets.Products {
			priceMap[p.ID] = p.Quantity
		}
	}
	for _, c := range data {
		for _, p := range c.Baskets.Products {
			sum += priceMap[p.ID]
			counter++
		}
	}
	fmt.Printf("Ortacha : %.3f\n", float32(sum)/float32(counter))
}

//task9 eng kop mahsulot mahsulot sotb olgan mijozni aniqlang
/*
	eng kop mahsulot mahsulot sotb olgan mijoz: Micheal Jordan (8ta mahsulot),Serena Williams (C006): 11 items
*/
func (s *Store) Task9() {
	fmt.Println("-----------------------------------task-9--------------------------------------------------")
	data := s.Data
	firstName := make(map[string]string)
	lastName := make(map[string]string)
	productCount := make(map[string]int)
	for _, d := range data {
		basket := d.Baskets.Products

		firstName[d.ID] = d.FirstName
		lastName[d.ID] = d.LastName
		var total int
		for _, b := range basket {
			total += b.Quantity
		}
		productCount[d.ID] = total
	}
	max := 0
	for p := range productCount {
		if productCount[p] > max {
			max = productCount[p]
		}
	}
	for p := range productCount {
		if productCount[p] == max {
			fmt.Printf("Eng kop mansulot sotib olgan mijoz:%s %s (%dta mahsulot)\n", firstName[p], lastName[p], max)
		}
	}
}

//task10 eng kop savdolarda korinadigan mahsulotni aniqlang va korsating
/*
	eng kop savdoda koringan mahsulot : Carrot (Micheal Jordan - 5 ta)
*/
func (s *Store) Task10() {
	fmt.Println("-----------------------------------task-10--------------------------------------------------")
	data := s.Data
	var (
		firstName   string
		lastName    string
		productName string
		max         int = 0
	)
	for _, d := range data {
		for _, b := range d.Baskets.Products {
			if max < b.Quantity {

				max = b.Quantity

				firstName = d.FirstName
				lastName = d.LastName
				productName = b.Name
			}
		}
	}
	fmt.Printf("Eng kop savdoda koringan mahsulot:%s (%s %s - %dta)\n", productName, firstName, lastName, max)
}

//task11 ortacha savdo mablagi boyicha eng kop mablag'ga ega bolgan customerni aniqlang
/*
	ortacha savdo mablagi : (59,000+42,000+85,000) = 62 000
	eng kop mablag ega balgan mijoz: Micheal Jordan
*/
func (s *Store) Task11() {
	fmt.Println("-----------------------------------task-11--------------------------------------------------")
	client := s.Data
	totalSum := 0
	count := 0

	for _, c := range client {
		totalSum += c.Baskets.Total
		count++
	}
	fmt.Println("Ortacha savdo mablag'i: ", float32(totalSum)/float32(count))

	var (
		lastName  string
		firstName string
		max       int = 0
	)
	for _, c := range client {
		if max < c.Cash {
			max = c.Cash

			lastName = c.LastName
			firstName = c.FirstName
		}

	}
	fmt.Printf("Eng kop mablag'ga ega bo'lgan mijoz:%s %s\n", lastName, firstName)
}

//task12 eng kop daromad (quantity*price) qilgan toifani aniqlang va korsating
/*
	eng kop daromad qilgan toifa: snack (chips - 8,000*4 32,00sum)
*/
func (s *Store) Task12() {
	fmt.Println("-----------------------------------task-12--------------------------------------------------")
	data := s.Data
	max := 0
	total := make(map[string]int)
	category := make(map[string]string)
	productName := make(map[string]string)

	for _, d := range data {
		for _, b := range d.Baskets.Products {
			total[b.ID] = b.Quantity * b.Price

			category[b.ID] = b.Category
			productName[b.ID] = b.Name
		}

	}
	for c := range total {
		if max < total[c] {
			max = total[c]
		}
	}
	for t := range total {
		if max == total[t] {
			fmt.Printf("Eng kop daromad qilgan toifa: %s (%s - %dsum)\n", category[t], productName[t], total[t])
		}
	}
}

//task13 eng qimmat productni o'z ichiga olgan savdoni aniqlang va korsating:
/*
	eng qimmat mahsulot oz ichiga olgan savdo: steak (micheal jordan - 30*1 = 30ming)
*/
func (s *Store) Task13() {
	fmt.Println("-----------------------------------task-13--------------------------------------------------")
	var (
		data        = s.Data
		firstName   string
		lastName    string
		max         = 0
		productName string
		sum         = 0
	)
	for _, d := range data {
		for _, b := range d.Baskets.Products {
			if max < b.Price {

				max = b.Price
				sum = b.Quantity * max

				firstName = d.FirstName
				lastName = d.LastName
				productName = b.Name
			}
		}

	}
	fmt.Printf("Eng qimmat mahsulot oz ichiga olgan savdo: %s (%s %s - %dsum)\n", productName, firstName, lastName, sum)
}

//task14 har bir mijoz un eng kop xarid qilgan toifani aniqlang
/*
	eng kop harid qilgan toifa: snack(emma watson - 8,000*4 + 5,000*2 = 42,000 sum)
*/
func (s *Store) Task14() {
	fmt.Println("-----------------------------------task-14--------------------------------------------------")
	client := s.Data
	for _, c := range client {
		total := 0
		category := ""

		for _, b := range c.Baskets.Products {
			total += b.Quantity * b.Price
			category = b.Category
		}

		fmt.Printf("Eng kop harid qilgan toifa: %s (%s %s - %dsum)\n", category, c.FirstName, c.LastName, total)
		fmt.Println("------------------------------------------------------------------------------------")
	}
}

//task15 har bitda productdan savdo paytida umumiy nechta sotilgannin aniqlagn va korsating
/*
	umumiy nechta : 2+3+4+2+1...=...ta
*/
func (s *Store) Task15() {
	fmt.Println("-----------------------------------task-15--------------------------------------------------")
	count := 0
	data := s.Data

	for _, d := range data {
		for _, b := range d.Baskets.Products {
			count += b.Quantity
		}
	}

	fmt.Println("Umumiy nechta:", count, "ta")
}
