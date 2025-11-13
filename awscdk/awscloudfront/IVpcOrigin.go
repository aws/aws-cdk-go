package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscloudfront"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a VPC origin.
type IVpcOrigin interface {
	awscdk.IResource
	interfacesawscloudfront.IVpcOriginRef
	// The domain name of the CloudFront VPC origin endpoint configuration.
	DomainName() *string
	// The VPC origin ARN.
	VpcOriginArn() *string
	// The VPC origin ID.
	VpcOriginId() *string
}

// The jsii proxy for IVpcOrigin
type jsiiProxy_IVpcOrigin struct {
	internal.Type__awscdkIResource
	internal.Type__interfacesawscloudfrontIVpcOriginRef
}

func (i *jsiiProxy_IVpcOrigin) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IVpcOrigin) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpcOrigin) VpcOriginArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcOriginArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpcOrigin) VpcOriginId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcOriginId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpcOrigin) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpcOrigin) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpcOrigin) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpcOrigin) VpcOriginRef() *interfacesawscloudfront.VpcOriginReference {
	var returns *interfacesawscloudfront.VpcOriginReference
	_jsii_.Get(
		j,
		"vpcOriginRef",
		&returns,
	)
	return returns
}

