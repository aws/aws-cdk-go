//go:build !no_runtime_type_checking

package customresources

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

func (p *jsiiProxy_PhysicalResourceIdReference) validateResolveParameters(context awscdk.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

