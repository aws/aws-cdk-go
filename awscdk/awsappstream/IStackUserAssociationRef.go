package awsappstream

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappstream/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StackUserAssociation.
// Experimental.
type IStackUserAssociationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IStackUserAssociationRef
type jsiiProxy_IStackUserAssociationRef struct {
	internal.Type__constructsIConstruct
}

