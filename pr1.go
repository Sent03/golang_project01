package main
import "fmt"

type BankAccount struct {
	owner   string
	balance float64
}

func (b *BankAccount) Deposit(amount float64) {
	b.balance += amount
	fmt.Println("Средства успешно зачислены на счет")
}
func (b *BankAccount) Withdraw(amount float64) error {
	if amount < 0 {
		return fmt.Errorf("сумма должна быть положительной")
	}
	if b.balance < amount {
		return fmt.Errorf("недостаточно средств на счете")
	}
	b.balance -= amount
	fmt.Println("Средства успешно сняты со счета")
	return nil
}

func (b *BankAccount) GetBalance() float64 {
	return b.balance
}

func main {
	
}