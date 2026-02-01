package redis

import (
	"context"
	"encoding/json"
	"time"
)

type TempUser struct {
	Nickname   string `json:"nickname"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	VerifyCode string `json:"verify_code"`
}

func SaveTempUser(tempUser TempUser) error {
	var bytes []byte
	var err error
	if bytes, err = json.Marshal(tempUser); err != nil {
		return err
	}

	ctx := context.Background()
	if err = Client.Set(ctx, tempUser.VerifyCode, string(bytes), time.Minute*5).Err(); err != nil {
		return err
	}
	return nil
}

func GetTempUser(verifyCode string) (*TempUser, error) {
	ctx := context.Background()

	result, err := Client.GetDel(ctx, verifyCode).Result()
	if err != nil {
		return nil, err
	}

	var tempUser TempUser
	if err = json.Unmarshal([]byte(result), &tempUser); err != nil {
		return nil, err
	}

	return &tempUser, nil
}
