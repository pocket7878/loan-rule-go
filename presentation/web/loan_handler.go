package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pocket7878/loan-rule-go/application/activity"
	"github.com/pocket7878/loan-rule-go/domain/model/collection"
	"github.com/pocket7878/loan-rule-go/domain/model/collection/book"
	"github.com/pocket7878/loan-rule-go/domain/model/member"
)

type LoanController struct {
	loanActivity activity.LoanActivity
}

// GET /:member_number/loans
func (c *LoanController) GetLoans(ctx *gin.Context) {
	memberNumberValue := ctx.Param("member_number")
	memberNumber, err := member.NewMemberNumber(memberNumberValue)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	loanContext, err := c.loanActivity.LoanContext(*memberNumber)
	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}
	entryList, err := c.loanActivity.EntryList()
	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}

	entryViews := make([]*entryView, 0)
	for _, entry := range entryList.List() {
		entryViews = append(entryViews, newEntryView(entry))
	}

	ctx.HTML(http.StatusOK, "loan/index.html.tmpl", gin.H{
		"member":       loanContext.Member(),
		"memberNumber": loanContext.MemberNumber(),
		"loanHistory":  loanContext.LoanHistory(),
		"entryList":    entryViews,
	})
}

// GET /:member_number/loanability/:book_number
func (c *LoanController) GetLoanability(ctx *gin.Context) {
	memberNumberValue := ctx.Param("member_number")
	memberNumber, err := member.NewMemberNumber(memberNumberValue)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	bookNumberValue := ctx.Param("book_number")
	bookNumber := book.NewBookNumber(bookNumberValue)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	loanContext, err := c.loanActivity.LoanContext(*memberNumber)
	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}

	loanabilityType, err := c.loanActivity.CheckLoanability(*loanContext, *bookNumber)
	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}

	ctx.HTML(http.StatusOK, "loanability/show.html.tmpl", gin.H{
		"loanability": loanabilityType.String(),
	})
}

type entryView struct {
	Description string
	BookNumber  string
}

func newEntryView(entry collection.Entry) *entryView {
	return &entryView{
		Description: entry.Description(),
		BookNumber:  entry.Number(),
	}
}
