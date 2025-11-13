package interfacesawscloudtrail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscloudtrail/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResourcePolicy.
// Experimental.
type IResourcePolicyRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ResourcePolicy resource.
	// Experimental.
	ResourcePolicyRef() *ResourcePolicyReference
}

// The jsii proxy for IResourcePolicyRef
type jsiiProxy_IResourcePolicyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IResourcePolicyRef) ResourcePolicyRef() *ResourcePolicyReference {
	var returns *ResourcePolicyReference
	_jsii_.Get(
		j,
		"resourcePolicyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResourcePolicyRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResourcePolicyRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

