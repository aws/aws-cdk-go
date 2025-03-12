//go:build !no_runtime_type_checking

package awsbedrock

import (
	"fmt"
)

func validateNewFoundationModelIdentifierParameters(modelId *string) error {
	if modelId == nil {
		return fmt.Errorf("parameter modelId is required, but nil was provided")
	}

	return nil
}

