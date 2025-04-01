package ton

import "github.com/ice-hermes/tongo/tlb"

type Transaction struct {
	tlb.Transaction
	BlockID BlockIDExt
}
