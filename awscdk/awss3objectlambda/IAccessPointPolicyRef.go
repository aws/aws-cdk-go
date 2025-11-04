package awss3objectlambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3objectlambda/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AccessPointPolicy.
// Experimental.
type IAccessPointPolicyRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a AccessPointPolicy resource.
	// Experimental.
	AccessPointPolicyRef() *AccessPointPolicyReference
}

// The jsii proxy for IAccessPointPolicyRef
type jsiiProxy_IAccessPointPolicyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IAccessPointPolicyRef) AccessPointPolicyRef() *AccessPointPolicyReference {
	var returns *AccessPointPolicyReference
	_jsii_.Get(
		j,
		"accessPointPolicyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccessPointPolicyRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccessPointPolicyRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

