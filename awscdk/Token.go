package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a special or lazily-evaluated value.
//
// Can be used to delay evaluation of a certain value in case, for example,
// that it requires some context or late-bound data. Can also be used to
// mark values that need special processing at document rendering time.
//
// Tokens can be embedded into strings while retaining their original
// semantics.
type Token interface {
}

// The jsii proxy struct for Token
type jsiiProxy_Token struct {
	_ byte // padding
}

// Return a resolvable representation of the given value.
func Token_AsAny(value interface{}) IResolvable {
	_init_.Initialize()

	if err := validateToken_AsAnyParameters(value); err != nil {
		panic(err)
	}
	var returns IResolvable

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Token",
		"asAny",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Return a reversible list representation of this token.
func Token_AsList(value interface{}, options *EncodingOptions) *[]*string {
	_init_.Initialize()

	if err := validateToken_AsListParameters(value, options); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Token",
		"asList",
		[]interface{}{value, options},
		&returns,
	)

	return returns
}

// Return a reversible number representation of this token.
func Token_AsNumber(value interface{}) *float64 {
	_init_.Initialize()

	if err := validateToken_AsNumberParameters(value); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Token",
		"asNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Return a reversible string representation of this token.
//
// If the Token is initialized with a literal, the stringified value of the
// literal is returned. Otherwise, a special quoted string representation
// of the Token is returned that can be embedded into other strings.
//
// Strings with quoted Tokens in them can be restored back into
// complex values with the Tokens restored by calling `resolve()`
// on the string.
func Token_AsString(value interface{}, options *EncodingOptions) *string {
	_init_.Initialize()

	if err := validateToken_AsStringParameters(value, options); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Token",
		"asString",
		[]interface{}{value, options},
		&returns,
	)

	return returns
}

// Compare two strings that might contain Tokens with each other.
func Token_CompareStrings(possibleToken1 *string, possibleToken2 *string) TokenComparison {
	_init_.Initialize()

	if err := validateToken_CompareStringsParameters(possibleToken1, possibleToken2); err != nil {
		panic(err)
	}
	var returns TokenComparison

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Token",
		"compareStrings",
		[]interface{}{possibleToken1, possibleToken2},
		&returns,
	)

	return returns
}

// Returns true if obj represents an unresolved value.
//
// One of these must be true:
//
// - `obj` is an IResolvable
// - `obj` is a string containing at least one encoded `IResolvable`
// - `obj` is either an encoded number or list
//
// This does NOT recurse into lists or objects to see if they
// contain resolvables.
func Token_IsUnresolved(obj interface{}) *bool {
	_init_.Initialize()

	if err := validateToken_IsUnresolvedParameters(obj); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Token",
		"isUnresolved",
		[]interface{}{obj},
		&returns,
	)

	return returns
}

