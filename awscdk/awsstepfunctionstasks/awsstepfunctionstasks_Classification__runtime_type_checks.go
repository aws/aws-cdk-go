//go:build !no_runtime_type_checking

package awsstepfunctionstasks

import (
	"fmt"
)

func validateNewClassificationParameters(classificationStatement *string) error {
	if classificationStatement == nil {
		return fmt.Errorf("parameter classificationStatement is required, but nil was provided")
	}

	return nil
}

