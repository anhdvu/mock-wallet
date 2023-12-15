package main

import (
	"errors"
	"strconv"
)

var (
	ErrKLVKeyNotExists = errors.New("the KLV key does not exist")
	ErrKLVTooShort     = errors.New("the provided KLV string is too short")
	ErrKLVNotCorrect   = errors.New("the provided KLV string is incorrect")
)

func klvLookUp(k string) (string, error) {
	klvDict := map[string]string{
		"002": "tracking number",
		"004": "original transaction amount",
		"010": "conversion rate",
		"026": "merchant category code",
		"032": "acquiring institution code",
		"037": "retrieval reference number",
		"041": "terminal id",
		"042": "merchant identifier",
		"043": "merchant description",
		"044": "merchant name",
		"045": "transaction type identifier",
		"048": "fraud scoring data",
		"049": "original currency code",
		"050": "from account",
		"052": "pin block",
		"061": "pos data",
		"063": "trace id",
		"067": "extended payment code",
		"068": "is recurring",
		"085": "markup amount",
		"108": "recipient name",
		"109": "recipient address",
		"110": "recipient account number",
		"111": "recipient account number type",
		"250": "capture mode",
		"251": "network",
		"252": "fee type",
		"253": "last four digits",
		"254": "digitized pan",
		"255": "digitized wallet id",
		"256": "adjustment reason",
		"257": "original deduct reference id",
		"258": "markup type",
		"259": "acquirer country",
		"260": "mobile number",
		"261": "transaction fee amount",
		"262": "transaction subtype",
		"263": "card issuer data (Colombia)",
		"264": "tax (Colombia)",
		"265": "tax amount base (Colombia)",
		"266": "retailer data (Colombia)",
		"267": "IAC tax amount (Colombia)",
		"268": "number of installments",
		"269": "customer id (Colombia)",
		"270": "security services data",
		"271": "on-behalf services",
		"272": "original merchant description",
		"273": "installments financing type (Brazil)",
		"274": "status",
		"275": "installments grace period (Mexico)",
		"276": "installments type of credit (Mexico)",
		"277": "payment initiator",
		"278": "payment initiator subtype (mastercard only)",
		"300": "additional amount (Colombia)",
		"301": "second additional amount",
		"302": "cashback pos currency code",
		"303": "cashback pos amount",
		"400": "sender name",
		"401": "sender address (visa direct only)",
		"402": "sender city (visa direct only)",
		"403": "sender state (visa direct only)",
		"404": "sender country (visa direct only)",
		"405": "sanction screening score (visa direct only)",
		"406": "business application identifier (visa direct only)",
		"408": "special condition indicator (visa direct only)",
		"409": "business tax id (visa direct only)",
		"410": "individual tax id (visa direct only)",
		"411": "source of funds",
		"412": "sender account number",
		"413": "sender account number type",
		"414": "mvv (visa direct only)",
		"415": "sender reference number (visa direct only)",
		"416": "is afd transaction",
		"417": "acquirer fee amount",
		"418": "card holder address verification result",
		"419": "card holder postal code",
		"420": "card holder street address",
		"421": "sender date of birth",
		"422": "oct activity check result",
		"423": "sender postal code",
		"424": "recipient city",
		"425": "recipient country",
		"900": "3d secure otp",
		"901": "digitization activation",
		"902": "digitization activation method type",
		"903": "digitization activation method value",
		"904": "digitization activation expiry",
		"905": "digitization final tokenization decision",
		"906": "device name",
		"910": "digitized device id",
		"911": "digitized pan expiry",
		"912": "digitized fpan masked",
		"913": "digitized token reference",
		"915": "digitized token requestor id",
		"916": "visa digitized pan",
		"917": "visa token type",
		"920": "pos transaction status",
		"921": "pos transaction security",
		"922": "pos authorization lifecycle",
		"923": "digitization event type",
		"924": "digitization event reason code",
		"925": "supports partial auth",
		"929": "digitization path",
		"930": "wallet recommendation",
		"931": "tokenization pan source",
		"932": "unique transaction reference",
		"933": "transaction purpose",
		"958": "get contact details reason",
		"998": "json transaction data",
		"999": "generic key",
	}

	val, ok := klvDict[k]
	if !ok {
		return "", ErrKLVKeyNotExists
	}

	return val, nil
}

func breakDownKLV(klv string) ([][4]string, error) {
	var result [][4]string

	if len(klv) < 5 {
		return nil, ErrKLVTooShort
	}

	for len(klv) > 4 {
		key := klv[:3]
		description, err := klvLookUp(key)
		if err != nil {
			return nil, err
		}

		length := klv[3:5]
		lengthInt, err := strconv.Atoi(length)
		if err != nil {
			return nil, err
		}

		if len(klv) < 5+lengthInt {
			return nil, ErrKLVNotCorrect
		}

		value := klv[5:(5 + lengthInt)]
		set := [4]string{key, description, length, value}
		result = append(result, set)
		klv = klv[(5 + lengthInt):]
	}

	return result, nil
}
