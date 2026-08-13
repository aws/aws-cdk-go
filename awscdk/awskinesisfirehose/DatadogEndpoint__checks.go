//go:build !no_runtime_type_checking

package awskinesisfirehose

import (
	"fmt"
)

func validateDatadogEndpoint_OfParameters(url *string) error {
	if url == nil {
		return fmt.Errorf("parameter url is required, but nil was provided")
	}

	return nil
}

