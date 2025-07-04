//go:build !no_runtime_type_checking

package appstagingsynthesizeralpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateDeploymentIdentities_DefaultBootstrapRolesParameters(options *DefaultBootstrapRolesOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateDeploymentIdentities_SpecifyRolesParameters(roles *BootstrapRoles) error {
	if roles == nil {
		return fmt.Errorf("parameter roles is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(roles, func() string { return "parameter roles" }); err != nil {
		return err
	}

	return nil
}

