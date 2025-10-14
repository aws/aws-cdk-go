package awsiotevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotevents/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AlarmModel.
// Experimental.
type IAlarmModelRef interface {
	constructs.IConstruct
	// A reference to a AlarmModel resource.
	// Experimental.
	AlarmModelRef() *AlarmModelReference
}

// The jsii proxy for IAlarmModelRef
type jsiiProxy_IAlarmModelRef struct {
	internal.Type__constructsIConstruct
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

