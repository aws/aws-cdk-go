package interfacesawsnotifications

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsnotifications/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NotificationConfiguration.
// Experimental.
type INotificationConfigurationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a NotificationConfiguration resource.
	// Experimental.
	NotificationConfigurationRef() *NotificationConfigurationReference
}

// The jsii proxy for INotificationConfigurationRef
type jsiiProxy_INotificationConfigurationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_INotificationConfigurationRef) NotificationConfigurationRef() *NotificationConfigurationReference {
	var returns *NotificationConfigurationReference
	_jsii_.Get(
		j,
		"notificationConfigurationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INotificationConfigurationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INotificationConfigurationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

