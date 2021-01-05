package library

import (
	"crypto/hmac"
	"crypto/sha256"
	"strings"
	"encoding/json"
	"encoding/base64"
	"encoding/hex"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JWTHeader struct {
	Type string `json:"type"`
	Algo string `json:"algo"`
}

/*
	GenerateBase64
*/
func GenerateBase64(p interface{}) string {
	je, _ := json.Marshal(p)

	stringEncoded := base64.StdEncoding.EncodeToString(je)

	sp := StringReplace(string(stringEncoded))

	return sp
}

/*
	@model interface{}, struct of model
	@modelId primitive.ObjectID, model Id
*/
func Sign(model interface{}, modelId primitive.ObjectID) string {
	jwtBody := GenerateBase64(model)
	header := JWTHeader{"jwt", "HSH256"}
	jwtHeader := GenerateBase64(header)
	payload := jwtHeader + "." + jwtBody
	signature := JWTSignature(modelId.String(), payload)

	return signature
}

/*
	JWTSignature
	@uid string, user ID
	@p string, JWT Payload
*/
func JWTSignature(uid string, p string) string {
	h := hmac.New(sha256.New, []byte(uid))

	h.Write([]byte(p))

	sha := hex.EncodeToString(h.Sum(nil))

	sign := GenerateBase64(sha)

	return sign
}

/*
	StringReplace replace special char
*/
func StringReplace(s string) string {
	replace := strings.NewReplacer("=", "", "+", "-", "/", "_")

	return replace.Replace(s)
}