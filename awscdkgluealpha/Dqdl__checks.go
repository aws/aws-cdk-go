//go:build !no_runtime_type_checking

package awscdkgluealpha

import (
	"fmt"
)

func validateDqdl_FromStringParameters(dqdl *string) error {
	if dqdl == nil {
		return fmt.Errorf("parameter dqdl is required, but nil was provided")
	}

	return nil
}

