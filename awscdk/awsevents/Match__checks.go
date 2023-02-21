//go:build !no_runtime_type_checking

package awsevents

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

func (m *jsiiProxy_Match) validateResolveParameters(context awscdk.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func validateMatch_AnythingButPrefixParameters(prefix *string) error {
	if prefix == nil {
		return fmt.Errorf("parameter prefix is required, but nil was provided")
	}

	return nil
}

func validateMatch_CidrParameters(range_ *string) error {
	if range_ == nil {
		return fmt.Errorf("parameter range_ is required, but nil was provided")
	}

	return nil
}

func validateMatch_EqualParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateMatch_EqualsIgnoreCaseParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateMatch_ExactStringParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateMatch_GreaterThanParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateMatch_GreaterThanOrEqualParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateMatch_IntervalParameters(lower *float64, upper *float64) error {
	if lower == nil {
		return fmt.Errorf("parameter lower is required, but nil was provided")
	}

	if upper == nil {
		return fmt.Errorf("parameter upper is required, but nil was provided")
	}

	return nil
}

func validateMatch_IpAddressRangeParameters(range_ *string) error {
	if range_ == nil {
		return fmt.Errorf("parameter range_ is required, but nil was provided")
	}

	return nil
}

func validateMatch_LessThanParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateMatch_LessThanOrEqualParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateMatch_PrefixParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateMatch_SuffixParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

