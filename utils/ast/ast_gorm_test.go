package ast

import (
	"testing"

	"github.com/Joword/chatbgi-manager/global"
	"github.com/Joword/chatbgi-manager/model/example"
)

const A = 123

func TestAddRegisterTablesAst(t *testing.T) {
	AddRegisterTablesAst("D:\\gin-vue-admin\\server\\utils\\ast_test.go", "Register", "test", "testDB", "testModel")
}

func Register() {
	test := global.GetGlobalDBByDBName("test")
	test.AutoMigrate(example.ExaFile{})
}
