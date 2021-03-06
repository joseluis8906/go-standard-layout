package jwt_test

import (
	"context"
	"testing"
	"time"

	"github.com/joseluis8906/go-standard-layout/pkg/jwt"
)

func TestEncodeDecode(t *testing.T) {
	ctx := context.Background()

	type Data struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	data := Data{
		ID:   "123-abc",
		Name: "John Doe",
	}

	secret := "my secret"

	token, err := jwt.Encode(ctx, secret, data, time.Duration(time.Now().Unix()+3))
	if err != nil {
		t.Error(err)
	}

	time.Sleep(2 * time.Second)

	var data2 Data

	err = jwt.Decode(ctx, secret, token, &data2)
	if err != nil {
		t.Error(err)
	}

	if data2 != data {
		t.Error("values don't match")
	}

}
