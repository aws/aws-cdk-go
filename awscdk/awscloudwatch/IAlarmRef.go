package awscloudwatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Alarm.
// Experimental.
type IAlarmRef interface {
	constructs.IConstruct
	// A reference to a Alarm resource.
	// Experimental.
	AlarmRef() *AlarmReference
}

// The jsii proxy for IAlarmRef
type jsiiProxy_IAlarmRef struct {
	internal.Type__constructsIConstruct
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

