package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Timed-metadata passthrough behavior for an M3U8 container.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   m3u8TimedMetadataBehavior := medialive_alpha.M3u8TimedMetadataBehavior_Of(jsii.String("value"))
//
// Experimental.
type M3u8TimedMetadataBehavior interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for M3u8TimedMetadataBehavior
type jsiiProxy_M3u8TimedMetadataBehavior struct {
	_ byte // padding
}

func (j *jsiiProxy_M3u8TimedMetadataBehavior) Value() *string {
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
func M3u8TimedMetadataBehavior_Of(value *string) M3u8TimedMetadataBehavior {
	_init_.Initialize()

	if err := validateM3u8TimedMetadataBehavior_OfParameters(value); err != nil {
		panic(err)
	}
	var returns M3u8TimedMetadataBehavior

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.M3u8TimedMetadataBehavior",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func M3u8TimedMetadataBehavior_NO_PASSTHROUGH() M3u8TimedMetadataBehavior {
	_init_.Initialize()
	var returns M3u8TimedMetadataBehavior
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M3u8TimedMetadataBehavior",
		"NO_PASSTHROUGH",
		&returns,
	)
	return returns
}

func M3u8TimedMetadataBehavior_PASSTHROUGH() M3u8TimedMetadataBehavior {
	_init_.Initialize()
	var returns M3u8TimedMetadataBehavior
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M3u8TimedMetadataBehavior",
		"PASSTHROUGH",
		&returns,
	)
	return returns
}

