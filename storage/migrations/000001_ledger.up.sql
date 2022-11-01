CREATE TYPE price AS (
   number NUMERIC,
   currency_code TEXT
);
CREATE TABLE IF NOT EXISTS ledger (
	TransactionId SERIAL PRIMARY KEY,
    RequestId SERIAL,
	TerminalId SERIAL,
	PartnerObjectId SERIAL,
    AmountTotal price,
	AmountOriginal price,
    CommissionPS price,
	CommissionClient price,
	CommissionProvider price,	
    DateInput DATE,
	DatePost DATE,
    Status TEXT,
	PaymentType TEXT,
	PaymentNumber TEXT,
	ServiceId SERIAL,
	Service TEXT,
	PayeeId SERIAL,
	PayeeName TEXT,
	PayeeBankMfo SERIAL,
	PayeeBankAccount TEXT,
	PaymentNarrative TEXT
);
