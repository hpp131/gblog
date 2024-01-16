package conf_test

import (
	"fmt"
	"testing"

	"github.com/hpp131/gblog/conf"
)

func TestDB(t *testing.T){
	conn := conf.C().DB()
	fmt.Printf("%#v", conn)
}


