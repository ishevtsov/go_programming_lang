// Package bank provides a concurrency-safe bank with one account.
package bank

var (
	sema    = make(chan struct{}, 1)
	balance int
)

func Deposit(amount int) {
	sema <- struct{}{} // acquire token
	balance = balance + amount
	<-sema // release token
}

func Balance() int {
	sema <- struct{}{}
	b := balance
	<-sema
	return b
}
