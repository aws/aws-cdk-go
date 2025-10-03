package awsappstream

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappstream/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApplicationFleetAssociation.
// Experimental.
type IApplicationFleetAssociationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IApplicationFleetAssociationRef
type jsiiProxy_IApplicationFleetAssociationRef struct {
	internal.Type__constructsIConstruct
}

