package awsroute53recoverycontrol

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53recoverycontrol/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Cluster.
// Experimental.
type IClusterRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Cluster resource.
	// Experimental.
	ClusterRef() *ClusterReference
}

// The jsii proxy for IClusterRef
type jsiiProxy_IClusterRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IClusterRef) ClusterRef() *ClusterReference {
	var returns *ClusterReference
	_jsii_.Get(
		j,
		"clusterRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IClusterRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IClusterRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

