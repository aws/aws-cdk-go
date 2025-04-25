package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A class that provides convenient access to special version tokens for LaunchTemplate versions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   launchTemplateSpecialVersions := awscdk.Aws_ec2.NewLaunchTemplateSpecialVersions()
//
type LaunchTemplateSpecialVersions interface {
}

// The jsii proxy struct for LaunchTemplateSpecialVersions
type jsiiProxy_LaunchTemplateSpecialVersions struct {
	_ byte // padding
}

func NewLaunchTemplateSpecialVersions() LaunchTemplateSpecialVersions {
	_init_.Initialize()

	j := jsiiProxy_LaunchTemplateSpecialVersions{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.LaunchTemplateSpecialVersions",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewLaunchTemplateSpecialVersions_Override(l LaunchTemplateSpecialVersions) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.LaunchTemplateSpecialVersions",
		nil, // no parameters
		l,
	)
}

func LaunchTemplateSpecialVersions_DEFAULT_VERSION() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.LaunchTemplateSpecialVersions",
		"DEFAULT_VERSION",
		&returns,
	)
	return returns
}

func LaunchTemplateSpecialVersions_LATEST_VERSION() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.LaunchTemplateSpecialVersions",
		"LATEST_VERSION",
		&returns,
	)
	return returns
}

