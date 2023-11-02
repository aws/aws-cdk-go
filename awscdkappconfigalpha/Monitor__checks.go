//go:build !no_runtime_type_checking

package awscdkappconfigalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappconfig"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

func validateMonitor_FromCfnMonitorsPropertyParameters(monitorsProperty *awsappconfig.CfnEnvironment_MonitorsProperty) error {
	if monitorsProperty == nil {
		return fmt.Errorf("parameter monitorsProperty is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(monitorsProperty, func() string { return "parameter monitorsProperty" }); err != nil {
		return err
	}

	return nil
}

func validateMonitor_FromCloudWatchAlarmParameters(alarm awscloudwatch.IAlarm) error {
	if alarm == nil {
		return fmt.Errorf("parameter alarm is required, but nil was provided")
	}

	return nil
}

