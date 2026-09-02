package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The stream type used for audio elementary streams.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   m2tsAudioStreamType := medialive_alpha.M2tsAudioStreamType_Of(jsii.String("value"))
//
// Experimental.
type M2tsAudioStreamType interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for M2tsAudioStreamType
type jsiiProxy_M2tsAudioStreamType struct {
	_ byte // padding
}

func (j *jsiiProxy_M2tsAudioStreamType) Value() *string {
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
func M2tsAudioStreamType_Of(value *string) M2tsAudioStreamType {
	_init_.Initialize()

	if err := validateM2tsAudioStreamType_OfParameters(value); err != nil {
		panic(err)
	}
	var returns M2tsAudioStreamType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.M2tsAudioStreamType",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func M2tsAudioStreamType_ATSC() M2tsAudioStreamType {
	_init_.Initialize()
	var returns M2tsAudioStreamType
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M2tsAudioStreamType",
		"ATSC",
		&returns,
	)
	return returns
}

func M2tsAudioStreamType_DVB() M2tsAudioStreamType {
	_init_.Initialize()
	var returns M2tsAudioStreamType
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M2tsAudioStreamType",
		"DVB",
		&returns,
	)
	return returns
}

