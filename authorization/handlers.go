package authorization

import (
	"crypto/hmac"
	"crypto/sha256"
	b64 "encoding/base64"
	"encoding/hex"
	"encoding/json"
	"os"
	"strconv"
	"time"
)

var JWT_HAEDER = map[string]string{
	"alg": "HS256",
	"typ": "JWT",
}

func GenerateToken(user_id uint64, username string) (*JWTToken, error) {
	token := &JWTToken{}
	token.AccessExpire = time.Now().Add(time.Minute * 20).Unix()
	token.RefreshExpire = time.Now().Add(time.Hour * 24).Unix()

	var jwt_payload map[string]string = make(map[string]string)
	jwt_payload["id"] = strconv.FormatUint(user_id, 10)
	jwt_payload["name"] = username
	jwt_payload["Aexp"] = strconv.FormatInt(token.AccessExpire, 10)
	jwt_payload["Rexp"] = strconv.FormatInt(token.RefreshExpire, 10)

	header_data, err := json.Marshal(&JWT_HAEDER)
	if err != nil {
		return nil, err
	}

	payload_data, err := json.Marshal(&jwt_payload)
	if err != nil {
		return nil, err
	}

	header_encoede := b64.StdEncoding.EncodeToString(header_data)
	payload_encode := b64.StdEncoding.EncodeToString(payload_data)

	encoded := header_encoede + "." + payload_encode

	access_secret := os.Getenv("ACCESS_SECRET")
	refresh_secret := os.Getenv("REFRESH_SECRET")

	access_crypto := hmac.New(sha256.New, []byte(access_secret))
	refresh_crypto := hmac.New(sha256.New, []byte(refresh_secret))

	access_crypto.Write([]byte(encoded))
	refresh_crypto.Write([]byte(encoded))

	token.AccessToken = encoded + "." + hex.EncodeToString(access_crypto.Sum(nil))
	token.RefreshToken = encoded + "." + hex.EncodeToString(refresh_crypto.Sum(nil))
	return token, nil
}
