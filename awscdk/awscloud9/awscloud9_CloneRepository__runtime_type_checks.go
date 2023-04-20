//go:build !no_runtime_type_checking

package awscloud9

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/awscodecommit"
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

