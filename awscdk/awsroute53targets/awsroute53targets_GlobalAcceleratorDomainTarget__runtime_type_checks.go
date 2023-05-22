//go:build !no_runtime_type_checking

package awsroute53targets

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/awsroute53"
)

func (g *jsiiProxy_GlobalAcceleratorDomainTarget) validateBindParameters(_record awsroute53.IRecordSet) error {
	if _record == nil {
		return fmt.Errorf("parameter _record is required, but nil was provided")
	}

	return nil
}

func validateNewGlobalAcceleratorDomainTargetParameters(acceleratorDomainName *string) error {
	if acceleratorDomainName == nil {
		return fmt.Errorf("parameter acceleratorDomainName is required, but nil was provided")
	}

	return nil
}

