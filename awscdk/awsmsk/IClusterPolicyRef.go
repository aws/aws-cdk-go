package awsmsk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsmsk/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ClusterPolicy.
// Experimental.
type IClusterPolicyRef interface {
	constructs.IConstruct
	// A reference to a ClusterPolicy resource.
	// Experimental.
	ClusterPolicyRef() *ClusterPolicyReference
}

// The jsii proxy for IClusterPolicyRef
type jsiiProxy_IClusterPolicyRef struct {
	internal.Type__constructsIConstruct
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

