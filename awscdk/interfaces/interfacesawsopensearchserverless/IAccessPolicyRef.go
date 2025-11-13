package interfacesawsopensearchserverless

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsopensearchserverless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AccessPolicy.
// Experimental.
type IAccessPolicyRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a AccessPolicy resource.
	// Experimental.
	AccessPolicyRef() *AccessPolicyReference
}

// The jsii proxy for IAccessPolicyRef
type jsiiProxy_IAccessPolicyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IAccessPolicyRef) AccessPolicyRef() *AccessPolicyReference {
	var returns *AccessPolicyReference
	_jsii_.Get(
		j,
		"accessPolicyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccessPolicyRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccessPolicyRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

