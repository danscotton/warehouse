package main

import (
	"testing"

	gm "github.com/onsi/gomega"
)

func TestNewWarehousesHaveNoStock(t *testing.T) {
	g := gm.NewGomegaWithT(t)
	w := NewWarehouse()

	g.Expect(w.GetStock()).To(gm.HaveLen(0))
}

func TestTwoItemsThatAreTheSameCreateOneEntry(t *testing.T) {
	g := gm.NewGomegaWithT(t)
	w := NewWarehouse()

	w.AddStock(NewRecord("Random Access Memories", "Daft Punk"), 5)
	w.AddStock(NewRecord("Random Access Memories", "Daft Punk"), 5)

	g.Expect(w.GetStock()).To(gm.SatisfyAll(
		gm.HaveKeyWithValue(NewRecord("Random Access Memories", "Daft Punk"), 10),
		gm.HaveLen(1)))
}

func TestHasStock(t *testing.T) {
	g := gm.NewGomegaWithT(t)
	w := NewWarehouse()

	r1 := NewRecord("Random Access Memories", "Daft Punk")
	r2 := NewRecord("Chutes Too Narrow", "The Shins")

	w.AddStock(r1, 5)

	g.Expect(w.HasStock(r1)).To(gm.BeTrue())
	g.Expect(w.HasStock(r2)).To(gm.BeFalse())
}
