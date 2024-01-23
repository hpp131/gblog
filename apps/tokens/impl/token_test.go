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
	req := tokens.NewRevokeTokenRequest("cmn41fosfi0usl19o7jg")
	err := tl.RevokeToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
}


func TestValidateToken(t *testing.T){
	req := tokens.NewValidateTokenRequest("cmnl1posfi0keb0b82jg", "cmnl1posfi0keb0b82k0")
	tk, err := tl.ValidateToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk)
}