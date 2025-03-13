//go:build !no_runtime_type_checking

package awsappmesh

import (
	"fmt"
)

func validateLoggingFormat_FromJsonParameters(jsonLoggingFormat *map[string]*string) error {
	if jsonLoggingFormat == nil {
		return fmt.Errorf("parameter jsonLoggingFormat is required, but nil was provided")
	}

	return nil
}

func validateLoggingFormat_FromTextParameters(text *string) error {
	if text == nil {
		return fmt.Errorf("parameter text is required, but nil was provided")
	}

	return nil
}

