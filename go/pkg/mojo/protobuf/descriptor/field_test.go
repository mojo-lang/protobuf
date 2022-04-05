package descriptor

import (
    "github.com/mojo-lang/core/go/pkg/mojo"
    "github.com/stretchr/testify/assert"
    "testing"
)

func newField() *Field {
    file := NewFile()
    message := NewMessage(file)
    return NewField(message, "foo")
}

func TestField_SetBoolOption(t *testing.T) {
    field := newField()

    field.SetOption(mojo.E_Alias, "alias")
    assert.Equal(t, "alias", field.GetStringOption(mojo.E_Alias))

    field.SetBoolOption(mojo.E_DbIgnore, true)
    assert.True(t, field.GetBoolOption(mojo.E_DbIgnore))
}

func TestField_SetNumber(t *testing.T) {
    field := newField()
    field.SetNumber(14)
    assert.Equal(t, int32(14), field.GetNumber())
}
