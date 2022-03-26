package loan

import (
	"github.com/pocket7878/loan-rule-go/domain/model/loan"
	"github.com/pocket7878/loan-rule-go/domain/model/member"
)

type LoanService struct {
	loanRepository LoanRepository
}

func NewLoanService(loanRepository LoanRepository) *LoanService {
	return &LoanService{loanRepository: loanRepository}
}

func (s *LoanService) LoanHistory(memberNumber member.MemberNumber) (*loan.LoanHistory, error) {
	return s.loanRepository.History(memberNumber)
}
