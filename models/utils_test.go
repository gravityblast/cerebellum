package models

import (
  "testing"
  assert "github.com/pilu/miniassert"
)

func TestIsValidUUID(t *testing.T) {
  assert.True(t, IsValidUUID("77a591d0-ed7c-0130-97ce-28cfe91367b5"))
  assert.False(t, IsValidUUID("77a591d0-ed7c-0130-97ce-28cfe91367b5i-XXX"))
  assert.False(t, IsValidUUID("bad UUI"))
}
