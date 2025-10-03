package awsappstream

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappstream/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StackFleetAssociation.
// Experimental.
type IStackFleetAssociationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IStackFleetAssociationRef
type jsiiProxy_IStackFleetAssociationRef struct {
	internal.Type__constructsIConstruct
}

