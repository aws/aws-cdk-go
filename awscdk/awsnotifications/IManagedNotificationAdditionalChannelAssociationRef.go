package awsnotifications

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsnotifications/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ManagedNotificationAdditionalChannelAssociation.
// Experimental.
type IManagedNotificationAdditionalChannelAssociationRef interface {
	constructs.IConstruct
	// A reference to a ManagedNotificationAdditionalChannelAssociation resource.
	// Experimental.
	ManagedNotificationAdditionalChannelAssociationRef() *ManagedNotificationAdditionalChannelAssociationReference
}

// The jsii proxy for IManagedNotificationAdditionalChannelAssociationRef
type jsiiProxy_IManagedNotificationAdditionalChannelAssociationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IManagedNotificationAdditionalChannelAssociationRef) ManagedNotificationAdditionalChannelAssociationRef() *ManagedNotificationAdditionalChannelAssociationReference {
	var returns *ManagedNotificationAdditionalChannelAssociationReference
	_jsii_.Get(
		j,
		"managedNotificationAdditionalChannelAssociationRef",
		&returns,
	)
	return returns
}

