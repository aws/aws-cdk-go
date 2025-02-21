//go:build !no_runtime_type_checking

package awscdkeks-v2alpha

import (
	"fmt"
)

func validateKubernetesVersion_OfParameters(version *string) error {
	if version == nil {
		return fmt.Errorf("parameter version is required, but nil was provided")
	}

	return nil
}

