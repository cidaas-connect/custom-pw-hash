package api

type PasswordVerification struct {
	PasswordVerification PasswordVerificationClass `json:"passwordVerification"`
	Password             string                    `json:"password"`
	Algorithm            string                    `json:"algorithm"`
}

type PasswordVerificationClass struct {
	PasswordHash string `json:"passwordHash"`
	HData        HData  `json:"hData"`
}

type HData struct {
	Salt string `json:"salt"`
}

type PasswordVerifiedData struct {
	Verified bool `json:"verified"`
}
