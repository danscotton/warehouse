package warehouse

import (
	"fmt"
	"log"
)

type Item interface {
	String() string
}

type Artist struct {
	Name string
}

type Record struct {
	Artist Artist
	Title  string
}

func (r Record) String() string {
	return fmt.Sprintf("%s – %s", r.Title, r.Artist)
}

type Author struct {
	Name string
}

type Book struct {
	Title string
	Author
}

func (b Book) String() string {
	return fmt.Sprintf("%s, %s", b.Title, b.Name)
}

type Warehouse struct {
	stock map[Item]int
}

func (w Warehouse) AddStock(s Item, c int) {
	w.stock[s] += c
}

func (w Warehouse) GetStock() map[Item]int {
	return w.stock
}

func (w Warehouse) HasStock(i Item) bool {
	_, found := w.stock[i]
	return found
}

func NewWarehouse() Warehouse {
	return Warehouse{
		stock: map[Item]int{},
	}
}

func NewRecord(title, artist string) Record {
	return Record{
		Title: title,
		Artist: Artist{
			Name: artist,
		},
	}
}

func NewBook(title, author string) Item {
	return Book{
		Title: title,
		Author: Author{
			Name: author,
		},
	}
}

func main() {
	w := NewWarehouse()

	w.AddStock(NewRecord("Random Access Memories", "Daft Punk"), 10)
	w.AddStock(NewRecord("Random Access Memories", "Daft Punk"), 5)
	w.AddStock(NewBook("Zen and the Art of Motorcycle Maintenence", "Robert Pirsig"), 5)

	log.Printf("%#v", w.GetStock())
}
