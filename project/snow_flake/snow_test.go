package snow_flake

import (
	"fmt"
	"testing"
)

func TestGenerateId(t *testing.T) {
	SetNode(1)
	id := GenerateId()
	fmt.Println(id.Int64())
}
