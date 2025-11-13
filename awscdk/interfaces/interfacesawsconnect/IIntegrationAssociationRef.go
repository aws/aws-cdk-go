package interfacesawsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IntegrationAssociation.
// Experimental.
type IIntegrationAssociationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a IntegrationAssociation resource.
	// Experimental.
	IntegrationAssociationRef() *IntegrationAssociationReference
}

// The jsii proxy for IIntegrationAssociationRef
type jsiiProxy_IIntegrationAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IIntegrationAssociationRef) IntegrationAssociationRef() *IntegrationAssociationReference {
	var returns *IntegrationAssociationReference
	_jsii_.Get(
		j,
		"integrationAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIntegrationAssociationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIntegrationAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

