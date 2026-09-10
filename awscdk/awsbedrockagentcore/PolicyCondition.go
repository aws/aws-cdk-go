package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A condition on a policy statement.
//
// A condition compares a request attribute against a value. Conditions are grouped
// into the `when` and `unless` clauses of a statement, where the members of a clause
// must all hold.
//
// Comparisons are named for the type of value they accept, so each one takes a
// concrete type rather than a union. Use `allOf` and `anyOf` to build a nested
// boolean expression, which also makes the grouping explicit in the generated Cedar.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   // principal.department == "Engineering"
//   awscdk.PolicyCondition_StringEquals(awscdk.PolicyAttribute_Principal(jsii.String("department")), jsii.String("Engineering"))
//
//   // (principal.department == "Engineering" || principal.department == "Support")
//   awscdk.PolicyCondition_AnyOf([]PolicyCondition{
//   	awscdk.PolicyCondition_StringEquals(awscdk.PolicyAttribute_Principal(jsii.String("department")), jsii.String("Engineering")),
//   	awscdk.PolicyCondition_StringEquals(awscdk.PolicyAttribute_Principal(jsii.String("department")), jsii.String("Support")),
//   })
//
type PolicyCondition interface {
}

// The jsii proxy struct for PolicyCondition
type jsiiProxy_PolicyCondition struct {
	_ byte // padding
}

// All of the given conditions must hold.
//
// Renders as a parenthesised `&&` group, so it can be nested inside `anyOf`
// without relying on operator precedence.
func PolicyCondition_AllOf(conditions *[]PolicyCondition) PolicyCondition {
	_init_.Initialize()

	if err := validatePolicyCondition_AllOfParameters(conditions); err != nil {
		panic(err)
	}
	var returns PolicyCondition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.PolicyCondition",
		"allOf",
		[]interface{}{conditions},
		&returns,
	)

	return returns
}

// At least one of the given conditions must hold.
//
// Renders as a parenthesised `||` group, so it can be nested inside `allOf` or
// combined with the surrounding clause without relying on operator precedence.
func PolicyCondition_AnyOf(conditions *[]PolicyCondition) PolicyCondition {
	_init_.Initialize()

	if err := validatePolicyCondition_AnyOfParameters(conditions); err != nil {
		panic(err)
	}
	var returns PolicyCondition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.PolicyCondition",
		"anyOf",
		[]interface{}{conditions},
		&returns,
	)

	return returns
}

// The attribute equals a boolean value.
func PolicyCondition_BooleanEquals(attribute PolicyAttribute, value *bool) PolicyCondition {
	_init_.Initialize()

	if err := validatePolicyCondition_BooleanEqualsParameters(attribute, value); err != nil {
		panic(err)
	}
	var returns PolicyCondition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.PolicyCondition",
		"booleanEquals",
		[]interface{}{attribute, value},
		&returns,
	)

	return returns
}

// The attribute is an IP address inside the given CIDR range.
func PolicyCondition_IpInRange(attribute PolicyAttribute, cidr *string) PolicyCondition {
	_init_.Initialize()

	if err := validatePolicyCondition_IpInRangeParameters(attribute, cidr); err != nil {
		panic(err)
	}
	var returns PolicyCondition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.PolicyCondition",
		"ipInRange",
		[]interface{}{attribute, cidr},
		&returns,
	)

	return returns
}

// The attribute equals a number value.
//
// Cedar whole numbers are 64-bit signed integers, so the value must be an integer.
func PolicyCondition_NumberEquals(attribute PolicyAttribute, value *float64) PolicyCondition {
	_init_.Initialize()

	if err := validatePolicyCondition_NumberEqualsParameters(attribute, value); err != nil {
		panic(err)
	}
	var returns PolicyCondition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.PolicyCondition",
		"numberEquals",
		[]interface{}{attribute, value},
		&returns,
	)

	return returns
}

// The attribute is greater than a number value.
func PolicyCondition_NumberGreaterThan(attribute PolicyAttribute, value *float64) PolicyCondition {
	_init_.Initialize()

	if err := validatePolicyCondition_NumberGreaterThanParameters(attribute, value); err != nil {
		panic(err)
	}
	var returns PolicyCondition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.PolicyCondition",
		"numberGreaterThan",
		[]interface{}{attribute, value},
		&returns,
	)

	return returns
}

