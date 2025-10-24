//go:build !no_runtime_type_checking

package awsecr

import (
	"fmt"
)

func validateImageTagMutabilityExclusionFilter_WildcardParameters(pattern *string) error {
	if pattern == nil {
		return fmt.Errorf("parameter pattern is required, but nil was provided")
	}

	return nil
}

