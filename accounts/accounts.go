package accounts

import "errors"

// Account struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("no money. you can't withraw")

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit method
func (receiver *Account) Deposit(amount int) {
	receiver.balance += amount
}

// Withraw method
func (receiver *Account) Withraw(amount int) error {
	if receiver.balance < amount {
		return errNoMoney
	}
	receiver.balance -= amount
	// 반환이 error이므로 아무 반환도 하고 싶지 않더라도 nil 반환해야 함
	return nil
}

// Balance show
func (receiver Account) Balance() int {
	return receiver.balance
}

// NewOwner func
func (receiver *Account) NewOwner(owner string) {
	receiver.owner = owner
}
