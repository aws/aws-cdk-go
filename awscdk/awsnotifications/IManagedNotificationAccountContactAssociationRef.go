package awsnotifications

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsnotifications/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ManagedNotificationAccountContactAssociation.
// Experimental.
type IManagedNotificationAccountContactAssociationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IManagedNotificationAccountContactAssociationRef
type jsiiProxy_IManagedNotificationAccountContactAssociationRef struct {
	internal.Type__constructsIConstruct
}

