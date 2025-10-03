package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
)

// Represents a VPC origin.
type IVpcOrigin interface {
	awscdk.IResource
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

