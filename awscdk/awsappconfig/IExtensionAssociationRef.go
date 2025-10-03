package awsappconfig

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ExtensionAssociation.
// Experimental.
type IExtensionAssociationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IExtensionAssociationRef
type jsiiProxy_IExtensionAssociationRef struct {
	internal.Type__constructsIConstruct
}

