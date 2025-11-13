package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Interface for LaunchTemplate-like objects.
type ILaunchTemplate interface {
	interfacesawsec2.ILaunchTemplateRef
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
	internal.Type__interfacesawsec2ILaunchTemplateRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_ILaunchTemplate) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
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

func (j *jsiiProxy_ILaunchTemplate) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILaunchTemplate) LaunchTemplateRef() *interfacesawsec2.LaunchTemplateReference {
	var returns *interfacesawsec2.LaunchTemplateReference
	_jsii_.Get(
		j,
		"launchTemplateRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILaunchTemplate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILaunchTemplate) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

