//go:build !no_runtime_type_checking

package awsecs

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
)

func validateCredentialSpec_ArnForS3ObjectParameters(bucket awss3.IBucket, key *string) error {
	if bucket == nil {
		return fmt.Errorf("parameter bucket is required, but nil was provided")
	}

	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func validateCredentialSpec_ArnForSsmParameterParameters(parameter awsssm.IParameter) error {
	if parameter == nil {
		return fmt.Errorf("parameter parameter is required, but nil was provided")
	}

	return nil
}

func validateNewCredentialSpecParameters(prefixId *string, fileLocation *string) error {
	if prefixId == nil {
		return fmt.Errorf("parameter prefixId is required, but nil was provided")
	}

	if fileLocation == nil {
		return fmt.Errorf("parameter fileLocation is required, but nil was provided")
	}

	return nil
}

