package soap

func NewSecurityHeader(action *Action, security *Security) *SecurityHeader {
	return &SecurityHeader{
		Wsa:      "http://www.w3.org/2005/08/addressing",
		Action:   action,
		Security: security,
	}
}

// SecurityHeader represents a header with WSSE
type SecurityHeader struct {
	Wsa      string    `xml:"xmlns:wsa,attr,omitempty"`
	Action   *Action   `xml:"wsa:Action,omitempty"`
	Security *Security `xml:"wsse:Security,omitempty"`
}

func NewAction(mustUnderstand, action string) *Action {
	return &Action{
		MustUnderstand: mustUnderstand,
		Value:          action,
	}
}

// Action is a struct for storing SOAP1.2 action header
type Action struct {
	MustUnderstand string `xml:"SOAP-ENV:mustUnderstand,attr,omitempty"`
	Value          string `xml:",chardata"`
}

func NewPasswordSecurity(username, password string) *Security {
	return &Security{
		Wsse: "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd",
		UsernameToken: &UsernameToken{
			Username: username,
			Password: Password{
				Type:  "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordText",
				Value: password,
			},
		},
	}
}

// Security provides authorization data
type Security struct {
	Wsse          string         `xml:"xmlns:wsse,attr,omitempty"`
	UsernameToken *UsernameToken `xml:"wsse:UsernameToken,omitempty"`
}

// UsernameToken stores username and password
type UsernameToken struct {
	Username string   `xml:"wsse:Username"`
	Password Password `xml:"wsse:Password"`
}

// Password represents a password type
type Password struct {
	Type  string `xml:"Type,attr"`
	Value string `xml:",chardata"`
}
