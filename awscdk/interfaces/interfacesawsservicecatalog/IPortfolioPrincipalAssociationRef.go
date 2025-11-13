package interfacesawsservicecatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsservicecatalog/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PortfolioPrincipalAssociation.
// Experimental.
type IPortfolioPrincipalAssociationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a PortfolioPrincipalAssociation resource.
	// Experimental.
	PortfolioPrincipalAssociationRef() *PortfolioPrincipalAssociationReference
}

// The jsii proxy for IPortfolioPrincipalAssociationRef
type jsiiProxy_IPortfolioPrincipalAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IPortfolioPrincipalAssociationRef) PortfolioPrincipalAssociationRef() *PortfolioPrincipalAssociationReference {
	var returns *PortfolioPrincipalAssociationReference
	_jsii_.Get(
		j,
		"portfolioPrincipalAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPortfolioPrincipalAssociationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPortfolioPrincipalAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

