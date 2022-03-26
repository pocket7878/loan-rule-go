package loan

import (
	"github.com/pocket7878/loan-rule-go/domain/model/loan"
	"github.com/pocket7878/loan-rule-go/domain/model/member"
)

type LoanRepository interface {
	History(memberNumber member.MemberNumber) (*loan.LoanHistory, error)
}
