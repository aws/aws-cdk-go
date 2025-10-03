package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CloudFrontOriginAccessIdentity.
// Experimental.
type ICloudFrontOriginAccessIdentityRef interface {
	constructs.IConstruct
}

// The jsii proxy for ICloudFrontOriginAccessIdentityRef
type jsiiProxy_ICloudFrontOriginAccessIdentityRef struct {
	internal.Type__constructsIConstruct
}

