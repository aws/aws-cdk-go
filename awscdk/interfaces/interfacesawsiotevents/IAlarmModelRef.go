package interfacesawsiotevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiotevents/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AlarmModel.
// Experimental.
type IAlarmModelRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a AlarmModel resource.
	// Experimental.
	AlarmModelRef() *AlarmModelReference
}

// The jsii proxy for IAlarmModelRef
type jsiiProxy_IAlarmModelRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IAlarmModelRef) AlarmModelRef() *AlarmModelReference {
	var returns *AlarmModelReference
	_jsii_.Get(
		j,
		"alarmModelRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlarmModelRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlarmModelRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

