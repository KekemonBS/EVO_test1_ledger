# Show Ledger info

Get Ledger rows depending on search terms.

**URL** : `/q?transactionid=&terminalid=&status=&paymenttype=&datepost=&paymentnarrative=&quantity=&pagenum=`

**Method** : `GET`

**Auth required** : NO
**Permissions required** : None

## Success Response

**Code** : `200 OK`

**Content examples**

For a TransactionId 11 on the local database where ledger records are saved

```json
{
	"TransactionId":11,
	"RequestId":20120,
	"TerminalId":3516,
	"PartnerObjectId":1111,
	"AmountTotal":{"number":"119.00",
	"currency":"UAH"},
	"AmountOriginal":{"number":"119.00","currency":"UAH"},
	"CommissionPS":{"number":"0.08","currency":"UAH"},
	"CommissionClient":{"number":"0.00","currency":"UAH"},
	"CommissionProvider":{"number":"-0.24","currency":"UAH"},
	"DateInput":"2022-08-23T00:00:00Z",
	"DatePost":"2022-08-23T00:00:00Z",
	"Status":"accepted",
	"PaymentType":"cash",
	"PaymentNumber":"PS16698305",
	"ServiceId":14080,
	"Service":"Поповнення карток",
	"PayeeId":15233155,
	"PayeeName":"privat",
	"PayeeBankMfo":264761,
	"PayeeBankAccount":"UA713550970919423",
	"PaymentNarrative":"Перерахування коштів згідно договору про надання послуг А11/27122 від 19.11.2020 р.",
}
```
