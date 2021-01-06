package library

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type jwtHeader struct {
	Type string `json:"type"`
	Algo string `json:"algo"`
}

/*GenerateBase64 ...
@desc Convert payload of type interface to Base64 String format
*/
func GenerateBase64(payload interface{}) string {
	jsonEncode, _ := json.Marshal(payload)

	stringEncoded := base64.StdEncoding.EncodeToString(jsonEncode)

	sp := stringReplace(string(stringEncoded))

	return sp
}

/*Sign ...
@param model interface{}, struct of model
@param modelId primitive.ObjectID, model Id
@desc sign payload to jwt token
*/
func Sign(model interface{}, modelID primitive.ObjectID) string {
	jwtBody := GenerateBase64(model)
	header := jwtHeader{"jwt", "HSH256"}
	jwtHeader := GenerateBase64(header)
	payload := jwtHeader + "." + jwtBody
	signature := jwtSignature(modelID.String(), payload)

	return signature
}

/*
	JWTSignature
	@uid string, user ID
	@p string, JWT Payload
*/
func jwtSignature(uid string, p string) string {
	h := hmac.New(sha256.New, []byte(uid))

	h.Write([]byte(p))

	sha := hex.EncodeToString(h.Sum(nil))

	sign := GenerateBase64(sha)

	return sign
}

/*
	StringReplace replace special char
*/
func stringReplace(s string) string {
	replace := strings.NewReplacer("=", "", "+", "-", "/", "_")

	return replace.Replace(s)
}
