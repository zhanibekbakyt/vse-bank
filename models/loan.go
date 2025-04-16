package models

type Loan struct {
	ID         uint    `json:"id" gorm:"primaryKey"`
	Amount     float64 `json:"amount"`
	UserID     uint    `json:"user_id"`
	BankID     uint    `json:"bank_id"`
	LoanTypeID uint    `json:"loan_type_id"`

	Bank Bank `json:"bank" gorm:"foreignKey:BankID"`
	User User `json:"user" gorm:"foreignKey:UserID"`
}
