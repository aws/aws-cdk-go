//go:build !no_runtime_type_checking

package awscdkapprunneralpha

import (
	"fmt"
)

func validateGitHubConnection_FromConnectionArnParameters(arn *string) error {
	if arn == nil {
		return fmt.Errorf("parameter arn is required, but nil was provided")
	}

	return nil
}

func validateNewGitHubConnectionParameters(arn *string) error {
	if arn == nil {
		return fmt.Errorf("parameter arn is required, but nil was provided")
	}

	return nil
}

