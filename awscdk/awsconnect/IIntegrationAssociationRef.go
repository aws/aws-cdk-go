package awsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IntegrationAssociation.
// Experimental.
type IIntegrationAssociationRef interface {
	constructs.IConstruct
	// A reference to a IntegrationAssociation resource.
	// Experimental.
	IntegrationAssociationRef() *IntegrationAssociationReference
}

// The jsii proxy for IIntegrationAssociationRef
type jsiiProxy_IIntegrationAssociationRef struct {
	internal.Type__constructsIConstruct
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

