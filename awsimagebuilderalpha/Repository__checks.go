//go:build !no_runtime_type_checking

package awsimagebuilderalpha

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
)

func validateRepository_FromEcrParameters(repository awsecr.IRepository) error {
	if repository == nil {
		return fmt.Errorf("parameter repository is required, but nil was provided")
	}

	return nil
}

