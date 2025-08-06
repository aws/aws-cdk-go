//go:build !no_runtime_type_checking

package awsecs

import (
	"fmt"
)

func validateTaskDefinitionRevision_OfParameters(revision *float64) error {
	if revision == nil {
		return fmt.Errorf("parameter revision is required, but nil was provided")
	}

	return nil
}

