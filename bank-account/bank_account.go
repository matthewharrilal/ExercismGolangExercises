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

	defer a.mutex.Unlock() // Release the lock only until the surrounding function has executed
	if !a.open {           // If the bank account isn't open then there is no payout and return false meaning that you cant close what is already closed
		return 0, false
	}
	currentBalance := a.balance
	a.open = false // Payout with the current balance set open status to false and return the current balance
	a.balance = 0
	return currentBalance, true
}

// Extending the account structure with this balance method and returns a balance and boolean determining status of operation
func (a *Account) Balance() (balance int64, ok bool) {
	a.mutex.Lock()         // Aquire the lock on the shared resource ... the balance
	defer a.mutex.Unlock() // Release lock once surrounding function has been executed
	if !a.open {           // If the bank account isnt open return false
		return 0, false
	}
	return a.balance, true
}

// Extending the account structure with the deposit mehtod takes in an initial amoutnand returns the new balance and whether or not operation was successful
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.mutex.Lock()         // Aquire the lock on the shared resource the balance
	defer a.mutex.Unlock() // Release lock when surrouding function has executed
	if !a.open || a.balance+amount < 0 {
		return 0, false // If the bank account is not open or even with the deposit the balance is still negativr
	}
	a.balance += amount
	return a.balance, true
}
