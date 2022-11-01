//go:build !no_runtime_type_checking

package awsiotactions

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/awsiot"
)

func (c *jsiiProxy_CloudWatchSetAlarmStateAction) validateBindParameters(topicRule awsiot.ITopicRule) error {
	if topicRule == nil {
		return fmt.Errorf("parameter topicRule is required, but nil was provided")
	}

	return nil
}

func validateNewCloudWatchSetAlarmStateActionParameters(alarm awscloudwatch.IAlarm, props *CloudWatchSetAlarmStateActionProps) error {
	if alarm == nil {
		return fmt.Errorf("parameter alarm is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

