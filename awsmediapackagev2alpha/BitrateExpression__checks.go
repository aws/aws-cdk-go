//go:build !no_runtime_type_checking

package awsmediapackagev2alpha

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

func validateBitrateExpression_RangeParameters(start awscdk.Bitrate, end awscdk.Bitrate) error {
	if start == nil {
		return fmt.Errorf("parameter start is required, but nil was provided")
	}

	if end == nil {
		return fmt.Errorf("parameter end is required, but nil was provided")
	}

	return nil
}

func validateBitrateExpression_ValueParameters(v awscdk.Bitrate) error {
	if v == nil {
		return fmt.Errorf("parameter v is required, but nil was provided")
	}

	return nil
}

func validateNewBitrateExpressionParameters(_expression *string, _values *[]*float64) error {
	if _expression == nil {
		return fmt.Errorf("parameter _expression is required, but nil was provided")
	}

	if _values == nil {
		return fmt.Errorf("parameter _values is required, but nil was provided")
	}

	return nil
}

