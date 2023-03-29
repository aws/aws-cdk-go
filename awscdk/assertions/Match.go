package assertions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Partial and special matching during template assertions.
type Match interface {
}

// The jsii proxy struct for Match
type jsiiProxy_Match struct {
	_ byte // padding
}

func NewMatch_Override(m Match) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.assertions.Match",
		nil, // no parameters
		m,
	)
}

// Use this matcher in the place of a field's value, if the field must not be present.
func Match_Absent() Matcher {
	_init_.Initialize()

	var returns Matcher

	_jsii_.StaticInvoke(
		"aws-cdk-lib.assertions.Match",
		"absent",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Matches any non-null value at the target.
func Match_AnyValue() Matcher {
	_init_.Initialize()

	var returns Matcher

	_jsii_.StaticInvoke(
		"aws-cdk-lib.assertions.Match",
		"anyValue",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Matches the specified pattern with the array found in the same relative path of the target.
//
// The set of elements (or matchers) must match exactly and in order.
func Match_ArrayEquals(pattern *[]interface{}) Matcher {
	_init_.Initialize()

	if err := validateMatch_ArrayEqualsParameters(pattern); err != nil {
		panic(err)
	}
	var returns Matcher

	_jsii_.StaticInvoke(
		"aws-cdk-lib.assertions.Match",
		"arrayEquals",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

// Matches the specified pattern with the array found in the same relative path of the target.
//
// The set of elements (or matchers) must be in the same order as would be found.
func Match_ArrayWith(pattern *[]interface{}) Matcher {
	_init_.Initialize()

	if err := validateMatch_ArrayWithParameters(pattern); err != nil {
		panic(err)
	}
	var returns Matcher

	_jsii_.StaticInvoke(
		"aws-cdk-lib.assertions.Match",
		"arrayWith",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

// Deep exact matching of the specified pattern to the target.
func Match_Exact(pattern interface{}) Matcher {
	_init_.Initialize()

	if err := validateMatch_ExactParameters(pattern); err != nil {
		panic(err)
	}
	var returns Matcher

	_jsii_.StaticInvoke(
		"aws-cdk-lib.assertions.Match",
		"exact",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

// Matches any target which does NOT follow the specified pattern.
func Match_Not(pattern interface{}) Matcher {
	_init_.Initialize()

	if err := validateMatch_NotParameters(pattern); err != nil {
		panic(err)
	}
	var returns Matcher

	_jsii_.StaticInvoke(
		"aws-cdk-lib.assertions.Match",
		"not",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

// Matches the specified pattern to an object found in the same relative path of the target.
//
// The keys and their values (or matchers) must match exactly with the target.
func Match_ObjectEquals(pattern *map[string]interface{}) Matcher {
	_init_.Initialize()

	if err := validateMatch_ObjectEqualsParameters(pattern); err != nil {
		panic(err)
	}
	var returns Matcher

	_jsii_.StaticInvoke(
		"aws-cdk-lib.assertions.Match",
		"objectEquals",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

// Matches the specified pattern to an object found in the same relative path of the target.
//
// The keys and their values (or matchers) must be present in the target but the target can be a superset.
func Match_ObjectLike(pattern *map[string]interface{}) Matcher {
	_init_.Initialize()

	if err := validateMatch_ObjectLikeParameters(pattern); err != nil {
		panic(err)
	}
	var returns Matcher

	_jsii_.StaticInvoke(
		"aws-cdk-lib.assertions.Match",
		"objectLike",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

// Matches any string-encoded JSON and applies the specified pattern after parsing it.
func Match_SerializedJson(pattern interface{}) Matcher {
	_init_.Initialize()

	if err := validateMatch_SerializedJsonParameters(pattern); err != nil {
		panic(err)
	}
	var returns Matcher

	_jsii_.StaticInvoke(
		"aws-cdk-lib.assertions.Match",
		"serializedJson",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

// Matches targets according to a regular expression.
func Match_StringLikeRegexp(pattern *string) Matcher {
	_init_.Initialize()

	if err := validateMatch_StringLikeRegexpParameters(pattern); err != nil {
		panic(err)
	}
	var returns Matcher

	_jsii_.StaticInvoke(
		"aws-cdk-lib.assertions.Match",
		"stringLikeRegexp",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

