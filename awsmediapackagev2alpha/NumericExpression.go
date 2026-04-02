package awsmediapackagev2alpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediapackagev2alpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a numeric filter expression segment — either a single value or a range.
//
// Use with `ManifestFilter.numericCombo()` to build filters that combine
// ranges and individual values (e.g. `video_height:240-360,720-1080,1440`).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediapackagev2_alpha "github.com/aws/aws-cdk-go/awsmediapackagev2alpha"
//
//   numericExpression := mediapackagev2_alpha.NumericExpression_Range(jsii.Number(123), jsii.Number(123))
//
// Experimental.
type NumericExpression interface {
}

// The jsii proxy struct for NumericExpression
type jsiiProxy_NumericExpression struct {
	_ byte // padding
}

// Experimental.
func NewNumericExpression(_expression *string, _values *[]*float64) NumericExpression {
	_init_.Initialize()

	if err := validateNewNumericExpressionParameters(_expression, _values); err != nil {
		panic(err)
	}
	j := jsiiProxy_NumericExpression{}

	_jsii_.Create(
		"@aws-cdk/aws-mediapackagev2-alpha.NumericExpression",
		[]interface{}{_expression, _values},
		&j,
	)

	return &j
}

// Experimental.
func NewNumericExpression_Override(n NumericExpression, _expression *string, _values *[]*float64) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-mediapackagev2-alpha.NumericExpression",
		[]interface{}{_expression, _values},
		n,
	)
}

// An inclusive numeric range.
// Experimental.
func NumericExpression_Range(start *float64, end *float64) NumericExpression {
	_init_.Initialize()

	if err := validateNumericExpression_RangeParameters(start, end); err != nil {
		panic(err)
	}
	var returns NumericExpression

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediapackagev2-alpha.NumericExpression",
		"range",
		[]interface{}{start, end},
		&returns,
	)

	return returns
}

// A single numeric value.
// Experimental.
func NumericExpression_Value(v *float64) NumericExpression {
	_init_.Initialize()

	if err := validateNumericExpression_ValueParameters(v); err != nil {
		panic(err)
	}
	var returns NumericExpression

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediapackagev2-alpha.NumericExpression",
		"value",
		[]interface{}{v},
		&returns,
	)

	return returns
}

