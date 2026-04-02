package awsmediapackagev2alpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediapackagev2alpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Represents a bitrate filter expression segment — either a single value or a range.
//
// Use with `ManifestFilter.bitrateCombo()` to build filters that combine
// ranges and individual bitrate values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediapackagev2_alpha "github.com/aws/aws-cdk-go/awsmediapackagev2alpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var bitrate Bitrate
//
//   bitrateExpression := mediapackagev2_alpha.BitrateExpression_Range(bitrate, bitrate)
//
// Experimental.
type BitrateExpression interface {
}

// The jsii proxy struct for BitrateExpression
type jsiiProxy_BitrateExpression struct {
	_ byte // padding
}

// Experimental.
func NewBitrateExpression(_expression *string, _values *[]*float64) BitrateExpression {
	_init_.Initialize()

	if err := validateNewBitrateExpressionParameters(_expression, _values); err != nil {
		panic(err)
	}
	j := jsiiProxy_BitrateExpression{}

	_jsii_.Create(
		"@aws-cdk/aws-mediapackagev2-alpha.BitrateExpression",
		[]interface{}{_expression, _values},
		&j,
	)

	return &j
}

// Experimental.
func NewBitrateExpression_Override(b BitrateExpression, _expression *string, _values *[]*float64) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-mediapackagev2-alpha.BitrateExpression",
		[]interface{}{_expression, _values},
		b,
	)
}

// An inclusive bitrate range.
// Experimental.
func BitrateExpression_Range(start awscdk.Bitrate, end awscdk.Bitrate) BitrateExpression {
	_init_.Initialize()

	if err := validateBitrateExpression_RangeParameters(start, end); err != nil {
		panic(err)
	}
	var returns BitrateExpression

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediapackagev2-alpha.BitrateExpression",
		"range",
		[]interface{}{start, end},
		&returns,
	)

	return returns
}

// A single bitrate value.
// Experimental.
func BitrateExpression_Value(v awscdk.Bitrate) BitrateExpression {
	_init_.Initialize()

	if err := validateBitrateExpression_ValueParameters(v); err != nil {
		panic(err)
	}
	var returns BitrateExpression

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediapackagev2-alpha.BitrateExpression",
		"value",
		[]interface{}{v},
		&returns,
	)

	return returns
}

