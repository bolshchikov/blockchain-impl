package blockchain

import "testing"

func TestBlockchainCreation(t *testing.T) {
	bc := New()
	bc.AddBlock("test")
	if bc.Len() != 2 {
		t.Fatalf("bc.Len() == %d, want %d ", bc.Len(), 2)
	}
	if IsValidChain(bc) != true {
		t.Fatalf("IsValidChain(bc) == %v, want true", IsValidChain(bc))
	}
}

func TestBlockchainDataTempering(t *testing.T) {
	bc := New()
	bc.AddBlock("test")
	block := bc.GetLastBlock()
	block.Data = "test2"
	if IsValidChain(bc) != false {
		t.Fatalf("IsValidChain(bc) == %v, want false", IsValidChain(bc))
	}
}

func TestBlockchainIndexTempering(t *testing.T) {
	bc := New()
	bc.AddBlock("test")
	block := bc.GetLastBlock()
	block.Index = 100
	if IsValidChain(bc) != false {
		t.Fatalf("IsValidChain(bc) == %v, want false", IsValidChain(bc))
	}
}
