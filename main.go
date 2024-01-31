package main

import (
	"fmt"
)

type User interface {
	Deposit(amount float64)
	Withdraw(amount float64)
}

type Account struct {
	Name    string
	Balance float64
}

type LoanerAccount struct {
	*Account
	CreditLimit float64
}

type PersonalAccount struct {
	*Account
	InvestmentValue float64
	Investments     []Investment
}

type BusinessAccount struct {
	*Account
}

type Investment struct {
	Amount float64
	Name   string
}

type PF struct {
	Name     string
	Accounts []User
}

type PJ struct {
	Name     string
	Accounts []BusinessAccount
}

type ErrCreateBusinessAccount bool

func (e ErrCreateBusinessAccount) Error() string {
	return "	     * The PJ can only have Business Accounts * \n"
}

func main() {

	// Fixed Data
	namesUser := []string{"User A", "User B", "User C", "User D", "User E"}
	balancesUser := []float64{1200.0, 5000.0, 500.0, 1000.0, 750.0}

	// Create a bank slice to add users infos
	var bankAccounts []User

	// Create a bank account to each user
	for i := 0; i < len(namesUser); i++ {

		var account User

		switch namesUser[i] {
		case "User C":
			account = &LoanerAccount{Account: &Account{Name: namesUser[i], Balance: balancesUser[i]}, CreditLimit: 500.0}
		case "User B":
			account = &PersonalAccount{Account: &Account{Name: namesUser[i], Balance: balancesUser[i]}, InvestmentValue: 0.0}
		case "User E":
			account = &BusinessAccount{Account: &Account{Name: namesUser[i], Balance: balancesUser[i]}}
		default:
			account = &Account{Name: namesUser[i], Balance: balancesUser[i]}
		}

		bankAccounts = append(bankAccounts, account)
	}

	// Create PF and Add accounts
	pf := PF{
		Name: "Natural Person",
		Accounts: []User{
			&Account{Name: "User F", Balance: 1000.0},
			&PersonalAccount{Account: &Account{Name: "User G", Balance: 5000.0}, InvestmentValue: 0.0},
			&LoanerAccount{Account: &Account{Name: "User H", Balance: 0.0}, CreditLimit: 1000.0},
		},
	}

	// Create PJ and add accounts
	pj := PJ{
		Name:     "Legal Person",
		Accounts: []BusinessAccount{},
	}

	// Try add a PersoalAccount to PJ
	fmt.Println("\n ---------- Trying add a PersonalAccount to Legal person ----------")
	pj.AddAccount(&PersonalAccount{Account: &Account{Name: "User I", Balance: 5000.0}, InvestmentValue: 0.0})

	// Add a BusinessAccount to PJ
	pj.AddAccount(&BusinessAccount{Account: &Account{Name: "User J", Balance: 10000.0}})

	//Create a userSlice with []businessAccount values
	var userSlice []User

	for _, bAccount := range pj.Accounts {
		userSlice = append(userSlice, &bAccount)
	}

	fmt.Println("Natural Person (PF):")
	ShowAccounts(pf.Accounts)

	fmt.Println("\nLegal Person (PJ):")
	ShowAccounts(userSlice)

	fmt.Println("\nCreated Accounts:")
	ShowAccounts(bankAccounts)

	//Operation DEPOSIT
	fmt.Println("\n**** Operation: User A deposited $250 ****")

	//Defining user
	//Deposit value
	bankAccounts[0].Deposit(250.0)

	//Operation WITHDRAW
	fmt.Println("\n**** Operation: User C withdrew $1.000 ****")

	//Defining user
	//Withdraw value
	bankAccounts[2].Withdraw(1000.0)

	//Operation INVEST
	fmt.Println("\n**** Operation: User B invested $1.000 ****")

	// Difining user
	// Investing value
	investorUserB := bankAccounts[1].(*PersonalAccount)
	investorUserB.Invest(1000.0, "RSI Consultancy")

	//Operation IaNVEST
	fmt.Println("\n**** Operation: User B invested $500 ****")

	// Difining user
	// Investing value
	investorUserB = bankAccounts[1].(*PersonalAccount)
	investorUserB.Invest(500.0, "RSI Technology")

	//Operation REMOVE
	fmt.Println("\n**** Operation: User B removed $500 investment of RSI Technology ****")
	investorUserB.RemoveInvestment(1)

	//Operation INVEST
	fmt.Println("\n**** Operation: User B invested $1.500 ****")

	// Difining user
	// Investing value
	investorUserB = bankAccounts[1].(*PersonalAccount)
	investorUserB.Invest(1500.0, "RSI Technology")

	//Operation Loan
	fmt.Println("\n**** Operation: User E loaned $200 ****")

	// Difining user
	// Investing value
	businessUser := bankAccounts[4].(*BusinessAccount)
	businessUser.Loan(200.0)

	// Final balance, accounts and investments
	fmt.Println("\n Final balance: ")
	ShowAccounts(bankAccounts)
	investorUserB.ViewInvestments()

	//Make a function to show the operations, repeating so much the Print ""operation, and remove fixed informations in texts.
}

