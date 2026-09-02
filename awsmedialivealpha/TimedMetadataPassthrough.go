package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// CMAF Ingest timed metadata passthrough.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   timedMetadataPassthrough := medialive_alpha.TimedMetadataPassthrough_Of(jsii.String("value"))
//
// Experimental.
type TimedMetadataPassthrough interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for TimedMetadataPassthrough
type jsiiProxy_TimedMetadataPassthrough struct {
	_ byte // padding
}

func (j *jsiiProxy_TimedMetadataPassthrough) Value() *string {
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
func TimedMetadataPassthrough_Of(value *string) TimedMetadataPassthrough {
	_init_.Initialize()

	if err := validateTimedMetadataPassthrough_OfParameters(value); err != nil {
		panic(err)
	}
	var returns TimedMetadataPassthrough

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.TimedMetadataPassthrough",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func TimedMetadataPassthrough_DISABLED() TimedMetadataPassthrough {
	_init_.Initialize()
	var returns TimedMetadataPassthrough
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.TimedMetadataPassthrough",
		"DISABLED",
		&returns,
	)
	return returns
}

func TimedMetadataPassthrough_ENABLED() TimedMetadataPassthrough {
	_init_.Initialize()
	var returns TimedMetadataPassthrough
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.TimedMetadataPassthrough",
		"ENABLED",
		&returns,
	)
	return returns
}

