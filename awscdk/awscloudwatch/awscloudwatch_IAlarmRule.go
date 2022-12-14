package awscloudwatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface for Alarm Rule.
type IAlarmRule interface {
	// serialized representation of Alarm Rule to be used when building the Composite Alarm resource.
	RenderAlarmRule() *string
}

// The jsii proxy for IAlarmRule
type jsiiProxy_IAlarmRule struct {
	_ byte // padding
}

func (i *jsiiProxy_IAlarmRule) RenderAlarmRule() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"renderAlarmRule",
		nil, // no parameters
		&returns,
	)

	return returns
}

