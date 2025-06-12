package awscdkapprunneralpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkapprunneralpha/v2/internal"
)

// Represents the App Runner VPC Ingress Connection.
// Experimental.
type IVpcIngressConnection interface {
	awscdk.IResource
	// The Amazon Resource Name (ARN) of the VPC Ingress Connection.
	// Experimental.
	VpcIngressConnectionArn() *string
	// The name of the VPC Ingress Connection.
	// Experimental.
	VpcIngressConnectionName() *string
}

// The jsii proxy for IVpcIngressConnection
type jsiiProxy_IVpcIngressConnection struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IVpcIngressConnection) VpcIngressConnectionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIngressConnectionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpcIngressConnection) VpcIngressConnectionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIngressConnectionName",
		&returns,
	)
	return returns
}

