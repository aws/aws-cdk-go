//go:build !no_runtime_type_checking

package awsbedrockagentcorealpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateEvaluatorRatingScale_CategoricalParameters(options *[]*CategoricalRatingOption) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	for idx_a793ab, v := range *options {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter options[%#v]", idx_a793ab) }); err != nil {
			return err
		}
	}

	return nil
}

func validateEvaluatorRatingScale_NumericalParameters(options *[]*NumericalRatingOption) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	for idx_a793ab, v := range *options {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter options[%#v]", idx_a793ab) }); err != nil {
			return err
		}
	}

	return nil
}

