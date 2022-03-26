package loan

import "fmt"

type LoanHistory struct {
	history []Loan
}

func NewLoanHistoryOf(history []Loan) *LoanHistory {
	return &LoanHistory{history: history}
}

func (l *LoanHistory) TotalLoandCount() int {
	return len(l.history)
}

func (l *LoanHistory) String() string {
	return fmt.Sprintf("LoanHistory{history=%v}", l.history)
}
