package internalModels

import (
	"encoding/base64"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
)

type PasswordObject struct {
	PasswordHash       string            `json:"passwordHash"`
	Method             string            `json:"method"`
	MethodData         map[string]string `json:"methodData"`
	PolicyCreatedUnder string            `json:"policyCreatedUnder"`
}

// GenerateAndSetPasswordHash hash the password using the bcrypt algorithm
func (p *PasswordObject) GenerateAndSetPasswordHash(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	p.PasswordHash = string(bytes)
	p.Method = "bcrypt"
	p.MethodData = map[string]string{
		"cost": "14",
	}
	return err
}

// ComparePasswordHash compare the password hash with the password
func (p *PasswordObject) ComparePasswordHash(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(p.PasswordHash), []byte(password))
}

// EncodeToBase64 convert the password object to a base64 encoded json string
func (p *PasswordObject) EncodeToBase64() (string, error) {
	// Serialize the object
	serialized, err := json.Marshal(p)
	if err != nil {
		return "", err
	}

	// Encode to base64
	encodedText := base64.StdEncoding.EncodeToString([]byte(serialized))
	return encodedText, nil
}

// DecodeFromBase64 convert the base64 encoded json string to a password object
func (p *PasswordObject) DecodeFromBase64(encodedText string) error {
	// Decode from base64
	decoded, err := base64.StdEncoding.DecodeString(encodedText)
	if err != nil {
		return err
	}

	// Deserialize the object
	err = json.Unmarshal(decoded, p)
	if err != nil {
		return err
	}
	return nil
}
