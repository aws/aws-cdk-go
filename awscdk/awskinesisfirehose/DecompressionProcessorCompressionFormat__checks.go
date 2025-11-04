//go:build !no_runtime_type_checking

package awskinesisfirehose

import (
	"fmt"
)

func validateDecompressionProcessorCompressionFormat_OfParameters(compressionFormat *string) error {
	if compressionFormat == nil {
		return fmt.Errorf("parameter compressionFormat is required, but nil was provided")
	}

	return nil
}

