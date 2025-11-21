//go:build !no_runtime_type_checking

package awsiam

import (
	"fmt"
)

func (i *jsiiProxy_IEncryptedResource) validateGrantOnKeyParameters(grantee IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

