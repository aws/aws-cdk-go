package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Nielsen ID3 passthrough behavior for an M3U8 container.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   m3u8NielsenId3Behavior := medialive_alpha.M3u8NielsenId3Behavior_Of(jsii.String("value"))
//
// Experimental.
type M3u8NielsenId3Behavior interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for M3u8NielsenId3Behavior
type jsiiProxy_M3u8NielsenId3Behavior struct {
	_ byte // padding
}

func (j *jsiiProxy_M3u8NielsenId3Behavior) Value() *string {
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
func M3u8NielsenId3Behavior_Of(value *string) M3u8NielsenId3Behavior {
	_init_.Initialize()

	if err := validateM3u8NielsenId3Behavior_OfParameters(value); err != nil {
		panic(err)
	}
	var returns M3u8NielsenId3Behavior

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.M3u8NielsenId3Behavior",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func M3u8NielsenId3Behavior_NO_PASSTHROUGH() M3u8NielsenId3Behavior {
	_init_.Initialize()
	var returns M3u8NielsenId3Behavior
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M3u8NielsenId3Behavior",
		"NO_PASSTHROUGH",
		&returns,
	)
	return returns
}

func M3u8NielsenId3Behavior_PASSTHROUGH() M3u8NielsenId3Behavior {
	_init_.Initialize()
	var returns M3u8NielsenId3Behavior
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M3u8NielsenId3Behavior",
		"PASSTHROUGH",
		&returns,
	)
	return returns
}

