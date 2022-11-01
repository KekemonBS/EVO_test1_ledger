package models

import (
	"time"

	"github.com/bojanz/currency"
)

type Transaction struct {
	TransactionId      int64           `db:"transactionid"`
	RequestId          int64           `db:"requestid"`
	TerminalId         int64           `db:"terminalid"`
	PartnerObjectId    int64           `db:"partnerobjectid"`
	AmountTotal        currency.Amount `db:"amounttotal"`
	AmountOriginal     currency.Amount `db:"amountoriginal"`
	CommissionPS       currency.Amount `db:"commissionps"`
	CommissionClient   currency.Amount `db:"commissionclient"`
	CommissionProvider currency.Amount `db:"commissionprovider"`
	DateInput          time.Time       `db:"dateinput"`
	DatePost           time.Time       `db:"datepost"`
	Status             string          `db:"status"`
	PaymentType        string          `db:"paymenttype"`
	PaymentNumber      string          `db:"paymentnumber"`
	ServiceId          int64           `db:"serviceid"`
	Service            string          `db:"service"`
	PayeeId            int64           `db:"payeeid"`
	PayeeName          string          `db:"payeename"`
	PayeeBankMfo       int64           `db:"payeebankmfo"`
	PayeeBankAccount   string          `db:"payeebankaccount"`
	PaymentNarrative   string          `db:"paymentnarrative"`
}
