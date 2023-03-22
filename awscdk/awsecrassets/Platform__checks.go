//go:build !no_runtime_type_checking

package awsecrassets

import (
	"fmt"
)

func validatePlatform_CustomParameters(platform *string) error {
	if platform == nil {
		return fmt.Errorf("parameter platform is required, but nil was provided")
	}

	return nil
}

