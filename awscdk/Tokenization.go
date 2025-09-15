package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Less oft-needed functions to manipulate Tokens.
type Tokenization interface {
}

// The jsii proxy struct for Tokenization
type jsiiProxy_Tokenization struct {
	_ byte // padding
}

// Return whether the given object is an IResolvable object.
//
// This is different from Token.isUnresolved() which will also check for
// encoded Tokens, whereas this method will only do a type check on the given
// object.
func Tokenization_IsResolvable(obj interface{}) *bool {
	_init_.Initialize()

	if err := validateTokenization_IsResolvableParameters(obj); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Tokenization",
		"isResolvable",
		[]interface{}{obj},
		&returns,
	)

	return returns
}

// Resolves an object by evaluating all tokens and removing any undefined or empty objects or arrays.
//
// Values can only be primitives, arrays or tokens. Other objects (i.e. with methods) will be rejected.
func Tokenization_Resolve(obj interface{}, options *ResolveOptions) interface{} {
	_init_.Initialize()

	if err := validateTokenization_ResolveParameters(obj, options); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Tokenization",
		"resolve",
		[]interface{}{obj, options},
		&returns,
	)

	return returns
}

// Reverse any value into a Resolvable, if possible.
//
// In case of a string, the string must not be a concatenation.
func Tokenization_Reverse(x interface{}, options *ReverseOptions) IResolvable {
	_init_.Initialize()

	if err := validateTokenization_ReverseParameters(x, options); err != nil {
		panic(err)
	}
	var returns IResolvable

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Tokenization",
		"reverse",
		[]interface{}{x, options},
		&returns,
	)

	return returns
}

// Un-encode a string which is either a complete encoded token, or doesn't contain tokens at all.
//
// It's illegal for the string to be a concatenation of an encoded token and something else.
func Tokenization_ReverseCompleteString(s *string) IResolvable {
	_init_.Initialize()

	if err := validateTokenization_ReverseCompleteStringParameters(s); err != nil {
		panic(err)
	}
	var returns IResolvable

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Tokenization",
		"reverseCompleteString",
		[]interface{}{s},
		&returns,
	)

	return returns
}

// Un-encode a Tokenized value from a list.
func Tokenization_ReverseList(l *[]*string) IResolvable {
	_init_.Initialize()

	if err := validateTokenization_ReverseListParameters(l); err != nil {
		panic(err)
	}
	var returns IResolvable

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Tokenization",
		"reverseList",
		[]interface{}{l},
		&returns,
	)

	return returns
}

// Un-encode a Tokenized value from a number.
func Tokenization_ReverseNumber(n *float64) IResolvable {
	_init_.Initialize()

	if err := validateTokenization_ReverseNumberParameters(n); err != nil {
		panic(err)
	}
	var returns IResolvable

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Tokenization",
		"reverseNumber",
		[]interface{}{n},
		&returns,
	)

	return returns
}

// Un-encode a string potentially containing encoded tokens.
func Tokenization_ReverseString(s *string) TokenizedStringFragments {
	_init_.Initialize()

	if err := validateTokenization_ReverseStringParameters(s); err != nil {
		panic(err)
	}
	var returns TokenizedStringFragments

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Tokenization",
		"reverseString",
		[]interface{}{s},
		&returns,
	)

	return returns
}

// Stringify a number directly or lazily if it's a Token.
//
// If it is an object (i.e., { Ref: 'SomeLogicalId' }), return it as-is.
func Tokenization_StringifyNumber(x *float64) *string {
	_init_.Initialize()

	if err := validateTokenization_StringifyNumberParameters(x); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Tokenization",
		"stringifyNumber",
		[]interface{}{x},
		&returns,
	)

	return returns
}

