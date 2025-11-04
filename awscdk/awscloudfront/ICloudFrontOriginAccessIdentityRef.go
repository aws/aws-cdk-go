package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CloudFrontOriginAccessIdentity.
// Experimental.
type ICloudFrontOriginAccessIdentityRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a CloudFrontOriginAccessIdentity resource.
	// Experimental.
	CloudFrontOriginAccessIdentityRef() *CloudFrontOriginAccessIdentityReference
}

// The jsii proxy for ICloudFrontOriginAccessIdentityRef
type jsiiProxy_ICloudFrontOriginAccessIdentityRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ICloudFrontOriginAccessIdentityRef) CloudFrontOriginAccessIdentityRef() *CloudFrontOriginAccessIdentityReference {
	var returns *CloudFrontOriginAccessIdentityReference
	_jsii_.Get(
		j,
		"cloudFrontOriginAccessIdentityRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICloudFrontOriginAccessIdentityRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICloudFrontOriginAccessIdentityRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

