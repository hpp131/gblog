package impl_test

import (
	"context"
	"testing"

	"github.com/hpp131/gblog/apps/tokens"
	"github.com/hpp131/gblog/apps/tokens/impl"
)

var (
	ctx = context.Background()
	tl = impl.NewTokenServiceImpl()
)

func TestIssueToken(t *testing.T){
	req := tokens.NewIssueTokenRequest("admin1", "123123")
	token, err := tl.IssueToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(token)
}


func TestRevokeToken(t *testing.T){

	
}


func TestValidateToken(t *testing.T){

	
}