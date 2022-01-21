package authorization

import (
	"crypto/hmac"
	"crypto/sha256"
	b64 "encoding/base64"
	"encoding/hex"
	"encoding/json"
	"hash"
	"os"
	"strconv"
	"strings"
	"time"
)

var JWT_HAEDER = map[string]string{
	"alg": "HS256",
	"typ": "JWT",
}

func sign(payload map[string]string, token_type string, token JWTToken) (string, error) {
	var secret string
	var crypto_f hash.Hash
	if token_type == "refresh" {
		payload["exp"] = strconv.FormatInt(token.RefreshExpire, 10)
		payload["typ"] = "refresh"
		secret = os.Getenv("REFRESH_SECRET")
	} else {
		payload["exp"] = strconv.FormatInt(token.AccessExpire, 10)
		payload["typ"] = "access"
		secret = os.Getenv("ACCESS_SECRET")
	}

	crypto_f = hmac.New(sha256.New, []byte(secret))

	header_data, err := json.Marshal(&JWT_HAEDER)
	if err != nil {
		return "Failed", err
	}

	payload_data, err := json.Marshal(&payload)
	if err != nil {
		return "Failed", err
	}

	header_encoede := b64.StdEncoding.EncodeToString(header_data)
	payload_encode := b64.StdEncoding.EncodeToString(payload_data)

	encoded := header_encoede + "." + payload_encode

	crypto_f.Write([]byte(encoded))

	return encoded + "." + hex.EncodeToString(crypto_f.Sum(nil)), nil
}

func GenerateToken(user_id uint64, username string) (*JWTToken, error) {
	token := &JWTToken{}
	token.AccessExpire = time.Now().Add(time.Minute * 20).Unix()
	token.RefreshExpire = time.Now().Add(time.Hour * 24).Unix()

	var jwt_payload map[string]string = make(map[string]string)
	jwt_payload["id"] = strconv.FormatUint(user_id, 10)
	jwt_payload["name"] = username

	if access_token, err := sign(jwt_payload, "access", *token); err != nil {
		return nil, err
	} else {
		token.AccessToken = access_token
	}
	if referesh_token, err := sign(jwt_payload, "refresh", *token); err != nil {
		return nil, err
	} else {
		token.RefreshToken = referesh_token
	}

	return token, nil
}

func extractTokenSection(token_sections []string, section int) (map[string]string, error) {
	binary, err := b64.StdEncoding.DecodeString(token_sections[section])

	if err != nil {
		return nil, err
	}

	var data map[string]string = make(map[string]string)
	if err := json.Unmarshal(binary, data); err != nil {
		return nil, err
	}
	return data, nil
}

func VerifySignature(token string) bool {
	sections := strings.Split(token, ".")

	header_data, err := extractTokenSection(sections, 0)
	if err != nil {
		print(err)
		return false
	}

	if header_data["alg"] != "HS256" {
		return false
	}

	payload_binary, err := extractTokenSection(sections, 1)
	if err != nil {
		print(err)
		return false
	}

	var secret string
	if payload_binary["typ"] == "access" {
		secret = os.Getenv("ACCESS_SECRET")

	} else if payload_binary["typ"] == "refresh" {
		secret = os.Getenv("REFRESH_SECRET")

	} else {
		return false
	}

	crypto_f := hmac.New(sha256.New, []byte(secret))

	encoded := sections[0] + "." + sections[1]

	crypto_f.Write([]byte(encoded))

	return hex.EncodeToString(crypto_f.Sum(nil)) == sections[2]
}
