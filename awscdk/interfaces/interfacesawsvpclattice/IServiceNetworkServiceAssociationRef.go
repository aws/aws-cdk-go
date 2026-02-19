package interfacesawsvpclattice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsvpclattice/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ServiceNetworkServiceAssociation.
// Experimental.
type IServiceNetworkServiceAssociationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ServiceNetworkServiceAssociation resource.
	// Experimental.
	ServiceNetworkServiceAssociationRef() *ServiceNetworkServiceAssociationReference
}

// The jsii proxy for IServiceNetworkServiceAssociationRef
type jsiiProxy_IServiceNetworkServiceAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IServiceNetworkServiceAssociationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IServiceNetworkServiceAssociationRef) ServiceNetworkServiceAssociationRef() *ServiceNetworkServiceAssociationReference {
	var returns *ServiceNetworkServiceAssociationReference
	_jsii_.Get(
		j,
		"serviceNetworkServiceAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServiceNetworkServiceAssociationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
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

