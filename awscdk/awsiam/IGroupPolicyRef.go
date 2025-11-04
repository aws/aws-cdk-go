package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GroupPolicy.
// Experimental.
type IGroupPolicyRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a GroupPolicy resource.
	// Experimental.
	GroupPolicyRef() *GroupPolicyReference
}

// The jsii proxy for IGroupPolicyRef
type jsiiProxy_IGroupPolicyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IGroupPolicyRef) GroupPolicyRef() *GroupPolicyReference {
	var returns *GroupPolicyReference
	_jsii_.Get(
		j,
		"groupPolicyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGroupPolicyRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGroupPolicyRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

