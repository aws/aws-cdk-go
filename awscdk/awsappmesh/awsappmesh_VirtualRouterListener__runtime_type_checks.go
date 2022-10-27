//go:build !no_runtime_type_checking

package awsappmesh

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk"
)

func (v *jsiiProxy_VirtualRouterListener) validateBindParameters(scope awscdk.Construct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

