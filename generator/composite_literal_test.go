package generator

import (
	"testing"

	"github.com/moznion/gowrtr/internal/errmsg"

	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateCompositeLiteralBeSuccess(t *testing.T) {
	composeGenerator := NewCompositeLiteral("&Struct").
		AddField("foo", NewRawStatement(`"foo-value"`)).
		AddFieldStr("bar", "bar-value").
		AddField("buz", NewAnonymousFunc(
			false,
			NewAnonymousFuncSignature().ReturnTypes("bool"),
			NewReturnStatement("true"),
		).Invocation(NewFuncInvocation())).
		AddFieldRaw("qux", 12345).
		AddFieldRaw("foobar", false)

	{
		gen, err := composeGenerator.Generate(0)

		assert.NoError(t, err)
		expected := `&Struct{
	foo: "foo-value",
	bar: "bar-value",
	buz: func() bool {
		return true
	}(),
	qux: 12345,
	foobar: false,
}
`
		assert.Equal(t, expected, gen)
	}

	{
		gen, err := composeGenerator.Generate(2)

		assert.NoError(t, err)
		expected := `		&Struct{
			foo: "foo-value",
			bar: "bar-value",
			buz: func() bool {
				return true
			}(),
			qux: 12345,
			foobar: false,
		}
`
		assert.Equal(t, expected, gen)
	}
}

func TestShouldGenerateCompositeLiteralRaiseError(t *testing.T) {
	_, err := NewCompositeLiteral("").AddField("foo", NewIf("")).Generate(0)
	assert.EqualError(t, err, errmsg.IfConditionIsEmptyError().Error())
}