// The attribute is greater than or equal to a number value.
func PolicyCondition_NumberGreaterThanOrEquals(attribute PolicyAttribute, value *float64) PolicyCondition {
	_init_.Initialize()

	if err := validatePolicyCondition_NumberGreaterThanOrEqualsParameters(attribute, value); err != nil {
		panic(err)
	}
	var returns PolicyCondition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.PolicyCondition",
		"numberGreaterThanOrEquals",
		[]interface{}{attribute, value},
		&returns,
	)

	return returns
}

// The attribute is one of the given number values.
func PolicyCondition_NumberIn(attribute PolicyAttribute, values *[]*float64) PolicyCondition {
	_init_.Initialize()

	if err := validatePolicyCondition_NumberInParameters(attribute, values); err != nil {
		panic(err)
	}
	var returns PolicyCondition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.PolicyCondition",
		"numberIn",
		[]interface{}{attribute, values},
		&returns,
	)

	return returns
}

// The attribute is less than a number value.
func PolicyCondition_NumberLessThan(attribute PolicyAttribute, value *float64) PolicyCondition {
	_init_.Initialize()

	if err := validatePolicyCondition_NumberLessThanParameters(attribute, value); err != nil {
		panic(err)
	}
	var returns PolicyCondition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.PolicyCondition",
		"numberLessThan",
		[]interface{}{attribute, value},
		&returns,
	)

	return returns
}

// The attribute is less than or equal to a number value.
func PolicyCondition_NumberLessThanOrEquals(attribute PolicyAttribute, value *float64) PolicyCondition {
	_init_.Initialize()

	if err := validatePolicyCondition_NumberLessThanOrEqualsParameters(attribute, value); err != nil {
		panic(err)
	}
	var returns PolicyCondition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.PolicyCondition",
		"numberLessThanOrEquals",
		[]interface{}{attribute, value},
		&returns,
	)

	return returns
}

// The attribute does not equal a number value.
func PolicyCondition_NumberNotEquals(attribute PolicyAttribute, value *float64) PolicyCondition {
	_init_.Initialize()

	if err := validatePolicyCondition_NumberNotEqualsParameters(attribute, value); err != nil {
		panic(err)
	}
	var returns PolicyCondition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.PolicyCondition",
		"numberNotEquals",
		[]interface{}{attribute, value},
		&returns,
	)

	return returns
}

// The attribute is a set that contains the given value.
//
// Use this when the attribute itself holds a set, for example `principal.groups`.
// To test a scalar attribute against a list of allowed values, use `stringIn` or
// `numberIn` instead.
func PolicyCondition_SetContains(attribute PolicyAttribute, value *string) PolicyCondition {
	_init_.Initialize()

	if err := validatePolicyCondition_SetContainsParameters(attribute, value); err != nil {
		panic(err)
	}
	var returns PolicyCondition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.PolicyCondition",
		"setContains",
		[]interface{}{attribute, value},
		&returns,
	)

	return returns
}

// The attribute equals a string value.
func PolicyCondition_StringEquals(attribute PolicyAttribute, value *string) PolicyCondition {
	_init_.Initialize()

	if err := validatePolicyCondition_StringEqualsParameters(attribute, value); err != nil {
		panic(err)
	}
	var returns PolicyCondition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.PolicyCondition",
		"stringEquals",
		[]interface{}{attribute, value},
		&returns,
	)

	return returns
}

// The attribute is one of the given string values.
func PolicyCondition_StringIn(attribute PolicyAttribute, values *[]*string) PolicyCondition {
	_init_.Initialize()

	if err := validatePolicyCondition_StringInParameters(attribute, values); err != nil {
		panic(err)
	}
	var returns PolicyCondition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.PolicyCondition",
		"stringIn",
		[]interface{}{attribute, values},
		&returns,
	)

	return returns
}

// The attribute does not equal a string value.
func PolicyCondition_StringNotEquals(attribute PolicyAttribute, value *string) PolicyCondition {
	_init_.Initialize()

	if err := validatePolicyCondition_StringNotEqualsParameters(attribute, value); err != nil {
		panic(err)
	}
	var returns PolicyCondition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.PolicyCondition",
		"stringNotEquals",
		[]interface{}{attribute, value},
		&returns,
	)

	return returns
}

