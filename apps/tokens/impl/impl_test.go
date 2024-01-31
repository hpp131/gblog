package impl_test

import (
	"context"
	"testing"

	"github.com/hpp131/gblog/apps/tokens"
	// "github.com/hpp131/gblog/apps/tokens/impl"
	_ "github.com/hpp131/gblog/apps"
	"github.com/hpp131/gblog/ioc"
)

var (
	ctx       = context.Background()
	tkImplObj tokens.Service
)

// 使用ioc的方式解决依赖；使用ioc重写单元测试
func init() {
	tkImplObj = ioc.Controller().Get(tokens.AppName).(tokens.Service)
	if err := ioc.Init(); err != nil {
		panic(err)
	}
}

func TestIssueToken(t *testing.T) {
	req := tokens.NewIssueTokenRequest("admin1", "123123")
	token, err := tkImplObj.IssueToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(token)
}

// func TestRevokeToken(t *testing.T){
// 	req := tokens.NewRevokeTokenRequest("cmn41fosfi0usl19o7jg")
// 	err := tl.RevokeToken(ctx, req)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }

// func TestValidateToken(t *testing.T){
// 	req := tokens.NewValidateTokenRequest("cmnl1posfi0keb0b82jg", "cmnl1posfi0keb0b82k0")
// 	tk, err := tl.ValidateToken(ctx, req)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Log(tk)
// }
