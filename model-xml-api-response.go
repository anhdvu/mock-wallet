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
		"AdministrativeMessage": &response{
			resultCode: 1,
			validCodes: []int{1},
			boiler:     DefaultXMLResponseBoiler,
		},
		"Balance": &response{
			resultCode: 1,
			validCodes: []int{1},
			boiler:     `<methodResponse><params><param><value><struct><member><name>resultCode</name><value><int>%d</int></value></member><member><name>balanceAmount</name><value><int>626900</int></value></member></struct></value></param></params></methodResponse>`,
		},
		//"BrazilianInstallmentSettled": &response{},
		"Deduct": &response{
			resultCode: 1,
			validCodes: []int{1, 2, -4, -7, -8, -9, -16, -17, -18, -19, -24, -25, -26, -27, -28, -29, -36, -37, -38, -39},
			boiler:     DefaultXMLResponseBoiler,
		},
		"DeductAdjustment": &response{
			resultCode: 1,
			validCodes: []int{1, -7, -8, -9},
			boiler:     DefaultXMLResponseBoiler,
		},
		"DeductReversal": &response{
			resultCode: 1,
			validCodes: []int{1, -7, -8, -9},
			boiler:     DefaultXMLResponseBoiler,
		},
		"LoadAdjustment": &response{
			resultCode: 1,
			validCodes: []int{1, -7, -8, -9},
			boiler:     DefaultXMLResponseBoiler,
		},
		"LoadAuth": &response{
			resultCode: 1,
			validCodes: []int{1, -7, -8, -9},
			boiler:     DefaultXMLResponseBoiler,
		},
		"LoadAuthReversal": &response{
			resultCode: 1,
			validCodes: []int{1, -7, -8, -9},
			boiler:     DefaultXMLResponseBoiler,
		},
		"LoadReversal": &response{
			resultCode: 1,
			validCodes: []int{1, -7, -8, -9},
			boiler:     DefaultXMLResponseBoiler,
		},
		//"MexicanInstallmentSettled":   &response{},
		"Stop": &response{
			resultCode: 1,
			validCodes: []int{1},
			boiler:     DefaultXMLResponseBoiler,
		},
		//"ValidatePIN":                 &response{},
	}

	return &companionResponses{
		d: d,
	}
}

func (cr *companionResponses) allResponses() map[string]int {
	o := make(map[string]int)
	for k, v := range cr.d {
		o[k] = v.resultCode
	}

	return o
}

func (cr *companionResponses) makeResponse(m string) ([]byte, error) {
	cr.mu.Lock()
	defer cr.mu.Unlock()
	if ok := cr.vefiryMethod(m); ok {
		r := fmt.Sprintf(cr.d[m].boiler, cr.d[m].resultCode)
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
