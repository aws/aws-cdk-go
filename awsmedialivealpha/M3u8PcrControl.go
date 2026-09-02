package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Controls insertion of the Program Clock Reference (PCR) in an M3U8 container.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   m3u8PcrControl := medialive_alpha.M3u8PcrControl_Of(jsii.String("value"))
//
// Experimental.
type M3u8PcrControl interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for M3u8PcrControl
type jsiiProxy_M3u8PcrControl struct {
	_ byte // padding
}

func (j *jsiiProxy_M3u8PcrControl) Value() *string {
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
func M3u8PcrControl_Of(value *string) M3u8PcrControl {
	_init_.Initialize()

	if err := validateM3u8PcrControl_OfParameters(value); err != nil {
		panic(err)
	}
	var returns M3u8PcrControl

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.M3u8PcrControl",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func M3u8PcrControl_CONFIGURED_PCR_PERIOD() M3u8PcrControl {
	_init_.Initialize()
	var returns M3u8PcrControl
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M3u8PcrControl",
		"CONFIGURED_PCR_PERIOD",
		&returns,
	)
	return returns
}

func M3u8PcrControl_PCR_EVERY_PES_PACKET() M3u8PcrControl {
	_init_.Initialize()
	var returns M3u8PcrControl
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M3u8PcrControl",
		"PCR_EVERY_PES_PACKET",
		&returns,
	)
	return returns
}

