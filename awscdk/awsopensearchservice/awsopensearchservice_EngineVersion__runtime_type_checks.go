//go:build !no_runtime_type_checking

package awsopensearchservice

import (
	"fmt"
)

func validateEngineVersion_ElasticsearchParameters(version *string) error {
	if version == nil {
		return fmt.Errorf("parameter version is required, but nil was provided")
	}

	return nil
}

func validateEngineVersion_OpenSearchParameters(version *string) error {
	if version == nil {
		return fmt.Errorf("parameter version is required, but nil was provided")
	}

	return nil
}

