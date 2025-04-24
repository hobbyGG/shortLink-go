package urlp

import (
	"testing"
)

type testArg struct {
	longURL string
}

func TestGetLastURL(t *testing.T) {
	// 测试用例
	tests := []struct {
		name     string
		args     testArg
		expected string
	}{
		{name: "正常示例", args: testArg{longURL: "abc.cn/asd"}, expected: "asd"},
		{name: "带query的测试", args: testArg{longURL: "abc.cn/asd?query=query"}, expected: "asd"},
		{name: "/结尾的测试", args: testArg{longURL: "abc.cn/asd/"}, expected: "asd"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			get, err := GetLastPath(test.args.longURL)
			if err != nil {
				t.Error(err)
				return
			}
			if get != test.expected {
				t.Errorf("get %v, expected %v", get, test.expected)
				return
			}
		})
	}
}
