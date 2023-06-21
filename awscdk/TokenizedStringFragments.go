package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Fragments of a concatenated string containing stringified Tokens.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   tokenizedStringFragments := cdk.NewTokenizedStringFragments()
//
type TokenizedStringFragments interface {
	FirstToken() IResolvable
	FirstValue() interface{}
	Length() *float64
	// Return all Tokens from this string.
	Tokens() *[]IResolvable
	AddIntrinsic(value interface{})
	AddLiteral(lit interface{})
	AddToken(token IResolvable)
	// Combine the string fragments using the given joiner.
	//
	// If there are any.
	Join(concat IFragmentConcatenator) interface{}
	// Apply a transformation function to all tokens in the string.
	MapTokens(mapper ITokenMapper) TokenizedStringFragments
}

// The jsii proxy struct for TokenizedStringFragments
type jsiiProxy_TokenizedStringFragments struct {
	_ byte // padding
}

func (j *jsiiProxy_TokenizedStringFragments) FirstToken() IResolvable {
	var returns IResolvable
	_jsii_.Get(
		j,
		"firstToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenizedStringFragments) FirstValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firstValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenizedStringFragments) Length() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"length",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenizedStringFragments) Tokens() *[]IResolvable {
	var returns *[]IResolvable
	_jsii_.Get(
		j,
		"tokens",
		&returns,
	)
	return returns
}


func NewTokenizedStringFragments() TokenizedStringFragments {
	_init_.Initialize()

	j := jsiiProxy_TokenizedStringFragments{}

	_jsii_.Create(
		"aws-cdk-lib.TokenizedStringFragments",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewTokenizedStringFragments_Override(t TokenizedStringFragments) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.TokenizedStringFragments",
		nil, // no parameters
		t,
	)
}

func (t *jsiiProxy_TokenizedStringFragments) AddIntrinsic(value interface{}) {
	if err := t.validateAddIntrinsicParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addIntrinsic",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TokenizedStringFragments) AddLiteral(lit interface{}) {
	if err := t.validateAddLiteralParameters(lit); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addLiteral",
		[]interface{}{lit},
	)
}

func (t *jsiiProxy_TokenizedStringFragments) AddToken(token IResolvable) {
	if err := t.validateAddTokenParameters(token); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addToken",
		[]interface{}{token},
	)
}

func (t *jsiiProxy_TokenizedStringFragments) Join(concat IFragmentConcatenator) interface{} {
	if err := t.validateJoinParameters(concat); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"join",
		[]interface{}{concat},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TokenizedStringFragments) MapTokens(mapper ITokenMapper) TokenizedStringFragments {
	if err := t.validateMapTokensParameters(mapper); err != nil {
		panic(err)
	}
	var returns TokenizedStringFragments

	_jsii_.Invoke(
		t,
		"mapTokens",
		[]interface{}{mapper},
		&returns,
	)

	return returns
}

