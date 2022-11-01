package postgresql

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/KekemonBS/ledgerTest/models"
	"github.com/jackc/pgx/v4"
)

var (
	ErrStore  = errors.New("store error :")
	ErrDelete = errors.New("delete error :")
	ErrCreate = errors.New("create error :")
	ErrRead   = errors.New("read error :")
	ErrQuery  = errors.New("query error :")
)

type DbImpl struct {
	storage *pgx.Conn
}

func New(db *pgx.Conn) *DbImpl {
	return &DbImpl{
		storage: db,
	}
}

func (db DbImpl) Create(ctx context.Context, tr models.Transaction) error {
	query := `
	INSERT INTO ledger 
	(
	TransactionId      ,
	RequestId          ,
	TerminalId         ,
	PartnerObjectId    ,
	AmountTotal        ,
	AmountOriginal     ,
	CommissionPS       ,
	CommissionClient   ,
	CommissionProvider ,
	DateInput          ,
	DatePost           ,
	Status             ,
	PaymentType        ,
	PaymentNumber      ,
	ServiceId          ,
	Service            ,
	PayeeId            ,
	PayeeName          ,
	PayeeBankMfo       ,
	PayeeBankAccount   ,
	PaymentNarrative   
	) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21)
	ON CONFLICT (TransactionId)
	DO UPDATE SET
		RequestId          = excluded.RequestId         ,
		TerminalId         = excluded.TerminalId        ,
		PartnerObjectId    = excluded.PartnerObjectId   ,
		AmountTotal        = excluded.AmountTotal       ,
		AmountOriginal     = excluded.AmountOriginal    ,
		CommissionPS       = excluded.CommissionPS      ,
		CommissionClient   = excluded.CommissionClient  ,
		CommissionProvider = excluded.CommissionProvider,
		DateInput          = excluded.DateInput         ,
		DatePost           = excluded.DatePost          ,
		Status             = excluded.Status            ,
		PaymentType        = excluded.PaymentType       ,
		PaymentNumber      = excluded.PaymentNumber     ,
		ServiceId          = excluded.ServiceId         ,
		Service            = excluded.Service           ,
		PayeeId            = excluded.PayeeId           ,
		PayeeName          = excluded.PayeeName         ,
		PayeeBankMfo       = excluded.PayeeBankMfo      ,
		PayeeBankAccount   = excluded.PayeeBankAccount  ,
		PaymentNarrative   = excluded.PaymentNarrative;
		`
	res, err := db.storage.Exec(ctx, query,
		tr.TransactionId,
		tr.RequestId,
		tr.TerminalId,
		tr.PartnerObjectId,
		tr.AmountTotal,
		tr.AmountOriginal,
		tr.CommissionPS,
		tr.CommissionClient,
		tr.CommissionProvider,
		tr.DateInput,
		tr.DatePost,
		tr.Status,
		tr.PaymentType,
		tr.PaymentNumber,
		tr.ServiceId,
		tr.Service,
		tr.PayeeId,
		tr.PayeeName,
		tr.PayeeBankMfo,
		tr.PayeeBankAccount,
		tr.PaymentNarrative,
	)

	if err != nil {
		return ErrStore
	}

	q := res.RowsAffected()
	if q == 0 {
		return ErrCreate
	}
	return nil
}
func (db DbImpl) Read(ctx context.Context, id int64) (models.Transaction, error) {
	query := `
	SELECT * FROM ledger 
	WHERE TransactionId = $1;
	`
	res, err := db.storage.Query(ctx, query,
		id,
	)
	defer res.Close()
	resTr := models.Transaction{}
	ok := res.Next()
	if !ok {
		return models.Transaction{}, ErrQuery
	}
	err = res.Scan(
		&resTr.TransactionId,
		&resTr.RequestId,
		&resTr.TerminalId,
		&resTr.PartnerObjectId,
		&resTr.AmountTotal,
		&resTr.AmountOriginal,
		&resTr.CommissionPS,
		&resTr.CommissionClient,
		&resTr.CommissionProvider,
		&resTr.DateInput,
		&resTr.DatePost,
		&resTr.Status,
		&resTr.PaymentType,
		&resTr.PaymentNumber,
		&resTr.ServiceId,
		&resTr.Service,
		&resTr.PayeeId,
		&resTr.PayeeName,
		&resTr.PayeeBankMfo,
		&resTr.PayeeBankAccount,
		&resTr.PaymentNarrative,
	)
	if err != nil {
		return models.Transaction{}, ErrRead
	}
	return resTr, nil
}

