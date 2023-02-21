//go:build !no_runtime_type_checking

// The CDK Construct Library for AWS::GameLift
package awscdkgameliftalpha

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func (i *jsiiProxy_IRuleSetContent) validateBindParameters(_scope constructs.Construct) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	return nil
}

