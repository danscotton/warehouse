package warehouse

type IWarehouse interface {
	AddStock(i Record, c int)
}

type RecordLabel struct {
	Name string
}

func (rl RecordLabel) SupplyTo(w IWarehouse, stock map[Record]int) {
	for r, c := range stock {
		w.AddStock(r, c)
	}
}
