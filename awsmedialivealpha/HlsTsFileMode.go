package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// HLS TS file mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   hlsTsFileMode := medialive_alpha.HlsTsFileMode_Of(jsii.String("value"))
//
// Experimental.
type HlsTsFileMode interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for HlsTsFileMode
type jsiiProxy_HlsTsFileMode struct {
	_ byte // padding
}

func (j *jsiiProxy_HlsTsFileMode) Value() *string {
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
func HlsTsFileMode_Of(value *string) HlsTsFileMode {
	_init_.Initialize()

	if err := validateHlsTsFileMode_OfParameters(value); err != nil {
		panic(err)
	}
	var returns HlsTsFileMode

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.HlsTsFileMode",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func HlsTsFileMode_SEGMENTED_FILES() HlsTsFileMode {
	_init_.Initialize()
	var returns HlsTsFileMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsTsFileMode",
		"SEGMENTED_FILES",
		&returns,
	)
	return returns
}

func HlsTsFileMode_SINGLE_FILE() HlsTsFileMode {
	_init_.Initialize()
	var returns HlsTsFileMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsTsFileMode",
		"SINGLE_FILE",
		&returns,
	)
	return returns
}

