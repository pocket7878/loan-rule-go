package loan

import (
	"fmt"

	"github.com/pocket7878/loan-rule-go/domain/model/member"
)

type LoanContext struct {
	member      member.Member
	loanHistory LoanHistory
}

func NewLoanContextOf(member member.Member, loanHistory LoanHistory) *LoanContext {
	return &LoanContext{member: member, loanHistory: loanHistory}
}

func (l *LoanContext) MaxLoanableCount() int {
	category := l.member.Category()
	return category.MaxLoanableStockCount()
}

func (l *LoanContext) TotalLoanCount() int {
	return l.loanHistory.TotalLoandCount()
}

func (l *LoanContext) Member() string {
	return l.member.String()
}

func (l *LoanContext) MemberNumber() string {
	return l.member.Number()
}

func (l *LoanContext) LoanHistory() string {
	return l.loanHistory.String()
}

func (l *LoanContext) String() string {
	return fmt.Sprintf("LoanContext{member=%s, loanHistory=%s}", l.member, l.loanHistory)
}
