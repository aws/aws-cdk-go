package awsredshift

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsredshift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ClusterParameterGroup.
// Experimental.
type IClusterParameterGroupRef interface {
	constructs.IConstruct
	// A reference to a ClusterParameterGroup resource.
	// Experimental.
	ClusterParameterGroupRef() *ClusterParameterGroupReference
}

// The jsii proxy for IClusterParameterGroupRef
type jsiiProxy_IClusterParameterGroupRef struct {
	internal.Type__constructsIConstruct
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

