package interfacesawsusernotifications

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsusernotifications/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ManagedNotificationConfiguration.
// Experimental.
type IManagedNotificationConfigurationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ManagedNotificationConfiguration resource.
	// Experimental.
	ManagedNotificationConfigurationRef() *ManagedNotificationConfigurationReference
}

// The jsii proxy for IManagedNotificationConfigurationRef
type jsiiProxy_IManagedNotificationConfigurationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IManagedNotificationConfigurationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IManagedNotificationConfigurationRef) ManagedNotificationConfigurationRef() *ManagedNotificationConfigurationReference {
	var returns *ManagedNotificationConfigurationReference
	_jsii_.Get(
		j,
		"managedNotificationConfigurationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IManagedNotificationConfigurationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IManagedNotificationConfigurationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

