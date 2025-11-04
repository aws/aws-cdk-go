package awsecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ClusterCapacityProviderAssociations.
// Experimental.
type IClusterCapacityProviderAssociationsRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ClusterCapacityProviderAssociations resource.
	// Experimental.
	ClusterCapacityProviderAssociationsRef() *ClusterCapacityProviderAssociationsReference
}

// The jsii proxy for IClusterCapacityProviderAssociationsRef
type jsiiProxy_IClusterCapacityProviderAssociationsRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IClusterCapacityProviderAssociationsRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IClusterCapacityProviderAssociationsRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

