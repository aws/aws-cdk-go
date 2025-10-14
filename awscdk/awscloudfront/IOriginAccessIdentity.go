package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// Interface for CloudFront OriginAccessIdentity.
type IOriginAccessIdentity interface {
	ICloudFrontOriginAccessIdentityRef
	awsiam.IGrantable
	awscdk.IResource
	// The Origin Access Identity Id (physical id) This was called originAccessIdentityName before.
	OriginAccessIdentityId() *string
	// The Origin Access Identity Id (physical id) It is misnamed and superseded by the correctly named originAccessIdentityId.
	// Deprecated: use originAccessIdentityId instead.
	OriginAccessIdentityName() *string
}

// The jsii proxy for IOriginAccessIdentity
type jsiiProxy_IOriginAccessIdentity struct {
	jsiiProxy_ICloudFrontOriginAccessIdentityRef
	internal.Type__awsiamIGrantable
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IOriginAccessIdentity) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IOriginAccessIdentity) OriginAccessIdentityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originAccessIdentityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOriginAccessIdentity) OriginAccessIdentityName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originAccessIdentityName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOriginAccessIdentity) CloudFrontOriginAccessIdentityRef() *CloudFrontOriginAccessIdentityReference {
	var returns *CloudFrontOriginAccessIdentityReference
	_jsii_.Get(
		j,
		"cloudFrontOriginAccessIdentityRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOriginAccessIdentity) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOriginAccessIdentity) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOriginAccessIdentity) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOriginAccessIdentity) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

