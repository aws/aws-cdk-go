//go:build !no_runtime_type_checking

package awscodepipelineactions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateStackSetDeploymentModel_OrganizationsParameters(props *OrganizationsDeploymentProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateStackSetDeploymentModel_SelfManagedParameters(props *SelfManagedDeploymentProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

