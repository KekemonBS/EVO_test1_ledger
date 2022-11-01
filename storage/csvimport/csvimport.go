package csvimport

import (
	"context"
	"encoding/csv"
	"errors"
	"io"
	"strconv"
	"time"

	"github.com/KekemonBS/ledgerTest/models"
	"github.com/bojanz/currency"
)

const (
	layout = "2006-01-02 15:04:05"
)

type DbImpl interface {
	Create(ctx context.Context, tr models.Transaction) error
}

//ImportCSV writes to storage loaded csv file
func ImportCSVFile(ctx context.Context, db DbImpl, code string, f io.Reader) error {
	reader := csv.NewReader(f)
	reader.Read() //Skip 1 line
	reader.Comma = ','
	reader.Comment = '#'
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		//PARSE TO MODEL
		err = ParseRow(ctx, db, code, record)
		if err != nil {
			return err
		}
		//-------------------------------
	}
	return nil
}
func ParseRow(ctx context.Context, db DbImpl, code string, record []string) error {
	var err error
	t := models.Transaction{}
	if t.TransactionId, err = strconv.ParseInt(record[0], 10, 64); err != nil {
		return errors.New("error parsing TransactionId")
	}
	if t.RequestId, err = strconv.ParseInt(record[1], 10, 64); err != nil {
		return errors.New("error parsing RequestId")
	}
	if t.TerminalId, err = strconv.ParseInt(record[2], 10, 64); err != nil {
		return errors.New("error parsing TerminalId")
	}
	if t.PartnerObjectId, err = strconv.ParseInt(record[3], 10, 64); err != nil {
		return errors.New("error parsing PartnerObjectId")
	}
	if t.AmountTotal, err = currency.NewAmount(record[4], code); err != nil {
		return errors.New("error parsing AmountTotal")
	}
	if t.AmountOriginal, err = currency.NewAmount(record[5], code); err != nil {
		return errors.New("error parsing AmountOriginal")
	}
	if t.CommissionPS, err = currency.NewAmount(record[6], code); err != nil {
		return errors.New("error parsing CommissionPS")
	}
	if t.CommissionClient, err = currency.NewAmount(record[7], code); err != nil {
		return errors.New("error parsing CommissionClient")
	}
	if t.CommissionProvider, err = currency.NewAmount(record[8], code); err != nil {
		return errors.New("error parsing CommissionProvider")
	}
	if t.DateInput, err = time.Parse(layout, record[9]); err != nil {
		return errors.New("error parsing TimeInput")
	}
	if t.DatePost, err = time.Parse(layout, record[10]); err != nil {
		return errors.New("error parsing TimePost")
	}
	t.Status = record[11]
	t.PaymentType = record[12]
	t.PaymentNumber = record[13]
	if t.ServiceId, err = strconv.ParseInt(record[14], 10, 64); err != nil {
		return errors.New("error parsing ServiceId")
	}
	t.Service = record[15]
	if t.PayeeId, err = strconv.ParseInt(record[16], 10, 64); err != nil {
		return errors.New("error parsing PayeeId")
	}
	t.PayeeName = record[17]
	if t.PayeeBankMfo, err = strconv.ParseInt(record[18], 10, 64); err != nil {
		return errors.New("error parsing PayeeBankMfo")
	}
	t.PayeeBankAccount = record[19]
	t.PaymentNarrative = record[20]

	//WRITE TO STORAGE
	//fmt.Println(t)
	err = db.Create(ctx, t)
	if err != nil {
		return err
	}
	return nil
}
