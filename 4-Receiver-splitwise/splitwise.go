package main

import (
	"fmt"
)

// Person struct represents an individual user.
type Person struct {
	First   string
	Last    string
	Balance float64
}

// FullName returns the full name of the person.
func (p Person) FullName() string {
	return p.First + " " + p.Last
}

// AdjustBalance adjusts the balance of the person by adding the amount.
// Using a pointer receiver since we want to modify the original balance.
func (p *Person) AdjustBalance(amount float64) {
	p.Balance += amount
}

// Expense struct represents an expense made by a person.
type Expense struct {
	Payee  *Person  // Who paid
	Amount float64  // How much was paid
	Splits []Person // Who shares the expense
}

// SplitExpense evenly splits the expense between the involved persons.
func (e *Expense) SplitExpense() {
	splitAmount := e.Amount / float64(len(e.Splits)+1) // +1 to include the payee themselves
	for i := range e.Splits {
		e.Splits[i].AdjustBalance(-splitAmount) // Deduct split amount from each user's balance
	}
	e.Payee.AdjustBalance(splitAmount * float64(len(e.Splits))) // Add total split amount to payee
}

// PrintExpense prints details of the expense and balance adjustments.
func (e *Expense) PrintExpense() {
	fmt.Printf("%s paid $%.2f. Split between:\n", e.Payee.FullName(), e.Amount)
	for _, person := range e.Splits {
		fmt.Printf("- %s (Balance: $%.2f)\n", person.FullName(), person.Balance)
	}
	fmt.Printf("Payee %s (Balance: $%.2f)\n", e.Payee.FullName(), e.Payee.Balance)
}

func main() {
	// Create persons (users)
	alice := Person{First: "Alice", Last: "Smith", Balance: 0}
	bob := Person{First: "Bob", Last: "Brown", Balance: 0}
	charlie := Person{First: "Charlie", Last: "Davis", Balance: 0}

	// Create an expense where Alice pays $100, split between Bob and Charlie
	expensebyalice := Expense{
		Payee:  &alice,
		Amount: 100,
		Splits: []Person{bob, charlie},
	}
	expensebybob := Expense{
		Payee:  &bob,
		Amount: 100,
		Splits: []Person{alice, charlie},
	}

	// Split the expense
	expensebyalice.SplitExpense()

	// Print the details of the expense and balances
	expensebyalice.PrintExpense()

	expensebybob.SplitExpense()

	// Print the details of the expense and balances
	expensebybob.PrintExpense()

	// Further, a callback function could process another expense (e.g., tax or tip)
	applyTax(expensebyalice.Amount, func(amount float64) {
		tax := amount * 0.1       // 10% tax
		alice.AdjustBalance(-tax) // Alice pays the tax
		fmt.Printf("Applied 10%% tax: $%.2f (Alice's new balance: $%.2f)\n", tax, alice.Balance)
	})
	applyTax(expensebybob.Amount, func(amount float64) {
		tax := amount * 0.1     // 10% tax
		bob.AdjustBalance(-tax) // Alice pays the tax
		fmt.Printf("Applied 10%% tax: $%.2f (Bob's new balance: $%.2f)\n", tax, bob.Balance)
	})
	fmt.Printf("Final balance: (Alice) $%.2f (Bob)v$%.2f  (Charlie)$%.2f   ", alice.Balance, bob.Balance, charlie.Balance)

}

// Callback function to process extra charges (e.g., tax or tip)
func applyTax(amount float64, callback func(float64)) {
	callback(amount) //anonymous functions
}
