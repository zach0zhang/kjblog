package kjconfig

import (
	"fmt"
	"testing"
)

func TestKjconfig(t *testing.T) {
	InitConfig("config.json")
	fmt.Println(Cfg)
}
