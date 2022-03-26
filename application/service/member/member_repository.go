package member

import (
	"github.com/pocket7878/loan-rule-go/domain/model/member"
)

type MemberRepository interface {
	Get(memberNumber member.MemberNumber) (*member.Member, error)
}
