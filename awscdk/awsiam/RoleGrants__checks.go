//go:build !no_runtime_type_checking

package awsiam

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiam"
)

func (r *jsiiProxy_RoleGrants) validateAssumeRoleParameters(identity IPrincipal) error {
	if identity == nil {
		return fmt.Errorf("parameter identity is required, but nil was provided")
	}

	return nil
}

func (r *jsiiProxy_RoleGrants) validatePassRoleParameters(identity IPrincipal) error {
	if identity == nil {
		return fmt.Errorf("parameter identity is required, but nil was provided")
	}

	return nil
}

func validateRoleGrants_FromRoleParameters(role interfacesawsiam.IRoleRef) error {
	if role == nil {
		return fmt.Errorf("parameter role is required, but nil was provided")
	}

	return nil
}

