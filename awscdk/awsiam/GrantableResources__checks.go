//go:build !no_runtime_type_checking

package awsiam

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
)

func validateGrantableResources_IsEncryptedResourceParameters(resource interfaces.IEnvironmentAware) error {
	if resource == nil {
		return fmt.Errorf("parameter resource is required, but nil was provided")
	}

	return nil
}

func validateGrantableResources_IsResourceWithPolicyParameters(resource interfaces.IEnvironmentAware) error {
	if resource == nil {
		return fmt.Errorf("parameter resource is required, but nil was provided")
	}

	return nil
}

