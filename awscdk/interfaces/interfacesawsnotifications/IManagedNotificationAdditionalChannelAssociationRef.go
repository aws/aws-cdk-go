package interfacesawsnotifications

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsnotifications/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ManagedNotificationAdditionalChannelAssociation.
// Experimental.
type IManagedNotificationAdditionalChannelAssociationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ManagedNotificationAdditionalChannelAssociation resource.
	// Experimental.
	ManagedNotificationAdditionalChannelAssociationRef() *ManagedNotificationAdditionalChannelAssociationReference
}

// The jsii proxy for IManagedNotificationAdditionalChannelAssociationRef
type jsiiProxy_IManagedNotificationAdditionalChannelAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IManagedNotificationAdditionalChannelAssociationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
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

func (j *jsiiProxy_IManagedNotificationAdditionalChannelAssociationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IManagedNotificationAdditionalChannelAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

