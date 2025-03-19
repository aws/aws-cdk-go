package awseks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseks/internal"
)

// Represents an Amazon EKS Add-On.
type IAddon interface {
	awscdk.IResource
	// ARN of the Add-On.
	AddonArn() *string
	// Name of the Add-On.
	AddonName() *string
}

// The jsii proxy for IAddon
type jsiiProxy_IAddon struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IAddon) AddonArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addonArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAddon) AddonName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addonName",
		&returns,
	)
	return returns
}

