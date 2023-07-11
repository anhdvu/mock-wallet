package main

import "encoding/xml"

type authorisation struct {
	MethodName      string `json:"method_name,omitempty"`
	Terminal        string `json:"terminal,omitempty"`
	Reference       string `json:"reference,omitempty"`
	Amount          string `json:"amount,omitempty"`
	Narrative       string `json:"narrative,omitempty"`
	TransactionType string `json:"transaction type,omitempty"`
	KLV             string `json:"klv,omitempty"`
	TransactionID   string `json:"transaction_id,omitempty"`
	TransactionDate string `json:"transaction_date,omitempty"`
	Checksum        string `json:"checksum,omitempty"`
}

type settlement struct {
	MethodName         string `json:"method_name"`
	Terminal           string `json:"terminal"`
	Reference          string `json:"reference"`
	Amount             string `json:"amount"`
	Narrative          string `json:"narrative"`
	KLV                string `json:"klv"`
	RefTransactionID   string `json:"reference_transaction_id"`
	RefTransactionDate string `json:"reference_transaction_date"`
	TransactionID      string `json:"transaction_id"`
	TransactionDate    string `json:"transaction_date"`
	Checksum           string `json:"checksum"`
}

type balance struct {
	MethodName      string `json:"method_name"`
	Terminal        string `json:"terminal"`
	Reference       string `json:"reference"`
	MsgType         string `json:"message_type"`
	KLV             string `json:"klv"`
	TransactionID   string `json:"transaction_id"`
	TransactionDate string `json:"transaction_date"`
	Checksum        string `json:"checksum"`
}

type stop struct {
	MethodName      string `json:"method_name"`
	Terminal        string `json:"terminal"`
	Reference       string `json:"reference"`
	CardNumber      string `json:"card_number"`
	ReasonCode      string `json:"reason_code"`
	KLV             string `json:"klv"`
	TransactionID   string `json:"transaction_id"`
	TransactionDate string `json:"transaction_date"`
	Checksum        string `json:"checksum"`
}

type administrativeMessage struct {
	MethodName      string `json:"method_name"`
	Terminal        string `json:"terminal"`
	Reference       string `json:"reference"`
	MsgType         string `json:"message_type"`
	KLVData         string `json:"klv"`
	TransactionID   string `json:"transaction_id"`
	TransactionDate string `json:"transaction_date"`
	Checksum        string `json:"checksum"`
}

type xmlPayload struct {
	XMLName    xml.Name `xml:"methodCall"`
	MethodName string   `xml:"methodName"`
	Params     []struct {
		Value struct {
			StringParam string `xml:"string,omitempty"`
			IntParam    string `xml:"int,omitempty"`
			TimeParam   string `xml:"dateTime.iso8601,omitempty"`
		} `xml:"value"`
	} `xml:"params>param"`
}
