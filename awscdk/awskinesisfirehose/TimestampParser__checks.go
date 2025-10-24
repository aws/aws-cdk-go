//go:build !no_runtime_type_checking

package awskinesisfirehose

import (
	"fmt"
)

func validateTimestampParser_FromFormatStringParameters(format *string) error {
	if format == nil {
		return fmt.Errorf("parameter format is required, but nil was provided")
	}

	return nil
}

