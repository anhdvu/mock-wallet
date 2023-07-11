package main

type message struct {
	Challenge                 string             `json:"challenge,omitempty"`
	CustomerReference         string             `json:"customerReference"`
	CurrencyCode              string             `json:"currencyCode,omitempty"`
	DigitizationPath          string             `json:"digitizationPath,omitempty"`
	DigitizedDeviceIdentifier string             `json:"digitizedDeviceIdentifier,omitempty"`
	DigitizedDeviceType       string             `json:"digitizedDeviceType,omitempty"`
	DigitizedFpanMasked       string             `json:"digitizedFpanMasked,omitempty"`
	DigitizedPan              string             `json:"digitizedPan,omitempty"`
	DigitizedPanExpiry        string             `json:"digitizedPanExpiry,omitempty"`
	DigitizedTokenReference   string             `json:"digitizedTokenReference,omitempty"`
	EventType                 string             `json:"eventType,omitempty"`
	MerchantDescription       string             `json:"merchantDescription,omitempty"`
	MessageType               string             `json:"messageType"`
	ResultCode                string             `json:"resultCode,omitempty"`
	TokenRequestorID          string             `json:"tokenRequestorId,omitempty"`
	TrackingNumber            string             `json:"trackingNumber"`
	TransactionAmount         string             `json:"transactionAmount,omitempty"`
	WalletIdentifier          string             `json:"walletIdentifier,omitempty"`
	WalletRecommendation      string             `json:"walletRecommendation,omitempty"`
	ActivationMethods         []activationMethod `json:"activationMethods,omitempty"`
}

type activationMethod struct {
	Value string `json:"value"`
	Type  int    `json:"type"`
}
