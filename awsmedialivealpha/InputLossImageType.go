package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The image MediaLive substitutes into the output on input loss.
//
// Example:
//   var slateBucket IBucket
//
//
//   inputLoss := &InputLossBehavior{
//   	BlackFrame: awscdk.Duration_Seconds(jsii.Number(1)),
//   	RepeatFrame: awscdk.Duration_*Seconds(jsii.Number(5)),
//   	ImageType: medialive.InputLossImageType_SLATE(),
//   	ImageSlate: medialive.FileLocation_FromBucket(slateBucket, jsii.String("slates/offline.png")),
//   }
//
// Experimental.
type InputLossImageType interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for InputLossImageType
type jsiiProxy_InputLossImageType struct {
	_ byte // padding
}

func (j *jsiiProxy_InputLossImageType) Value() *string {
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
func InputLossImageType_Of(value *string) InputLossImageType {
	_init_.Initialize()

	if err := validateInputLossImageType_OfParameters(value); err != nil {
		panic(err)
	}
	var returns InputLossImageType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.InputLossImageType",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func InputLossImageType_COLOR() InputLossImageType {
	_init_.Initialize()
	var returns InputLossImageType
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.InputLossImageType",
		"COLOR",
		&returns,
	)
	return returns
}

func InputLossImageType_SLATE() InputLossImageType {
	_init_.Initialize()
	var returns InputLossImageType
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.InputLossImageType",
		"SLATE",
		&returns,
	)
	return returns
}

