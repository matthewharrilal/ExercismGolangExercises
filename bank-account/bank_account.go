package account

import "sync"

// Creating account structure
type Account struct {
	open    bool       // Denotes the status of the bank account
	balance int64      // Current balance of the bank account
	mutex   sync.Mutex // Aquiring lock to concurrent transactions or that morre than one thread doesn't have access to a bank account at once
}

// Opens bank account takes in an initial deposit and returns the value for the pointer pointing to account
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 { // No bank account for you if you don't put in that money
		return nil
	}
	return &Account{open: true, balance: initialDeposit, mutex: sync.Mutex{}} // opens the bank account and passes a lock with the account to perform operations
}

// Closes the account and returns a payout and a boolean determining if it was a succesful operation
func (a *Account) Close() (payout int64, ok bool) {
	a.mutex.Lock() // Aquire a lock on the bank account so only one thread has access to this shared resource
	defer a.mutex.Unlock()
	if !a.open {
		return 0, false
	}
	currentBalance := a.balance
	a.open = false
	a.balance = 0
	return currentBalance, true
}

// Balance returns the current balance
func (a *Account) Balance() (balance int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if !a.open {
		return 0, false
	}
	return a.balance, true
}

// Deposit deposit or withdraws a given amount from the account (if the resulting balance is not negative)
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if !a.open || a.balance+amount < 0 {
		return 0, false
	}
	a.balance += amount
	return a.balance, true
}
