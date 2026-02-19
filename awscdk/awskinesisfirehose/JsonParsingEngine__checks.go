//go:build !no_runtime_type_checking

package awskinesisfirehose

import (
	"fmt"
)

func validateJsonParsingEngine_OfParameters(parsingEngine *string) error {
	if parsingEngine == nil {
		return fmt.Errorf("parameter parsingEngine is required, but nil was provided")
	}

	return nil
}

