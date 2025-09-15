package awsecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ClusterCapacityProviderAssociations.
// Experimental.
type IClusterCapacityProviderAssociationsRef interface {
	constructs.IConstruct
	// A reference to a ClusterCapacityProviderAssociations resource.
	// Experimental.
	ClusterCapacityProviderAssociationsRef() *ClusterCapacityProviderAssociationsReference
}

// The jsii proxy for IClusterCapacityProviderAssociationsRef
type jsiiProxy_IClusterCapacityProviderAssociationsRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IClusterCapacityProviderAssociationsRef) ClusterCapacityProviderAssociationsRef() *ClusterCapacityProviderAssociationsReference {
	var returns *ClusterCapacityProviderAssociationsReference
	_jsii_.Get(
		j,
		"clusterCapacityProviderAssociationsRef",
		&returns,
	)
	return returns
}

