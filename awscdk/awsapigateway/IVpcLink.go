package awsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway/internal"
)

// Represents an API Gateway VpcLink.
type IVpcLink interface {
	awscdk.IResource
	// Physical ID of the VpcLink resource.
	VpcLinkId() *string
}

// The jsii proxy for IVpcLink
type jsiiProxy_IVpcLink struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IVpcLink) VpcLinkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcLinkId",
		&returns,
	)
	return returns
}

