package awsemr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsemr/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Cluster.
// Experimental.
type IClusterRef interface {
	constructs.IConstruct
	// A reference to a Cluster resource.
	// Experimental.
	ClusterRef() *ClusterReference
}

// The jsii proxy for IClusterRef
type jsiiProxy_IClusterRef struct {
	internal.Type__constructsIConstruct
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

