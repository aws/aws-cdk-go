package awsredshift

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsredshift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ClusterParameterGroup.
// Experimental.
type IClusterParameterGroupRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ClusterParameterGroup resource.
	// Experimental.
	ClusterParameterGroupRef() *ClusterParameterGroupReference
}

// The jsii proxy for IClusterParameterGroupRef
type jsiiProxy_IClusterParameterGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IClusterParameterGroupRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
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

