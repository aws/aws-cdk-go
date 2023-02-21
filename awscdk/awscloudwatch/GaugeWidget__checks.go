//go:build !no_runtime_type_checking

package awscloudwatch

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (g *jsiiProxy_GaugeWidget) validateAddMetricParameters(metric IMetric) error {
	if metric == nil {
		return fmt.Errorf("parameter metric is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GaugeWidget) validatePositionParameters(x *float64, y *float64) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	if y == nil {
		return fmt.Errorf("parameter y is required, but nil was provided")
	}

	return nil
}

func validateNewGaugeWidgetParameters(props *GaugeWidgetProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

