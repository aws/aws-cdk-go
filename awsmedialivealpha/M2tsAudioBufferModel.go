package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The buffer model used for Dolby Digital audio.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   m2tsAudioBufferModel := medialive_alpha.M2tsAudioBufferModel_Of(jsii.String("value"))
//
// Experimental.
type M2tsAudioBufferModel interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for M2tsAudioBufferModel
type jsiiProxy_M2tsAudioBufferModel struct {
	_ byte // padding
}

func (j *jsiiProxy_M2tsAudioBufferModel) Value() *string {
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
func M2tsAudioBufferModel_Of(value *string) M2tsAudioBufferModel {
	_init_.Initialize()

	if err := validateM2tsAudioBufferModel_OfParameters(value); err != nil {
		panic(err)
	}
	var returns M2tsAudioBufferModel

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.M2tsAudioBufferModel",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func M2tsAudioBufferModel_ATSC() M2tsAudioBufferModel {
	_init_.Initialize()
	var returns M2tsAudioBufferModel
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M2tsAudioBufferModel",
		"ATSC",
		&returns,
	)
	return returns
}

func M2tsAudioBufferModel_DVB() M2tsAudioBufferModel {
	_init_.Initialize()
	var returns M2tsAudioBufferModel
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M2tsAudioBufferModel",
		"DVB",
		&returns,
	)
	return returns
}

