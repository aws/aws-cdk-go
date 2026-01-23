package awscloudwatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Interface for objects that can be the targets of CloudWatch alarm actions.
type IAlarmAction interface {
	// Return the properties required to send alarm actions to this CloudWatch alarm.
	Bind(scope constructs.Construct, alarm IAlarm) *AlarmActionConfig
}

// The jsii proxy for IAlarmAction
type jsiiProxy_IAlarmAction struct {
	_ byte // padding
}

func (i *jsiiProxy_IAlarmAction) Bind(scope constructs.Construct, alarm IAlarm) *AlarmActionConfig {
	if err := i.validateBindParameters(scope, alarm); err != nil {
		panic(err)
	}
	var returns *AlarmActionConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope, alarm},
		&returns,
	)

	return returns
}

