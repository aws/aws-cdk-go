package interfacesawscloudwatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscloudwatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Alarm.
// Experimental.
type IAlarmRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Alarm resource.
	// Experimental.
	AlarmRef() *AlarmReference
}

// The jsii proxy for IAlarmRef
type jsiiProxy_IAlarmRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IAlarmRef) AlarmRef() *AlarmReference {
	var returns *AlarmReference
	_jsii_.Get(
		j,
		"alarmRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlarmRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlarmRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

