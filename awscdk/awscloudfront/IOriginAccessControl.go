package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscloudfront"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a CloudFront Origin Access Control.
type IOriginAccessControl interface {
	interfacesawscloudfront.IOriginAccessControlRef
	awscdk.IResource
	// The unique identifier of the origin access control.
	OriginAccessControlId() *string
}

// The jsii proxy for IOriginAccessControl
type jsiiProxy_IOriginAccessControl struct {
	internal.Type__interfacesawscloudfrontIOriginAccessControlRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IOriginAccessControl) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
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

func (j *jsiiProxy_IOriginAccessControl) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOriginAccessControl) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOriginAccessControl) OriginAccessControlRef() *interfacesawscloudfront.OriginAccessControlReference {
	var returns *interfacesawscloudfront.OriginAccessControlReference
	_jsii_.Get(
		j,
		"originAccessControlRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOriginAccessControl) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

