//go:build !no_runtime_type_checking

package customresources

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

func (p *jsiiProxy_PhysicalResourceIdReference) validateResolveParameters(_arg awscdk.IResolveContext) error {
	if _arg == nil {
		return fmt.Errorf("parameter _arg is required, but nil was provided")
	}

	return nil
}

