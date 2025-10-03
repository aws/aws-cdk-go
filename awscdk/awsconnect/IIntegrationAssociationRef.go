package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IntegrationAssociation.
// Experimental.
type IIntegrationAssociationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IIntegrationAssociationRef
type jsiiProxy_IIntegrationAssociationRef struct {
	internal.Type__constructsIConstruct
}

