package member

import (
	"fmt"
	"strings"
)

type Member struct {
	number   MemberNumber
	name     string
	category MemberCategory
}

func (m *Member) Category() MemberCategory {
	return m.category
}

func (m *Member) Number() string {
	return m.number.String()
}

func (m *Member) String() string {
	var b strings.Builder
	fmt.Fprintf(&b, "Member {")
	fmt.Fprintf(&b, "memberNumber=%s", m.number)
	fmt.Fprintf(&b, ", name=%s", m.name)
	fmt.Fprintf(&b, ", category=%s", m.category)
	fmt.Fprintf(&b, "}")

	return b.String()
}
