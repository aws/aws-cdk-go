//go:build !no_runtime_type_checking

package awsiam

import (
	"fmt"
)

func (i *jsiiProxy_IRole) validateGrantParameters(grantee IPrincipal) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IRole) validateGrantAssumeRoleParameters(grantee IPrincipal) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IRole) validateGrantPassRoleParameters(grantee IPrincipal) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

