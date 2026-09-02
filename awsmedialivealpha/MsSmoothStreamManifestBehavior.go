package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// MS Smooth stream manifest behavior.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   msSmoothStreamManifestBehavior := medialive_alpha.MsSmoothStreamManifestBehavior_Of(jsii.String("value"))
//
// Experimental.
type MsSmoothStreamManifestBehavior interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for MsSmoothStreamManifestBehavior
type jsiiProxy_MsSmoothStreamManifestBehavior struct {
	_ byte // padding
}

func (j *jsiiProxy_MsSmoothStreamManifestBehavior) Value() *string {
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
func MsSmoothStreamManifestBehavior_Of(value *string) MsSmoothStreamManifestBehavior {
	_init_.Initialize()

	if err := validateMsSmoothStreamManifestBehavior_OfParameters(value); err != nil {
		panic(err)
	}
	var returns MsSmoothStreamManifestBehavior

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.MsSmoothStreamManifestBehavior",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func MsSmoothStreamManifestBehavior_DO_NOT_SEND() MsSmoothStreamManifestBehavior {
	_init_.Initialize()
	var returns MsSmoothStreamManifestBehavior
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.MsSmoothStreamManifestBehavior",
		"DO_NOT_SEND",
		&returns,
	)
	return returns
}

func MsSmoothStreamManifestBehavior_SEND() MsSmoothStreamManifestBehavior {
	_init_.Initialize()
	var returns MsSmoothStreamManifestBehavior
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.MsSmoothStreamManifestBehavior",
		"SEND",
		&returns,
	)
	return returns
}

