package entity

type Person struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName:"`
	Age       int8   `json:"age" binding:"gte=1,lte=130"`
	Email     string `json:"email" validate:"required,email"`
}

type Vote struct {
	Title       string `json:"Title" binding:"required"`
	Description string `json:"Description"`
	VoteScript  string `json:"VoteScript" validate:"is-cool"`
	Author      Person `json:"author" binding:"required"`
}

type VoteBuild struct {
	BashCommand string `json:"bash_command" binding:"required"`
	Vote        string `json:"vote" binding:"required"`
	TX_ID       string `json:"tx_id" binding:"required"`
	Index       string `json:"index" binding:"required"`
	DRepKeyFile string `json:"drep_key_file" binding:"required"`
	OutFile     string `json:"outfile" binding:"required"`
}

type TransactionSign struct {
	TxBodyFile  string `json:"tx_body_file" binding:"required"`
	PaymentSkey string `json:"payment_skey" binding:"required"`
	DRepSkey    string `json:"drep_skey" binding:"required"`
	OutFile     string `json:"out_file" binding:"required"`
}
