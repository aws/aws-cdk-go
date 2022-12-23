//go:build !no_runtime_type_checking

package awsiam

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func (p *jsiiProxy_PermissionsBoundary) validateApplyParameters(boundaryPolicy IManagedPolicy) error {
	if boundaryPolicy == nil {
		return fmt.Errorf("parameter boundaryPolicy is required, but nil was provided")
	}

	return nil
}

func validatePermissionsBoundary_OfParameters(scope constructs.IConstruct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

