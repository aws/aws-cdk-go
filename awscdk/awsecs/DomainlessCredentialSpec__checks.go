//go:build !no_runtime_type_checking

package awsecs

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
)

func validateDomainlessCredentialSpec_ArnForS3ObjectParameters(bucket awss3.IBucket, key *string) error {
	if bucket == nil {
		return fmt.Errorf("parameter bucket is required, but nil was provided")
	}

	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func validateDomainlessCredentialSpec_ArnForSsmParameterParameters(parameter awsssm.IParameter) error {
	if parameter == nil {
		return fmt.Errorf("parameter parameter is required, but nil was provided")
	}

	return nil
}

func validateDomainlessCredentialSpec_FromS3BucketParameters(bucket awss3.IBucket, key *string) error {
	if bucket == nil {
		return fmt.Errorf("parameter bucket is required, but nil was provided")
	}

	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func validateDomainlessCredentialSpec_FromSsmParameterParameters(parameter awsssm.IParameter) error {
	if parameter == nil {
		return fmt.Errorf("parameter parameter is required, but nil was provided")
	}

	return nil
}

func validateNewDomainlessCredentialSpecParameters(fileLocation *string) error {
	if fileLocation == nil {
		return fmt.Errorf("parameter fileLocation is required, but nil was provided")
	}

	return nil
}

