package interfacesawsvpclattice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsvpclattice/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ServiceNetworkVpcAssociation.
// Experimental.
type IServiceNetworkVpcAssociationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ServiceNetworkVpcAssociation resource.
	// Experimental.
	ServiceNetworkVpcAssociationRef() *ServiceNetworkVpcAssociationReference
}

// The jsii proxy for IServiceNetworkVpcAssociationRef
type jsiiProxy_IServiceNetworkVpcAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IServiceNetworkVpcAssociationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IServiceNetworkVpcAssociationRef) ServiceNetworkVpcAssociationRef() *ServiceNetworkVpcAssociationReference {
	var returns *ServiceNetworkVpcAssociationReference
	_jsii_.Get(
		j,
		"serviceNetworkVpcAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServiceNetworkVpcAssociationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServiceNetworkVpcAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

