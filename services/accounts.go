package services

import "time"

type AccountService interface {
	CreateAccount(dto AccountCreatedDTO) (*AccountDTO, error)
	Transfer(dto AccountTransferDTO) (TransferedStatus, error)
	StoreValue(dto AccountTransferDTO) (TransferedStatus, error)
	GetEnvelopeAccountByUserId(userId string) *AccountDTO
}

//账户交易的参与者
type TradeParticipator struct {
	AccountNo string
	UserId string
	Username string
}

//账户描述
type AccountTransferDTO struct {
	TradeNo string
	TradeBody TradeParticipator
	TradeTarget TradeParticipator
	AmountStr string
	ChangeType ChangeType
	ChangeFlag ChangeFlag
	Decs string
}

//创建账户
type AccountCreatedDTO struct {
	UserId string
	Username string
	AccountName string
	AccountType ChangeType
	CurrencyCode string
	Amount string
}
//账户信息
type AccountDTO struct {
	AccountCreatedDTO
	AccountNo string
	CreatedAt time.Time
}