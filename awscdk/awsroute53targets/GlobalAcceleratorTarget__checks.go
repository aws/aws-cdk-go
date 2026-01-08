//go:build !no_runtime_type_checking

package awsroute53targets

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsglobalaccelerator"
)

func (g *jsiiProxy_GlobalAcceleratorTarget) validateBindParameters(record awsroute53.IRecordSet) error {
	if record == nil {
		return fmt.Errorf("parameter record is required, but nil was provided")
	}

	return nil
}

func validateNewGlobalAcceleratorTargetParameters(accelerator interfacesawsglobalaccelerator.IAcceleratorRef) error {
	if accelerator == nil {
		return fmt.Errorf("parameter accelerator is required, but nil was provided")
	}

	return nil
}

