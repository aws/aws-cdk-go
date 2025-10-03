package awsdeadline

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdeadline/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a QueueLimitAssociation.
// Experimental.
type IQueueLimitAssociationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IQueueLimitAssociationRef
type jsiiProxy_IQueueLimitAssociationRef struct {
	internal.Type__constructsIConstruct
}