func (db DbImpl) Delete(ctx context.Context, id int64) error {
	query := `
	DELETE FROM ledger 
	WHERE TransactionId = $1;
	`
	res, err := db.storage.Exec(ctx, query,
		id,
	)
	if err != nil {
		return ErrQuery
	}

	q := res.RowsAffected()
	if q == 0 {
		return ErrCreate
	}
	return nil
}

func (db DbImpl) Search(ctx context.Context, v url.Values) ([]models.Transaction, error) {
	query, args, err := BuildQuery(v)
	if err != nil {
		return nil, ErrQuery
	}
	fmt.Printf("Built query : \n%s\n", query)
	fmt.Printf("Args slice : \n%s\n", args)
	res, err := db.storage.Query(ctx, query, args...)
	defer res.Close()
	if err != nil {
		return nil, ErrQuery
	}

	var resTrs []models.Transaction
	for res.Next() {
		var resTr models.Transaction
		err = res.Scan(
			&resTr.TransactionId,
			&resTr.RequestId,
			&resTr.TerminalId,
			&resTr.PartnerObjectId,
			&resTr.AmountTotal,
			&resTr.AmountOriginal,
			&resTr.CommissionPS,
			&resTr.CommissionClient,
			&resTr.CommissionProvider,
			&resTr.DateInput,
			&resTr.DatePost,
			&resTr.Status,
			&resTr.PaymentType,
			&resTr.PaymentNumber,
			&resTr.ServiceId,
			&resTr.Service,
			&resTr.PayeeId,
			&resTr.PayeeName,
			&resTr.PayeeBankMfo,
			&resTr.PayeeBankAccount,
			&resTr.PaymentNarrative,
		)
		if err != nil {
			return nil, ErrRead
		}
		resTrs = append(resTrs, resTr)
		//fmt.Println("--------------RES------------")
		//fmt.Println(resTr)
	}
	return resTrs, nil
}

func BuildQuery(vm url.Values) (string, []any, error) {
	sizet := vm["quantity"][0]
	nt := vm["pagenum"][0]

	size, err := strconv.Atoi(sizet)
	if err != nil {
		return "", nil, err
	}
	n, err := strconv.Atoi(nt)
	if err != nil {
		return "", nil, err
	}

	offset := size * (n - 1)
	query := `SELECT * FROM ledger WHERE `
	fieldNames := []string{
		"transactionid",    //any id
		"terminalid",       //id1, id2, ...
		"status",           //accepted/denied
		"paymenttype",      //cash/card
		"datepost",         //xxxx-xx-xx/xxxx-xx-xx (<from>/<to>, ISO 8601)
		"paymentnarrative", //partial or full narrative
	}

	var vmf [][2]string
	i := 0
	for k, v := range vm {
		if v[0] != "" && k != "quantity" && k != "pagenum" {
			vmf = append(vmf, [2]string{k, v[0]})
		}
		i++
	}
	argCounter := 0
	var args []any
	for i, val := range vmf {
		k := val[0]
		v := val[1]
		if !inSlice(k, fieldNames) {
			return "", nil, ErrQuery
		}
		switch k {
		case "terminalid": //IN ( $1, $2, ... )
			q := strings.Split(v, ",")
			query += k + "::TEXT IN ( "
			for i := 0; i < len(q)-1; i++ {
				query += "$" + fmt.Sprint(argCounter+i+1) + ", "
			}
			query += "$" + fmt.Sprint(argCounter+len(q)) + " ) "

			argCounter += len(q)
			for _, e := range q {
				args = append(args, e)
			}
			break
		case "datepost": // BETWEEN ? AND ?;
			dates := strings.Split(v, "/")
			query += k + "::TEXT "
			query += "BETWEEN $" + fmt.Sprint(argCounter+1) + " AND $" + fmt.Sprint(argCounter+2) + " "

			argCounter += 2
			args = append(args, dates[0], dates[1])
			break
		default:
			query += k + "::TEXT ~ $" + fmt.Sprint(argCounter+1) + " "

			argCounter += 1
			args = append(args, v)
		}

		if i == len(vmf)-1 {
			continue
		}
		query += " AND "
	}
	query += "LIMIT $" + fmt.Sprint(argCounter+1) + " OFFSET $" + fmt.Sprint(argCounter+2) + ";"
	args = append(args, fmt.Sprint(size), fmt.Sprint(offset))
	return query, args, nil
}

func inSlice[T comparable](e T, s []T) bool {
	for _, i := range s {
		if i == e {
			return true
		}
	}
	return false
}
