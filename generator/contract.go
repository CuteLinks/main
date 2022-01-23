//go:generate mockgen -source=contract.go -destination contract_mocks.go -package $GOPACKAGE
package generator

type randomizer interface {
	GetRandomInt(n int) int
}


