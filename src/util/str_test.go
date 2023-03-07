package util

import (
	"testing"
)

func TestGetValueFromBracket(t *testing.T) {
	res := GetValueFromBracket("{{console}}{{unemeta_admin:UneUne202.}}{{34.85.98.88:3306}}")
	if len(res) < 3 {
		t.Fatal(res)
	}
	if res[0] != "console" {
		t.Fatal("expect console, actual:", res[0])
	}
	if res[1] != "unemeta_admin:UneUne202." {
		t.Fatalf("expect %s, actual:%s", "unemeta_admin:UneUne202.", res[1])
	}
	if res[2] != "34.85.98.88:3306" {
		t.Fatalf("expect %s, actual:%s", "34.85.98.88:3306", res[2])
	}
}
