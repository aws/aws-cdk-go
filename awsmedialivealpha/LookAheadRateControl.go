package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Lookahead rate control.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   lookAheadRateControl := medialive_alpha.LookAheadRateControl_HIGH()
//
// Experimental.
type LookAheadRateControl interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for LookAheadRateControl
type jsiiProxy_LookAheadRateControl struct {
	_ byte // padding
}

func (j *jsiiProxy_LookAheadRateControl) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// A value not yet modelled by AWS CDK.
// Experimental.
func LookAheadRateControl_Of(value *string) LookAheadRateControl {
	_init_.Initialize()

	if err := validateLookAheadRateControl_OfParameters(value); err != nil {
		panic(err)
	}
	var returns LookAheadRateControl

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.LookAheadRateControl",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func LookAheadRateControl_HIGH() LookAheadRateControl {
	_init_.Initialize()
	var returns LookAheadRateControl
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.LookAheadRateControl",
		"HIGH",
		&returns,
	)
	return returns
}

func LookAheadRateControl_LOW() LookAheadRateControl {
	_init_.Initialize()
	var returns LookAheadRateControl
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.LookAheadRateControl",
		"LOW",
		&returns,
	)
	return returns
}

func LookAheadRateControl_MEDIUM() LookAheadRateControl {
	_init_.Initialize()
	var returns LookAheadRateControl
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.LookAheadRateControl",
		"MEDIUM",
		&returns,
	)
	return returns
}

