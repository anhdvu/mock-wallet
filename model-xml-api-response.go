package main

import (
	"errors"
	"fmt"
	"sync"
)

const (
	DefaultXMLResponseBoiler = `<methodResponse><params><param><value><struct><member><name>resultCode</name><value><int>%d</int></value></member></struct></value></param></params></methodResponse>`
)

var (
	ErrInvalidResponseCode = errors.New("the given response code was invalid")
	ErrMethodNotExist      = errors.New("the given method does not exist")
)

type companionResponses struct {
	d  map[string]*response
	am map[string]*response
	mu sync.RWMutex
}

type response struct {
	boiler     string
	validCodes []int
	resultCode int
}

func defaultCompanionResponses() *companionResponses {
	d := map[string]*response{
		"AdministrativeMessage": {
			resultCode: 1,
			validCodes: []int{1},
			boiler:     DefaultXMLResponseBoiler,
		},
		"Balance": {
			resultCode: 1,
			validCodes: []int{1},
			boiler:     `<methodResponse><params><param><value><struct><member><name>resultCode</name><value><int>%d</int></value></member><member><name>balanceAmount</name><value><int>626900</int></value></member></struct></value></param></params></methodResponse>`,
		},
		//"BrazilianInstallmentSettled": &response{},
		"Deduct": {
			resultCode: 1,
			validCodes: []int{1, 2, -4, -7, -8, -9, -16, -17, -18, -19, -24, -25, -26, -27, -28, -29, -36, -37, -38, -39},
			boiler:     DefaultXMLResponseBoiler,
		},
		"DeductAdjustment": {
			resultCode: 1,
			validCodes: []int{1, -7, -8, -9},
			boiler:     DefaultXMLResponseBoiler,
		},
		"DeductReversal": {
			resultCode: 1,
			validCodes: []int{1, -7, -8, -9},
			boiler:     DefaultXMLResponseBoiler,
		},
		"LoadAdjustment": {
			resultCode: 1,
			validCodes: []int{1, -7, -8, -9},
			boiler:     DefaultXMLResponseBoiler,
		},
		"LoadAuth": {
			resultCode: 1,
			validCodes: []int{1, -7, -8, -9},
			boiler:     DefaultXMLResponseBoiler,
		},
		"LoadAuthReversal": {
			resultCode: 1,
			validCodes: []int{1, -7, -8, -9},
			boiler:     DefaultXMLResponseBoiler,
		},
		"LoadReversal": {
			resultCode: 1,
			validCodes: []int{1, -7, -8, -9},
			boiler:     DefaultXMLResponseBoiler,
		},
		//"MexicanInstallmentSettled":   &response{},
		"Stop": {
			resultCode: 1,
			validCodes: []int{1},
			boiler:     DefaultXMLResponseBoiler,
		},
		//"ValidatePIN":                 &response{},
	}

	am := map[string]*response{
		"3DSecureOTP": {DefaultXMLResponseBoiler, []int{1}, 1},
		"cardholder.maskedContactDetails": {
			resultCode: 1,
			validCodes: []int{1},
			boiler:     `<methodresponse><params><param><value><struct><member><name>resultCode</name><value><int>%d</int></value></member><member><name>maskedContactDetails</name><value><array><data><value><struct><member><name>type</name><value>phoneNumber</value></member><member><name>value</name><value>(###) ### 4321</value></member></struct></value><value><struct><member><name>type</name><value>emailAddress</value></member><member><name>value</name><value>joh***n@anymail.com</value></member></struct></value></data></array></value></member></struct></value></param></params></methodresponse>`,
		},
		"digitization.activation": {DefaultXMLResponseBoiler, []int{1}, 1},
		"digitization.activationmethods": {
			resultCode: 1,
			validCodes: []int{1},
			boiler:     `<methodResponse><params><param><value><struct><member><name>resultCode</name><value><int>%d</int></value></member><member><name>activationMethods</name><value><array><data><value><struct><member><name>type</name><value>1</value></member><member><name>value</name><value>+1(###) ### 4567</value></member></struct></value><value><struct><member><name>type</name><value>2</value></member><member><name>value</name><value>joh***n@anymail.com</value></member></struct></value></data></array></value></member></struct></value></param></params></methodResponse>`,
		},
		"digitization.complete":                     {DefaultXMLResponseBoiler, []int{1}, 1},
		"digitization.event.Deleted":                {DefaultXMLResponseBoiler, []int{1}, 1},
		"digitization.event.Deleted_from_device":    {DefaultXMLResponseBoiler, []int{1}, 1},
		"digitization.event.Stopped":                {DefaultXMLResponseBoiler, []int{1}, 1},
		"digitization.event.Digitized":              {DefaultXMLResponseBoiler, []int{1}, 1},
		"digitization.event.Digitization_Exception": {DefaultXMLResponseBoiler, []int{1}, 1},
		"digitization.event.Replacement":            {DefaultXMLResponseBoiler, []int{1}, 1},
	}

	return &companionResponses{
		d:  d,
		am: am,
	}
}

func (cr *companionResponses) allResponses() map[string]int {
	o := make(map[string]int)
	for k, v := range cr.d {
		o[k] = v.resultCode
	}

	return o
}

func (cr *companionResponses) makeResponse(method, extraMessage string) ([]byte, error) {
	cr.mu.Lock()
	defer cr.mu.Unlock()
	if ok := cr.vefiryMethod(method); ok {
		if method == "AdministrativeMessage" {
			r := fmt.Sprintf(cr.am[extraMessage].boiler, cr.am[extraMessage].resultCode)
			return []byte(r), nil
		}
		r := fmt.Sprintf(cr.d[method].boiler, cr.d[method].resultCode)
		return []byte(r), nil

	}
	return nil, ErrMethodNotExist
}

func (cr *companionResponses) vefiryMethod(m string) bool {
	if _, ok := cr.d[m]; ok {
		return ok
	}
	return false
}

func (cr *companionResponses) updateResponseCode(m string, code int) error {
	cr.mu.Lock()
	defer cr.mu.Unlock()
	if ok := cr.vefiryMethod(m); ok {
		return cr.d[m].updateCode(code)
	}

	return ErrMethodNotExist
}

func (r *response) validateCode(code int) bool {
	for _, c := range r.validCodes {
		if code == c {
			return true
		}
	}
	return false
}

func (r *response) updateCode(code int) error {
	if ok := r.validateCode(code); ok {
		r.resultCode = code
		return nil
	}
	return ErrInvalidResponseCode
}
