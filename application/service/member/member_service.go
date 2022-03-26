package member

import (
	"github.com/pocket7878/loan-rule-go/domain/model/member"
)

type MemberService struct {
	memberRepository MemberRepository
}

func NewMemberService(memberRepository MemberRepository) *MemberService {
	return &MemberService{memberRepository: memberRepository}
}

func (s *MemberService) Member(memberNumber member.MemberNumber) (*member.Member, error) {
	member, err := s.memberRepository.Get(memberNumber)
	if err != nil {
		return nil, err
	}

	return member, nil
}
