package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ServiceLinkedRole.
// Experimental.
type IServiceLinkedRoleRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ServiceLinkedRole resource.
	// Experimental.
	ServiceLinkedRoleRef() *ServiceLinkedRoleReference
}

// The jsii proxy for IServiceLinkedRoleRef
type jsiiProxy_IServiceLinkedRoleRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IServiceLinkedRoleRef) ServiceLinkedRoleRef() *ServiceLinkedRoleReference {
	var returns *ServiceLinkedRoleReference
	_jsii_.Get(
		j,
		"serviceLinkedRoleRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServiceLinkedRoleRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServiceLinkedRoleRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

