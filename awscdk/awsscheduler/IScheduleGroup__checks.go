//go:build !no_runtime_type_checking

package awsscheduler

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

func (i *jsiiProxy_IScheduleGroup) validateGrantParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IScheduleGroup) validateGrantDeleteSchedulesParameters(identity awsiam.IGrantable) error {
	if identity == nil {
		return fmt.Errorf("parameter identity is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IScheduleGroup) validateGrantReadSchedulesParameters(identity awsiam.IGrantable) error {
	if identity == nil {
		return fmt.Errorf("parameter identity is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IScheduleGroup) validateGrantWriteSchedulesParameters(identity awsiam.IGrantable) error {
	if identity == nil {
		return fmt.Errorf("parameter identity is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IScheduleGroup) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	if metricName == nil {
		return fmt.Errorf("parameter metricName is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IScheduleGroup) validateMetricAttemptsParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IScheduleGroup) validateMetricDroppedParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IScheduleGroup) validateMetricFailedToBeSentToDLQParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IScheduleGroup) validateMetricSentToDLQParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IScheduleGroup) validateMetricSentToDLQTruncatedParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IScheduleGroup) validateMetricTargetErrorsParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IScheduleGroup) validateMetricTargetThrottledParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IScheduleGroup) validateMetricThrottledParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IScheduleGroup) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

