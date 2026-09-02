package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// CMAF Ingest Nielsen ID3 behavior.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   nielsenId3Behavior := medialive_alpha.NielsenId3Behavior_Of(jsii.String("value"))
//
// Experimental.
type NielsenId3Behavior interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for NielsenId3Behavior
type jsiiProxy_NielsenId3Behavior struct {
	_ byte // padding
}

func (j *jsiiProxy_NielsenId3Behavior) Value() *string {
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
func NielsenId3Behavior_Of(value *string) NielsenId3Behavior {
	_init_.Initialize()

	if err := validateNielsenId3Behavior_OfParameters(value); err != nil {
		panic(err)
	}
	var returns NielsenId3Behavior

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.NielsenId3Behavior",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func NielsenId3Behavior_NO_PASSTHROUGH() NielsenId3Behavior {
	_init_.Initialize()
	var returns NielsenId3Behavior
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.NielsenId3Behavior",
		"NO_PASSTHROUGH",
		&returns,
	)
	return returns
}

func NielsenId3Behavior_PASSTHROUGH() NielsenId3Behavior {
	_init_.Initialize()
	var returns NielsenId3Behavior
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.NielsenId3Behavior",
		"PASSTHROUGH",
		&returns,
	)
	return returns
}

