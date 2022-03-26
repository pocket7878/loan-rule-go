package loanability

type LoanabilityType int

const (
	Loanable LoanabilityType = iota
	Unloanable
)

func (l *LoanabilityType) String() string {
	switch *l {
	case Loanable:
		return "Loanable"
	case Unloanable:
		return "Unloanable"
	default:
		return "Unknown"
	}
}
