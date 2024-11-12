package controller

import (
	"example/hello/entity"
	"example/hello/service"
	validators "example/hello/validtors"
	"fmt"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VoteController interface {
	FindAll(ctx *gin.Context) []entity.Vote
	Save(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
}

type voteController struct {
	service service.VoteService
}

var voteValidator *validator.Validate

func New(service service.VoteService) VoteController {
	voteValidator = validator.New()
	voteValidator.RegisterValidation("is-cool", validators.ValidateVote)
	return &voteController{
		service: service,
	}
}

func (c *voteController) FindAll(ctx *gin.Context) []entity.Vote {
	return c.service.FindAll()
}

func (c *voteController) Save(ctx *gin.Context) error {
	var vote entity.Vote
	err := ctx.ShouldBind(&vote)
	if err != nil {
		return err
	}
	err = voteValidator.Struct(vote)
	c.service.Save(vote)
	return err
}

func (c *voteController) ShowAll(ctx *gin.Context) {
	votes := c.service.FindAll()
	data := gin.H{
		"title:": "Vote Page",
		"votes":  votes,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}

func SignTransaction(c *gin.Context) {
	var signData entity.TransactionSign
	if err := c.ShouldBindJSON(&signData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Here you would use the fields from signData to construct and run the cardano-cli command
	// Example:
	cmd := fmt.Sprintf("cardano-cli conway transaction sign --tx-body-file %s --signing-key-file %s --signing-key-file %s --out-file %s",
		signData.TxBodyFile,
		signData.PaymentSkey,
		signData.DRepSkey,
		signData.OutFile)

	// Execute the command (Note: This is a simplistic example, actual implementation would require error handling, security considerations etc.)
	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sign transaction: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction signed successfully", "output": string(out)})
}
