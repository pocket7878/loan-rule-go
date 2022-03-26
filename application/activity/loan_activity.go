package activity

import (
	serviceCollection "github.com/pocket7878/loan-rule-go/application/service/collection"
	serviceLoan "github.com/pocket7878/loan-rule-go/application/service/loan"
	serviceMember "github.com/pocket7878/loan-rule-go/application/service/member"

	domainCollection "github.com/pocket7878/loan-rule-go/domain/model/collection"
	domainBook "github.com/pocket7878/loan-rule-go/domain/model/collection/book"
	domainLoan "github.com/pocket7878/loan-rule-go/domain/model/loan"
	domainLoanability "github.com/pocket7878/loan-rule-go/domain/model/loanability"
	domainMember "github.com/pocket7878/loan-rule-go/domain/model/member"
)

type LoanActivity struct {
	memberService     serviceMember.MemberService
	collectionService serviceCollection.CollectionService
	loanService       serviceLoan.LoanService
}

func NewLoanActivity(memberService serviceMember.MemberService, collectionService serviceCollection.CollectionService, loanService serviceLoan.LoanService) *LoanActivity {
	return &LoanActivity{memberService: memberService, collectionService: collectionService, loanService: loanService}
}

func (a *LoanActivity) LoanContext(memberNumber domainMember.MemberNumber) (*domainLoan.LoanContext, error) {
	member, err := a.memberService.Member(memberNumber)
	if err != nil {
		return nil, err
	}

	loanHistory, err := a.loanService.LoanHistory(memberNumber)
	if err != nil {
		return nil, err
	}

	return domainLoan.NewLoanContextOf(*member, *loanHistory), nil
}

func (a *LoanActivity) EntryList() (*domainCollection.EntryList, error) {
	return a.collectionService.EntryList()
}

func (a *LoanActivity) CheckLoanability(loanContext domainLoan.LoanContext, bookNumber domainBook.BookNumber) (*domainLoanability.LoanabilityType, error) {
	entry, err := a.collectionService.Entry(bookNumber)
	if err != nil {
		return nil, err
	}
	loanability := domainLoanability.NewLoanabilityOf(loanContext, *entry)
	loanabilityType := loanability.Judge()

	return &loanabilityType, nil
}
