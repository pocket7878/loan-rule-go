package loan

import (
	"fmt"
	"time"

	"github.com/pocket7878/loan-rule-go/domain/model/collection"
	"github.com/pocket7878/loan-rule-go/domain/model/member"
)

type Loan struct {
	stock        collection.Item
	memberNumber member.MemberNumber
	loanedAt     time.Time
}

func NewLoan(stock collection.Item, memberNumber member.MemberNumber, loanedAt time.Time) *Loan {
	return &Loan{stock: stock, memberNumber: memberNumber, loanedAt: loanedAt}
}

func (l *Loan) String() string {
	return fmt.Sprintf("Loan{stock=%s, memberNumber=%s, loanedAt=%s}", l.stock, l.memberNumber, l.loanedAt)
}
