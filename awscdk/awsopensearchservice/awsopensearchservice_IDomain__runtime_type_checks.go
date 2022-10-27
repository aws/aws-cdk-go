//go:build !no_runtime_type_checking

package awsopensearchservice

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

func (i *jsiiProxy_IDomain) validateGrantIndexReadParameters(index *string, identity awsiam.IGrantable) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	if identity == nil {
		return fmt.Errorf("parameter identity is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IDomain) validateGrantIndexReadWriteParameters(index *string, identity awsiam.IGrantable) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	if identity == nil {
		return fmt.Errorf("parameter identity is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IDomain) validateGrantIndexWriteParameters(index *string, identity awsiam.IGrantable) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	if identity == nil {
		return fmt.Errorf("parameter identity is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IDomain) validateGrantPathReadParameters(path *string, identity awsiam.IGrantable) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if identity == nil {
		return fmt.Errorf("parameter identity is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IDomain) validateGrantPathReadWriteParameters(path *string, identity awsiam.IGrantable) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if identity == nil {
		return fmt.Errorf("parameter identity is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IDomain) validateGrantPathWriteParameters(path *string, identity awsiam.IGrantable) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if identity == nil {
		return fmt.Errorf("parameter identity is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IDomain) validateGrantReadParameters(identity awsiam.IGrantable) error {
	if identity == nil {
		return fmt.Errorf("parameter identity is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IDomain) validateGrantReadWriteParameters(identity awsiam.IGrantable) error {
	if identity == nil {
		return fmt.Errorf("parameter identity is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IDomain) validateGrantWriteParameters(identity awsiam.IGrantable) error {
	if identity == nil {
		return fmt.Errorf("parameter identity is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IDomain) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	if metricName == nil {
		return fmt.Errorf("parameter metricName is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IDomain) validateMetricAutomatedSnapshotFailureParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IDomain) validateMetricClusterIndexWritesBlockedParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IDomain) validateMetricClusterStatusRedParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IDomain) validateMetricClusterStatusYellowParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IDomain) validateMetricCPUUtilizationParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IDomain) validateMetricFreeStorageSpaceParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IDomain) validateMetricIndexingLatencyParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IDomain) validateMetricJVMMemoryPressureParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IDomain) validateMetricKMSKeyErrorParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IDomain) validateMetricKMSKeyInaccessibleParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IDomain) validateMetricMasterCPUUtilizationParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IDomain) validateMetricMasterJVMMemoryPressureParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IDomain) validateMetricNodesParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IDomain) validateMetricSearchableDocumentsParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IDomain) validateMetricSearchLatencyParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

