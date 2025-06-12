//go:build !no_runtime_type_checking

package awselasticloadbalancingv2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateListenerCondition_HostHeadersParameters(values *[]*string) error {
	if values == nil {
		return fmt.Errorf("parameter values is required, but nil was provided")
	}

	return nil
}

func validateListenerCondition_HttpHeaderParameters(name *string, values *[]*string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if values == nil {
		return fmt.Errorf("parameter values is required, but nil was provided")
	}

	return nil
}

func validateListenerCondition_HttpRequestMethodsParameters(values *[]*string) error {
	if values == nil {
		return fmt.Errorf("parameter values is required, but nil was provided")
	}

	return nil
}

func validateListenerCondition_PathPatternsParameters(values *[]*string) error {
	if values == nil {
		return fmt.Errorf("parameter values is required, but nil was provided")
	}

	return nil
}

func validateListenerCondition_QueryStringsParameters(values *[]*QueryStringCondition) error {
	if values == nil {
		return fmt.Errorf("parameter values is required, but nil was provided")
	}
	for idx_89445e, v := range *values {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter values[%#v]", idx_89445e) }); err != nil {
			return err
		}
	}

	return nil
}

func validateListenerCondition_SourceIpsParameters(values *[]*string) error {
	if values == nil {
		return fmt.Errorf("parameter values is required, but nil was provided")
	}

	return nil
}

