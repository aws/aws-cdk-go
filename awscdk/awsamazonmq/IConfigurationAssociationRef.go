package awsamazonmq

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsamazonmq/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConfigurationAssociation.
// Experimental.
type IConfigurationAssociationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IConfigurationAssociationRef
type jsiiProxy_IConfigurationAssociationRef struct {
	internal.Type__constructsIConstruct
}

