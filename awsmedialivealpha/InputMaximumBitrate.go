package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The maximum input bitrate for the input specification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   inputMaximumBitrate := medialive_alpha.InputMaximumBitrate_MAX_10_MBPS()
//
// Experimental.
type InputMaximumBitrate interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for InputMaximumBitrate
type jsiiProxy_InputMaximumBitrate struct {
	_ byte // padding
}

func (j *jsiiProxy_InputMaximumBitrate) Value() *string {
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
func InputMaximumBitrate_Of(value *string) InputMaximumBitrate {
	_init_.Initialize()

	if err := validateInputMaximumBitrate_OfParameters(value); err != nil {
		panic(err)
	}
	var returns InputMaximumBitrate

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.InputMaximumBitrate",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func InputMaximumBitrate_MAX_10_MBPS() InputMaximumBitrate {
	_init_.Initialize()
	var returns InputMaximumBitrate
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.InputMaximumBitrate",
		"MAX_10_MBPS",
		&returns,
	)
	return returns
}

func InputMaximumBitrate_MAX_20_MBPS() InputMaximumBitrate {
	_init_.Initialize()
	var returns InputMaximumBitrate
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.InputMaximumBitrate",
		"MAX_20_MBPS",
		&returns,
	)
	return returns
}

func InputMaximumBitrate_MAX_50_MBPS() InputMaximumBitrate {
	_init_.Initialize()
	var returns InputMaximumBitrate
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.InputMaximumBitrate",
		"MAX_50_MBPS",
		&returns,
	)
	return returns
}

