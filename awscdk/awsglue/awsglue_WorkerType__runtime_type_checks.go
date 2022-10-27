//go:build !no_runtime_type_checking

package awsglue

import (
	"fmt"
)

func validateWorkerType_OfParameters(workerType *string) error {
	if workerType == nil {
		return fmt.Errorf("parameter workerType is required, but nil was provided")
	}

	return nil
}

