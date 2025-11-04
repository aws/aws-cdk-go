package awsvpclattice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsvpclattice/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ServiceNetworkServiceAssociation.
// Experimental.
type IServiceNetworkServiceAssociationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ServiceNetworkServiceAssociation resource.
	// Experimental.
	ServiceNetworkServiceAssociationRef() *ServiceNetworkServiceAssociationReference
}

// The jsii proxy for IServiceNetworkServiceAssociationRef
type jsiiProxy_IServiceNetworkServiceAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IServiceNetworkServiceAssociationRef) ServiceNetworkServiceAssociationRef() *ServiceNetworkServiceAssociationReference {
	var returns *ServiceNetworkServiceAssociationReference
	_jsii_.Get(
		j,
		"serviceNetworkServiceAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServiceNetworkServiceAssociationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServiceNetworkServiceAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

