//go:build !no_runtime_type_checking

package awsecs

import (
	"fmt"
)

func validateCpuArchitecture_OfParameters(cpuArchitecture *string) error {
	if cpuArchitecture == nil {
		return fmt.Errorf("parameter cpuArchitecture is required, but nil was provided")
	}

	return nil
}

