package member

import (
	"fmt"
	"strconv"
)

// 会員番号
type MemberNumber struct {
	value int
}

func NewMemberNumber(value string) (*MemberNumber, error) {
	int_value, err := strconv.Atoi(value)
	if err != nil {
		return nil, err
	}

	return &MemberNumber{value: int_value}, nil
}

func (m *MemberNumber) String() string {
	return fmt.Sprintf("%d", m.value)
}
