//go:build !no_runtime_type_checking

package awsiam

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/constructs-go/constructs/v10"
)

func validateGrantableResources_IsEncryptedResourceParameters(resource constructs.IConstruct) error {
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

