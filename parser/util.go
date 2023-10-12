package parser

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"referral-system/constant"
	"strconv"

	"github.com/gorilla/mux"
)

type clientRequest struct {
	req *http.Request
}

func (client clientRequest) param(s string) string {
	param := client.req.URL.Query().Get(s)
	return param
}

func (client clientRequest) paramInt(s string) int {
	param := client.req.URL.Query().Get(s)
	if param != "" {
		v, err := strconv.Atoi(param)
		if err != nil {
			return constant.IntParamEmpty
		}
		return v
	}
	return constant.IntParamEmpty
}

func (client clientRequest) pathParamInt(s string) int {
	vars := mux.Vars(client.req)
	if vars[s] == "" {
		return constant.IntParamEmpty
	}

	v, err := strconv.Atoi(vars[s])
	if err != nil {
		return constant.IntParamEmpty
	}

	return v
}

func (client clientRequest) pathParam(s string) string {
	vars := mux.Vars(client.req)
	return vars[s]
}

func (client clientRequest) getEntityID() int {
	param := client.paramInt("id")
	path := client.pathParamInt("id")

	if path != constant.IntParamEmpty {
		return path
	}

	return param
}

func jsonParser(r *http.Request, v interface{}) error {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&v); err != nil {
		return err
	}

	return nil
}

func getUserID(r *http.Request) int {
	var userID string

	cookies := r.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == "userID" {
			userID = cookie.Value
		}
	}

	userIDInt, _ := strconv.Atoi(userID)
	return userIDInt
}

func HashPassword(password string) string {
	hasher := md5.New()
	hasher.Write([]byte(password))
	hashedBytes := hasher.Sum(nil)
	hashedString := hex.EncodeToString(hashedBytes)

	return hashedString
}

func EntityStringify(v interface{}) string {
	byteData, err := json.Marshal(v)
	if err != nil {
		fmt.Println("e:jsonMarshal", byteData)
	}

	return string(byteData)
}
