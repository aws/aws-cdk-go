package awscdkintegtestsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkintegtestsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Partial and special matching during assertions.
// Experimental.
type Match interface {
}

// The jsii proxy struct for Match
type jsiiProxy_Match struct {
	_ byte // padding
}

// Experimental.
func NewMatch_Override(m Match) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/integ-tests-alpha.Match",
		nil, // no parameters
		m,
	)
}

// Matches the specified pattern with the array found in the same relative path of the target.
//
// The set of elements (or matchers) must be in the same order as would be found.
// Experimental.
func Match_ArrayWith(pattern *[]interface{}) *map[string]*[]interface{} {
	_init_.Initialize()

	if err := validateMatch_ArrayWithParameters(pattern); err != nil {
		panic(err)
	}
	var returns *map[string]*[]interface{}

	_jsii_.StaticInvoke(
		"@aws-cdk/integ-tests-alpha.Match",
		"arrayWith",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

// Matches the specified pattern to an object found in the same relative path of the target.
//
// The keys and their values (or matchers) must be present in the target but the target can be a superset.
// Experimental.
func Match_ObjectLike(pattern *map[string]interface{}) *map[string]*map[string]interface{} {
	_init_.Initialize()

	if err := validateMatch_ObjectLikeParameters(pattern); err != nil {
		panic(err)
	}
	var returns *map[string]*map[string]interface{}

	_jsii_.StaticInvoke(
		"@aws-cdk/integ-tests-alpha.Match",
		"objectLike",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

// Matches any string-encoded JSON and applies the specified pattern after parsing it.
// Experimental.
func Match_SerializedJson(pattern *map[string]interface{}) *map[string]*map[string]interface{} {
	_init_.Initialize()

	if err := validateMatch_SerializedJsonParameters(pattern); err != nil {
		panic(err)
	}
	var returns *map[string]*map[string]interface{}

	_jsii_.StaticInvoke(
		"@aws-cdk/integ-tests-alpha.Match",
		"serializedJson",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

// Matches targets according to a regular expression.
// Experimental.
func Match_StringLikeRegexp(pattern *string) *map[string]*string {
	_init_.Initialize()

	if err := validateMatch_StringLikeRegexpParameters(pattern); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.StaticInvoke(
		"@aws-cdk/integ-tests-alpha.Match",
		"stringLikeRegexp",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

