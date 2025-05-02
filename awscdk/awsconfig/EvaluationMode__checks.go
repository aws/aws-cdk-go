//go:build !no_runtime_type_checking

package awsconfig

import (
	"fmt"
)

func validateNewEvaluationModeParameters(modes *[]*string) error {
	if modes == nil {
		return fmt.Errorf("parameter modes is required, but nil was provided")
	}

	return nil
}

