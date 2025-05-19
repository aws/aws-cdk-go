package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
)

// Represents a CloudFront Origin Access Control.
type IOriginAccessControl interface {
	awscdk.IResource
	// The unique identifier of the origin access control.
	OriginAccessControlId() *string
}

// The jsii proxy for IOriginAccessControl
type jsiiProxy_IOriginAccessControl struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IOriginAccessControl) OriginAccessControlId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originAccessControlId",
		&returns,
	)
	return returns
}

