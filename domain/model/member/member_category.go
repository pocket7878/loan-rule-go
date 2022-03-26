package member

type MemberCategory struct {
	max_loanable_stock_count int
}

func NewNormalMemberCategory() *MemberCategory {
	return &MemberCategory{max_loanable_stock_count: 3}
}

func NewChildMemberCategory() *MemberCategory {
	return &MemberCategory{max_loanable_stock_count: 2}
}

func (c *MemberCategory) MaxLoanableStockCount() int {
	return c.max_loanable_stock_count
}
