//go:build !no_runtime_type_checking

package awscdkgluealpha

import (
	"fmt"
)

func validateNewSerializationLibraryParameters(className *string) error {
	if className == nil {
		return fmt.Errorf("parameter className is required, but nil was provided")
	}

	return nil
}

