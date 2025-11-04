package awsmsk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmsk/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ClusterPolicy.
// Experimental.
type IClusterPolicyRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ClusterPolicy resource.
	// Experimental.
	ClusterPolicyRef() *ClusterPolicyReference
}

// The jsii proxy for IClusterPolicyRef
type jsiiProxy_IClusterPolicyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IClusterPolicyRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
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

