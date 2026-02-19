//go:build !no_runtime_type_checking

package awsiam

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/constructs-go/constructs/v10"
)

func validateEncryptedResources_OfParameters(resource interfaces.IEnvironmentAware) error {
	if resource == nil {
		return fmt.Errorf("parameter resource is required, but nil was provided")
	}

	return nil
}

func validateEncryptedResources_RegisterParameters(scope constructs.IConstruct, cfnType *string, factory IEncryptedResourceFactory) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if cfnType == nil {
		return fmt.Errorf("parameter cfnType is required, but nil was provided")
	}

	if factory == nil {
		return fmt.Errorf("parameter factory is required, but nil was provided")
	}

	return nil
}

