package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// AAC VBR quality level.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   aacVbrQuality := medialive_alpha.AacVbrQuality_HIGH()
//
// Experimental.
type AacVbrQuality interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for AacVbrQuality
type jsiiProxy_AacVbrQuality struct {
	_ byte // padding
}

func (j *jsiiProxy_AacVbrQuality) Value() *string {
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
func AacVbrQuality_Of(value *string) AacVbrQuality {
	_init_.Initialize()

	if err := validateAacVbrQuality_OfParameters(value); err != nil {
		panic(err)
	}
	var returns AacVbrQuality

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.AacVbrQuality",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func AacVbrQuality_HIGH() AacVbrQuality {
	_init_.Initialize()
	var returns AacVbrQuality
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AacVbrQuality",
		"HIGH",
		&returns,
	)
	return returns
}

func AacVbrQuality_LOW() AacVbrQuality {
	_init_.Initialize()
	var returns AacVbrQuality
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AacVbrQuality",
		"LOW",
		&returns,
	)
	return returns
}

func AacVbrQuality_MEDIUM_HIGH() AacVbrQuality {
	_init_.Initialize()
	var returns AacVbrQuality
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AacVbrQuality",
		"MEDIUM_HIGH",
		&returns,
	)
	return returns
}

func AacVbrQuality_MEDIUM_LOW() AacVbrQuality {
	_init_.Initialize()
	var returns AacVbrQuality
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AacVbrQuality",
		"MEDIUM_LOW",
		&returns,
	)
	return returns
}

