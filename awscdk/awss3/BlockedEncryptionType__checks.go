//go:build !no_runtime_type_checking

package awss3

import (
	"fmt"
)

func validateBlockedEncryptionType_CustomParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

