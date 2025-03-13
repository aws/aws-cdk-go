//go:build !no_runtime_type_checking

package awscdkcloud9alpha

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodecommit"
)

func validateCloneRepository_FromCodeCommitParameters(repository awscodecommit.IRepository, path *string) error {
	if repository == nil {
		return fmt.Errorf("parameter repository is required, but nil was provided")
	}

	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

