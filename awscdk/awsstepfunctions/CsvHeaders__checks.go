//go:build !no_runtime_type_checking

package awsstepfunctions

import (
	"fmt"
)

func validateCsvHeaders_UseParameters(headers *[]*string) error {
	if headers == nil {
		return fmt.Errorf("parameter headers is required, but nil was provided")
	}

	return nil
}

