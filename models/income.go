package models

import "time"

// Income model for storing income data
type Income struct {
	ID             int       // Unique identifier for each income record
	Num            string    // Number of the income
	Title          string    // Title or description of the income
	Description    string    // Additional information about the income
	Date           time.Time // Date of the income
	AmountReceived float64   // Amount received for this income
	Subtotal       float64   // Subtotal before taxes or deductions
	IVA            float64   // IVA or tax applied
	Retention      float64   // Any retention or deduction
	Category       string    // Category of the income
	PaymentMethod  string    // Payment method (e.g., bank transfer, cash)
	IncomeSource   string    // Source of the income, such as a client or employer
	IsRecurring    bool      // Whether the income is recurring or not
	Client         Client    // Pointer to the client associated with the income
}

func (i Income) GetDateFormatted() string {
	return i.Date.Format("02/01/06")
}

type Incomes []Income

func (incomes Incomes) GetIncomingAmount() float64 {
	var amount float64
	for _, income := range incomes {
		amount += income.AmountReceived
	}
	return amount
}

func (incomes Incomes) GetIncomesByDate(month Month, year Year) Incomes {
	var filteredIncomes Incomes
	for _, income := range incomes {
		if (income.Date.Month() == time.Month(month.Time) || month.Time == 0) && (income.Date.Year() == year.Time.Year() || year.Time == time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)) {
			filteredIncomes = append(filteredIncomes, income)
		}
	}
	return filteredIncomes
}
