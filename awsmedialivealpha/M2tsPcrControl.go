package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Controls insertion of the Program Clock Reference (PCR).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   m2tsPcrControl := medialive_alpha.M2tsPcrControl_Of(jsii.String("value"))
//
// Experimental.
type M2tsPcrControl interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for M2tsPcrControl
type jsiiProxy_M2tsPcrControl struct {
	_ byte // padding
}

func (j *jsiiProxy_M2tsPcrControl) Value() *string {
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
func M2tsPcrControl_Of(value *string) M2tsPcrControl {
	_init_.Initialize()

	if err := validateM2tsPcrControl_OfParameters(value); err != nil {
		panic(err)
	}
	var returns M2tsPcrControl

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.M2tsPcrControl",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func M2tsPcrControl_CONFIGURED_PCR_PERIOD() M2tsPcrControl {
	_init_.Initialize()
	var returns M2tsPcrControl
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M2tsPcrControl",
		"CONFIGURED_PCR_PERIOD",
		&returns,
	)
	return returns
}

func M2tsPcrControl_PCR_EVERY_PES_PACKET() M2tsPcrControl {
	_init_.Initialize()
	var returns M2tsPcrControl
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M2tsPcrControl",
		"PCR_EVERY_PES_PACKET",
		&returns,
	)
	return returns
}

