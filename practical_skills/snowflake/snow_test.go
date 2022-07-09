package snowflake

import (
	"testing"
)

func TestGenerateId(t *testing.T) {
	t.Log(New(1).GenerateId().Int64())
}
