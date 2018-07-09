package warehouse

import (
	"testing"

	gm "github.com/onsi/gomega"
)

type StubWarehouse struct {
	Stock map[Record]int
}

func (sw StubWarehouse) AddStock(i Record, c int) {
	sw.Stock[i] = c
}

func TestRecordLabelsCanSupplyToWarehouses(t *testing.T) {
	g := gm.NewGomegaWithT(t)

	rl := RecordLabel{Name: "Sub Pop"}
	r1 := NewRecord("Chutes Too Narrow", "The Shins")
	r2 := NewRecord("Fleet Foxes", "Fleet Foxes")
	w := StubWarehouse{map[Record]int{}}

	rl.SupplyTo(w, map[Record]int{
		r1: 1000,
		r2: 500,
	})

	g.Expect(w.Stock).To(gm.HaveKeyWithValue(r1, 1000))
	g.Expect(w.Stock).To(gm.HaveKeyWithValue(r2, 500))
}
