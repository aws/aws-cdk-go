//go:build !no_runtime_type_checking

package previewawsstepfunctionsmixins

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
)

func (c *jsiiProxy_CfnStateMachineStandardLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef) error {
	if logGroup == nil {
		return fmt.Errorf("parameter logGroup is required, but nil was provided")
	}

	return nil
}