// Deposit to common user
func (account *Account) Deposit(amount float64) {

	// Depositing value
	account.Balance += amount
	//Showing transtion
	fmt.Printf("----------------- Deposit of $%.2f successfully made to %s. New balance: $%.2f \n", amount, account.Name, account.Balance)
}

// Withdraw commom user
func (account *Account) Withdraw(amount float64) {

	if amount > account.Balance {
		fmt.Printf("----------------- The balance is not enough for this withdrawal. Balance: %.2f \n", account.Balance)

	} else {
		//Withdrawing value
		account.Balance -= amount
		fmt.Printf("----------------- Withdraw made successfully. Current Balance: %.2f \n", account.Balance)
	}

}

// Withdraw LoanerUser
func (loaner *LoanerAccount) Withdraw(amount float64) {
	if amount > loaner.Balance+loaner.CreditLimit {
		fmt.Printf("----------------- Balance Available is not enough for this withdraw. Balance with credit: %.2f \n", loaner.Balance+loaner.CreditLimit)
	} else {
		loaner.Balance -= amount
		fmt.Printf("----------------- Withdraw made successfully. Current balance: %.2f \n", loaner.Balance)
	}
}

// Show Accounts all Users
func ShowAccounts(users []User) {
	for _, user := range users {
		switch u := user.(type) {
		case *Account:
			fmt.Printf("----------------- %s: $%.2f\n", u.Name, u.Balance)
		case *LoanerAccount:
			fmt.Printf("----------------- %s [Loaner]: $%.2f (Credit Limit: $%.2f) \n", u.Name, u.Balance, u.CreditLimit)
		case *PersonalAccount:
			fmt.Printf("----------------- %s [Investor]: $%.2f \n", u.Name, u.Balance)
		case *BusinessAccount:
			fmt.Printf("----------------- %s [Business]: $%.2f \n", u.Account.Name, u.Balance)
		}
	}
}

// Only PersonalUser
func (i *PersonalAccount) Invest(amount float64, name string) {
	//
	if amount > i.Balance {
		fmt.Printf("----------------- Balance Available is not enough for this withdraw. Current balance: %.2f \n", i.Balance)
	} else {
		investment := Investment{Amount: amount, Name: name}
		// Setting to history of invesments
		i.Investments = append(i.Investments, investment)

		// Removing value of balance and add value in InvestmentValue
		i.Balance -= amount
		i.InvestmentValue += amount
		fmt.Printf("----------------- Investment of $%.2f successfully made to %s. New balance: $%.2f \n", amount, name, i.Balance)
	}

}

// Only Personal account
func (i *PersonalAccount) RemoveInvestment(index int) {

	// Validate if have investments or the index make a sense
	if index < 0 || index >= len(i.Investments) {
		fmt.Printf("The index is invalid for %s \n", i.Name)
	} else {
		//add value of balance and removing value in InvestmentValue
		investment := i.Investments[index]
		i.Balance += investment.Amount
		i.InvestmentValue -= investment.Amount

		// :index Values before 'index' and index+1: to get values after 'index'. The '...' need to concat the 2 splits.
		i.Investments = append(i.Investments[:index], i.Investments[index+1:]...)
		fmt.Printf("-----------------The investment removed successfully for %s. New balance: %.2f \n", i.Name, i.Balance)
	}

}

// Only personal account
func (i *PersonalAccount) ViewInvestments() {

	fmt.Printf("\n \n -------------------- Investments of %s: -----------------\n", i.Name)

	// Validate if have investments
	if len(i.Investments) == 0 {
		fmt.Println("No investments stored.")
	} else {

		//Reset value before loop
		i.InvestmentValue = 0.0

		for index, investment := range i.Investments {
			fmt.Printf("%d. Name: %s, Amount: $%2.f \n", index, investment.Name, investment.Amount)
			i.InvestmentValue += investment.Amount
		}

		fmt.Printf("\nTotal invested value: $%.2f \n", i.InvestmentValue)
	}
}

// Only businnes
func (b *BusinessAccount) Loan(amount float64) {
	b.Balance += amount
	fmt.Printf("----------------- Loan of $%.2f successfully made to %s. New balance: $%.2f \n", amount, b.Name, b.Balance)
}

func (pj *PJ) AddAccount(account User) error {

	// Validate if account is BusinessAccount type
	businessAccount, isBusinessAccount := account.(*BusinessAccount)
	if !isBusinessAccount {
		err := ErrCreateBusinessAccount(isBusinessAccount)
		fmt.Println(err.Error())
		return err
	} else {

		//Add to pj
		pj.Accounts = append(pj.Accounts, *businessAccount)
		return nil
	}

}
