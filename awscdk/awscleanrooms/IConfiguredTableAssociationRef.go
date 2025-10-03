package awscleanrooms

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscleanrooms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConfiguredTableAssociation.
// Experimental.
type IConfiguredTableAssociationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IConfiguredTableAssociationRef
type jsiiProxy_IConfiguredTableAssociationRef struct {
	internal.Type__constructsIConstruct
}

