package awsdeadline

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdeadline/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a QueueFleetAssociation.
// Experimental.
type IQueueFleetAssociationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IQueueFleetAssociationRef
type jsiiProxy_IQueueFleetAssociationRef struct {
	internal.Type__constructsIConstruct
}

