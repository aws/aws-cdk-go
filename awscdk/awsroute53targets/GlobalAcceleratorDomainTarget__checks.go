//go:build !no_runtime_type_checking

package awsroute53targets

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
)

func (g *jsiiProxy_GlobalAcceleratorDomainTarget) validateBindParameters(record awsroute53.IRecordSet) error {
	if record == nil {
		return fmt.Errorf("parameter record is required, but nil was provided")
	}

	return nil
}

func validateNewGlobalAcceleratorDomainTargetParameters(acceleratorDomainName *string) error {
	if acceleratorDomainName == nil {
		return fmt.Errorf("parameter acceleratorDomainName is required, but nil was provided")
	}

	return nil
}

