package loanability

import (
	"github.com/pocket7878/loan-rule-go/domain/model/collection"
	"github.com/pocket7878/loan-rule-go/domain/model/loan"
)

type Loanability struct {
	loanContext loan.LoanContext
	entry       collection.Entry
}

func NewLoanabilityOf(loanContext loan.LoanContext, entry collection.Entry) *Loanability {
	return &Loanability{loanContext: loanContext, entry: entry}
}

func (l *Loanability) Judge() LoanabilityType {
	if l.isExceededLoanLimit() {
		return Unloanable
	}
	if l.isLoanSameBook() {
		return Unloanable
	}
	if l.isExceededLoanPeriodExists() {
		return Unloanable
	}

	return Loanable
}

func (l *Loanability) isLoanSameBook() bool {
	return false
}

func (l *Loanability) isExceededLoanLimit() bool {
	currentLoans := l.loanContext.TotalLoanCount()
	maxLoans := l.loanContext.MaxLoanableCount()

	return currentLoans >= maxLoans
}

func (l *Loanability) isExceededLoanPeriodExists() bool {
	// 貸出期限は2週間
	return false
}
