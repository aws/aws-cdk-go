//go:build !no_runtime_type_checking

package awsmediaconnectalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

func (i *jsiiProxy_IBridge) validateAddOutputParameters(id *string, networkOutput *BridgeNetworkOutput) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if networkOutput == nil {
		return fmt.Errorf("parameter networkOutput is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(networkOutput, func() string { return "parameter networkOutput" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IBridge) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	if metricName == nil {
		return fmt.Errorf("parameter metricName is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IBridge) validateMetricFailoverSwitchesParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IBridge) validateMetricSourceBitrateParameters(bridgeSourceName *string, props *awscloudwatch.MetricOptions) error {
	if bridgeSourceName == nil {
		return fmt.Errorf("parameter bridgeSourceName is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IBridge) validateMetricSourcePacketLossPercentParameters(bridgeSourceName *string, props *awscloudwatch.MetricOptions) error {
	if bridgeSourceName == nil {
		return fmt.Errorf("parameter bridgeSourceName is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IBridge) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

