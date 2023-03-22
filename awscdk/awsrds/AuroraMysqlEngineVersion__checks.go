//go:build !no_runtime_type_checking

package awsrds

import (
	"fmt"
)

func validateAuroraMysqlEngineVersion_OfParameters(auroraMysqlFullVersion *string) error {
	if auroraMysqlFullVersion == nil {
		return fmt.Errorf("parameter auroraMysqlFullVersion is required, but nil was provided")
	}

	return nil
}

