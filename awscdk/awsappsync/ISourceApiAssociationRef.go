package awsappsync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SourceApiAssociation.
// Experimental.
type ISourceApiAssociationRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISourceApiAssociationRef
type jsiiProxy_ISourceApiAssociationRef struct {
	internal.Type__constructsIConstruct
}

