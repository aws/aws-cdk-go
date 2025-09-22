package awsnotifications

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsnotifications/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ManagedNotificationAccountContactAssociation.
// Experimental.
type IManagedNotificationAccountContactAssociationRef interface {
	constructs.IConstruct
	// A reference to a ManagedNotificationAccountContactAssociation resource.
	// Experimental.
	ManagedNotificationAccountContactAssociationRef() *ManagedNotificationAccountContactAssociationReference
}

// The jsii proxy for IManagedNotificationAccountContactAssociationRef
type jsiiProxy_IManagedNotificationAccountContactAssociationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IManagedNotificationAccountContactAssociationRef) ManagedNotificationAccountContactAssociationRef() *ManagedNotificationAccountContactAssociationReference {
	var returns *ManagedNotificationAccountContactAssociationReference
	_jsii_.Get(
		j,
		"managedNotificationAccountContactAssociationRef",
		&returns,
	)
	return returns
}

