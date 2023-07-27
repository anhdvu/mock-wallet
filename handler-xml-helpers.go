package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (app *application) respondXML(w http.ResponseWriter, status int, response []byte) error {
	response = append([]byte(xml.Header), response...)
	w.Header().Add("Content-Type", "application/xml")
	w.WriteHeader(status)
	w.Write(response)

	return nil
}

func (app *application) readXML(r *http.Request, dst any, l *logRecord) error {
	var b bytes.Buffer

	decoder := xml.NewDecoder(io.TeeReader(r.Body, &b))

	err := decoder.Decode(dst)
	if err != nil {
		var syntaxError *xml.SyntaxError
		var tagPathError *xml.TagPathError

		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formatted XML - %s", syntaxError.Error())
		case errors.As(err, &tagPathError):
			return fmt.Errorf("body contains tag path error Field1 - %s, Field2 - %s, Tag1 - %s, Tag2 - %s", tagPathError.Field1, tagPathError.Field2, tagPathError.Tag1, tagPathError.Tag2)
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")
		default:
			return err
		}
	}

	err = decoder.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must contain a single XML value")
	}

	l.RawRequest = b.String()

	return nil
}

func (app *application) processXMLPayload(x *xmlPayload, log *logRecord) error {
	switch x.MethodName {
	case "Deduct", "LoadAuth":
		p := &authorisation{
			MethodName:      x.MethodName,
			Terminal:        x.Params[0].Value.StringParam,
			Reference:       x.Params[1].Value.StringParam,
			Amount:          x.Params[2].Value.IntParam,
			Narrative:       x.Params[3].Value.StringParam,
			TransactionType: x.Params[4].Value.StringParam,
			KLV:             x.Params[5].Value.StringParam,
			TransactionID:   x.Params[6].Value.StringParam,
			TransactionDate: x.Params[7].Value.TimeParam,
			Checksum:        x.Params[8].Value.StringParam,
		}

		log.Checksum = p.Checksum
		b, err := json.MarshalIndent(p, "", "  ")
		if err != nil {
			return err
		}
		log.JSON = string(b)

		klv, err := breakDownKLV(p.KLV)
		if err != nil {
			log.KLVBreakdown = nil
			return err
		}
		log.KLVBreakdown = klv

	case "Balance":
		p := &balance{
			MethodName:      x.MethodName,
			Terminal:        x.Params[0].Value.StringParam,
			Reference:       x.Params[1].Value.StringParam,
			MsgType:         x.Params[2].Value.StringParam,
			KLV:             x.Params[3].Value.StringParam,
			TransactionID:   x.Params[4].Value.StringParam,
			TransactionDate: x.Params[5].Value.TimeParam,
			Checksum:        x.Params[6].Value.StringParam,
		}

		log.Checksum = p.Checksum
		b, err := json.MarshalIndent(p, "", "  ")
		if err != nil {
			return err
		}
		log.JSON = string(b)

		klv, err := breakDownKLV(p.KLV)
		if err != nil {
			log.KLVBreakdown = nil
			return err
		}
		log.KLVBreakdown = klv

	case "Stop":
		p := &stop{
			MethodName:      x.MethodName,
			Terminal:        x.Params[0].Value.StringParam,
			Reference:       x.Params[1].Value.StringParam,
			CardNumber:      x.Params[2].Value.StringParam,
			ReasonCode:      x.Params[3].Value.IntParam,
			KLV:             x.Params[4].Value.StringParam,
			TransactionID:   x.Params[5].Value.StringParam,
			TransactionDate: x.Params[6].Value.TimeParam,
			Checksum:        x.Params[7].Value.StringParam,
		}

		log.Checksum = p.Checksum
		b, err := json.MarshalIndent(p, "", "  ")
		if err != nil {
			return err
		}
		log.JSON = string(b)

		klv, err := breakDownKLV(p.KLV)
		if err != nil {
			log.KLVBreakdown = nil
			return err
		}
		log.KLVBreakdown = klv

	case "AdministrativeMessage":
		p := &administrativeMessage{
			MethodName:      x.MethodName,
			Terminal:        x.Params[0].Value.StringParam,
			Reference:       x.Params[1].Value.StringParam,
			MsgType:         x.Params[2].Value.StringParam,
			KLV:             x.Params[3].Value.StringParam,
			TransactionID:   x.Params[4].Value.StringParam,
			TransactionDate: x.Params[5].Value.TimeParam,
			Checksum:        x.Params[6].Value.StringParam,
		}

		log.Checksum = p.Checksum
		b, err := json.MarshalIndent(p, "", "  ")
		if err != nil {
			return err
		}
		log.JSON = string(b)

		klv, err := breakDownKLV(p.KLV)
		if err != nil {
			log.KLVBreakdown = nil
			return err
		}
		log.KLVBreakdown = klv

	case "Load":
		p := &authorisation{
			MethodName:      x.MethodName,
			Terminal:        x.Params[0].Value.StringParam,
			Reference:       x.Params[1].Value.StringParam,
			Amount:          x.Params[2].Value.IntParam,
			Narrative:       x.Params[3].Value.StringParam,
			TransactionType: x.Params[4].Value.StringParam,
			KLV:             "",
			TransactionID:   x.Params[5].Value.StringParam,
			TransactionDate: x.Params[6].Value.TimeParam,
			Checksum:        x.Params[7].Value.StringParam,
		}

		log.Checksum = p.Checksum
		b, err := json.MarshalIndent(p, "", "  ")
		if err != nil {
			return err
		}
		log.JSON = string(b)

		klv, err := breakDownKLV(p.KLV)
		if err != nil {
			log.KLVBreakdown = nil
			return err
		}
		log.KLVBreakdown = klv

	default:
		p := &settlement{
			MethodName:         x.MethodName,
			Terminal:           x.Params[0].Value.StringParam,
			Reference:          x.Params[1].Value.StringParam,
			Amount:             x.Params[2].Value.IntParam,
			Narrative:          x.Params[3].Value.StringParam,
			KLV:                x.Params[4].Value.StringParam,
			RefTransactionID:   x.Params[5].Value.StringParam,
			RefTransactionDate: x.Params[6].Value.TimeParam,
			TransactionID:      x.Params[7].Value.StringParam,
			TransactionDate:    x.Params[8].Value.TimeParam,
			Checksum:           x.Params[9].Value.StringParam,
		}

		log.Checksum = p.Checksum
		b, err := json.MarshalIndent(p, "", "  ")
		if err != nil {
			return err
		}
		log.JSON = string(b)

		klv, err := breakDownKLV(p.KLV)
		if err != nil {
			log.KLVBreakdown = nil
			return err
		}
		log.KLVBreakdown = klv
	}

	return nil
}
