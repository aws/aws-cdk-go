package interfacesawsredshift

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsredshift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ClusterParameterGroup.
// Experimental.
type IClusterParameterGroupRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ClusterParameterGroup resource.
	// Experimental.
	ClusterParameterGroupRef() *ClusterParameterGroupReference
}

// The jsii proxy for IClusterParameterGroupRef
type jsiiProxy_IClusterParameterGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IClusterParameterGroupRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IClusterParameterGroupRef) ClusterParameterGroupRef() *ClusterParameterGroupReference {
	var returns *ClusterParameterGroupReference
	_jsii_.Get(
		j,
		"clusterParameterGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IClusterParameterGroupRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IClusterParameterGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

