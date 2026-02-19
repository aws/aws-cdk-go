package interfacesawsmsk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmsk/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ClusterPolicy.
// Experimental.
type IClusterPolicyRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ClusterPolicy resource.
	// Experimental.
	ClusterPolicyRef() *ClusterPolicyReference
}

// The jsii proxy for IClusterPolicyRef
type jsiiProxy_IClusterPolicyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IClusterPolicyRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IClusterPolicyRef) ClusterPolicyRef() *ClusterPolicyReference {
	var returns *ClusterPolicyReference
	_jsii_.Get(
		j,
		"clusterPolicyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IClusterPolicyRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IClusterPolicyRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

