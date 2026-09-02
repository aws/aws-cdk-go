package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Controls placement of audio Encoder Boundary Point (EBP) markers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   m2tsEbpAudioInterval := medialive_alpha.M2tsEbpAudioInterval_Of(jsii.String("value"))
//
// Experimental.
type M2tsEbpAudioInterval interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for M2tsEbpAudioInterval
type jsiiProxy_M2tsEbpAudioInterval struct {
	_ byte // padding
}

func (j *jsiiProxy_M2tsEbpAudioInterval) Value() *string {
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
func M2tsEbpAudioInterval_Of(value *string) M2tsEbpAudioInterval {
	_init_.Initialize()

	if err := validateM2tsEbpAudioInterval_OfParameters(value); err != nil {
		panic(err)
	}
	var returns M2tsEbpAudioInterval

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.M2tsEbpAudioInterval",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func M2tsEbpAudioInterval_VIDEO_AND_FIXED_INTERVALS() M2tsEbpAudioInterval {
	_init_.Initialize()
	var returns M2tsEbpAudioInterval
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M2tsEbpAudioInterval",
		"VIDEO_AND_FIXED_INTERVALS",
		&returns,
	)
	return returns
}

func M2tsEbpAudioInterval_VIDEO_INTERVAL() M2tsEbpAudioInterval {
	_init_.Initialize()
	var returns M2tsEbpAudioInterval
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M2tsEbpAudioInterval",
		"VIDEO_INTERVAL",
		&returns,
	)
	return returns
}

