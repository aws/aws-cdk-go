//go:build !no_runtime_type_checking

package awslambda

import (
	"fmt"
)

func validateWebFunctionTraffic_SingleParameters(revisionId *string) error {
	if revisionId == nil {
		return fmt.Errorf("parameter revisionId is required, but nil was provided")
	}

	return nil
}

