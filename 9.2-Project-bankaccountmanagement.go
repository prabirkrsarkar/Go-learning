package main

import (
	"fmt"
)

type Account struct {
	AccountNumber string
	OwnerName     string
	Balance       float64
}

// We use Pointer receiver for the Deposit method because it modifies the state of the Account struct
// by updating the Balance field.
func (acc *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("Deposit amount must be positive. Account: %s", acc.AccountNumber)
	}
	acc.Balance += amount
	fmt.Printf("Deposited $%.2f to account %s. New balance: $%.2f\n", amount, acc.AccountNumber, acc.Balance)
	return nil
}

func (acc *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("Withdrawal amount must be positive. Account: %s", acc.AccountNumber)
	}
	if amount > acc.Balance {
		return fmt.Errorf("Insufficient funds. Account: %s", acc.AccountNumber)
	}
	acc.Balance -= amount
	fmt.Printf("Withdrew $%.2f from account %s. New balance: $%.2f\n", amount, acc.AccountNumber, acc.Balance)
	return nil
}

// Though we could have used a value receiver for the GetBalance method but using a pointer receiver
// for consistency.
func (acc *Account) GetBalance() float64 {
	return acc.Balance
}

// Let Account implement the Stringer interface to provide a custom string representation
// of the account details.
func (acc *Account) String() string {
	return fmt.Sprintf("Account Number: %s, Owner: %s, Balance: $%.2f", acc.AccountNumber, acc.OwnerName, acc.Balance)
}

type SavingsAccount struct {
	Account      // Embedding the Account struct to achieve composition
	InterestRate float64
}

func (sa *SavingsAccount) ApplyInterest() {
	interest := sa.Balance * sa.InterestRate / 100 // Direct access to the Balance field of the embedded Account struct
	fmt.Printf("Applied %.2f%% interest to account %s. Interest amount: $%.2f, New balance: $%.2f\n",
		sa.InterestRate, sa.AccountNumber, interest, sa.Balance)

	// Using the promoted Deposit method of the embedded Account struct to add the interest to the balance.
	err := sa.Deposit(interest) // Since Deposit returns an error, we should handle it here to ensure that the interest is applied correctly.
	if err != nil {
		fmt.Println("ApplyInterest(): Error applying interest:", err)
	}
}

type OverdraftAccount struct {
	Account        // Embedding the Account struct to achieve composition
	OverdraftLimit float64
}

// Overriding the Withdraw method of the embedded Account struct to allow overdraft up
// to the specified limit.
func (oa *OverdraftAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("Withdrawal amount must be positive. Account: %s", oa.AccountNumber)
	}
	if amount > (oa.Balance + oa.OverdraftLimit) {
		return fmt.Errorf("Withdrawal of $%.2f exceeds overdraft limit of $%.2f. Account: %s", amount, oa.OverdraftLimit, oa.AccountNumber)
	}
	oa.Balance -= amount
	fmt.Printf("Withdrew $%.2f from account %s. New balance: $%.2f\n", amount, oa.AccountNumber, oa.Balance)
	return nil
}

func main() {
	fmt.Printf("--- Welcome to the Bank Account Management System! ---\n")

	// Create a SavingsAccount instance
	savingsAcc := SavingsAccount{
		Account: Account{
			AccountNumber: "SA123456",
			OwnerName:     "Alice Smith",
			Balance:       1000.00,
		},
		InterestRate: 2.5,
	}

	// When we print the savingsAcc, it will use the String() method of the embedded Account struct
	// to display the account details in a formatted manner.
	fmt.Println(savingsAcc) // Using the String() method of the embedded Account struct

	err := savingsAcc.Deposit(-500.00) // Using the promoted Deposit method of the embedded Account struct to add funds to the savings account
	if err != nil {
		fmt.Println(err)
	}

	savingsAcc.ApplyInterest()        // Applying interest to the savings account
	err = savingsAcc.Withdraw(200.00) // Using the promoted Withdraw method of the embedded Account struct to withdraw funds from the savings account
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("--- Savings Account Details After Transactions ---")
	fmt.Println(savingsAcc)

	overdraftacc := OverdraftAccount{
		Account: Account{
			AccountNumber: "OA654321",
			OwnerName:     "Bob Johnson",
			Balance:       500.00,
		},
		OverdraftLimit: 200.00,
	}

	fmt.Println(overdraftacc) // Using the String() method of the embedded Account struct

	// Using the promoted Deposit method of the embedded Account struct to add funds to the overdraft account
	err = overdraftacc.Deposit(200.00) // Using the promoted Deposit method of the embedded Account struct to add funds to the overdraft account
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Overdraft Account Balance, Account %s: $%.2f\n", overdraftacc.AccountNumber, overdraftacc.GetBalance()) // Using the promoted GetBalance method of the embedded Account struct to check the balance of the overdraft account

	err = overdraftacc.Withdraw(600.00) // Using the overridden Withdraw method of the OverdraftAccount struct to withdraw funds, allowing overdraft up to the specified limit
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("--- Overdraft Account Details After Transactions ---")
	fmt.Println(overdraftacc) // Calls the String() method of the embedded Account struct to display the account details in a formatted manner

	err = overdraftacc.Withdraw(1000.00) // Using the overridden Withdraw method of the OverdraftAccount struct to withdraw funds, allowing overdraft up to the specified limit
	if err != nil {
		fmt.Println(err)
	}
}
