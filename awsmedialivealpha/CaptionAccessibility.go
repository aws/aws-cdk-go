package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Whether a caption track implements accessibility features (written descriptions of dialog, music, and sounds).
//
// Signaled in HLS and MediaPackage output groups.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   captionAccessibility := medialive_alpha.CaptionAccessibility_Of(jsii.String("value"))
//
// Experimental.
type CaptionAccessibility interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for CaptionAccessibility
type jsiiProxy_CaptionAccessibility struct {
	_ byte // padding
}

func (j *jsiiProxy_CaptionAccessibility) Value() *string {
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
func CaptionAccessibility_Of(value *string) CaptionAccessibility {
	_init_.Initialize()

	if err := validateCaptionAccessibility_OfParameters(value); err != nil {
		panic(err)
	}
	var returns CaptionAccessibility

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.CaptionAccessibility",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CaptionAccessibility_DOES_NOT_IMPLEMENT_ACCESSIBILITY_FEATURES() CaptionAccessibility {
	_init_.Initialize()
	var returns CaptionAccessibility
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.CaptionAccessibility",
		"DOES_NOT_IMPLEMENT_ACCESSIBILITY_FEATURES",
		&returns,
	)
	return returns
}

func CaptionAccessibility_IMPLEMENTS_ACCESSIBILITY_FEATURES() CaptionAccessibility {
	_init_.Initialize()
	var returns CaptionAccessibility
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.CaptionAccessibility",
		"IMPLEMENTS_ACCESSIBILITY_FEATURES",
		&returns,
	)
	return returns
}

