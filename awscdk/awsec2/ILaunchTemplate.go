package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
)

// Interface for LaunchTemplate-like objects.
type ILaunchTemplate interface {
	awscdk.IResource
	// The identifier of the Launch Template.
	//
	// Exactly one of `launchTemplateId` and `launchTemplateName` will be set.
	LaunchTemplateId() *string
	// The name of the Launch Template.
	//
	// Exactly one of `launchTemplateId` and `launchTemplateName` will be set.
	LaunchTemplateName() *string
	// The version number of this launch template to use.
	VersionNumber() *string
}

// The jsii proxy for ILaunchTemplate
type jsiiProxy_ILaunchTemplate struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_ILaunchTemplate) LaunchTemplateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchTemplateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILaunchTemplate) LaunchTemplateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchTemplateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILaunchTemplate) VersionNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionNumber",
		&returns,
	)
	return returns
}

