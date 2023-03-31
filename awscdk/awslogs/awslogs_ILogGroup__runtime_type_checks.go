//go:build !no_runtime_type_checking

package awslogs

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

func (i *jsiiProxy_ILogGroup) validateAddMetricFilterParameters(id *string, props *MetricFilterOptions) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_ILogGroup) validateAddStreamParameters(id *string, props *StreamOptions) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_ILogGroup) validateAddSubscriptionFilterParameters(id *string, props *SubscriptionFilterOptions) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_ILogGroup) validateExtractMetricParameters(jsonField *string, metricNamespace *string, metricName *string) error {
	if jsonField == nil {
		return fmt.Errorf("parameter jsonField is required, but nil was provided")
	}

	if metricNamespace == nil {
		return fmt.Errorf("parameter metricNamespace is required, but nil was provided")
	}

	if metricName == nil {
		return fmt.Errorf("parameter metricName is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_ILogGroup) validateGrantParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_ILogGroup) validateGrantWriteParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

