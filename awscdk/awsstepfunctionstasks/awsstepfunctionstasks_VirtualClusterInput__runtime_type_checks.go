//go:build !no_runtime_type_checking

package awsstepfunctionstasks

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctions"
)

func validateVirtualClusterInput_FromTaskInputParameters(taskInput awsstepfunctions.TaskInput) error {
	if taskInput == nil {
		return fmt.Errorf("parameter taskInput is required, but nil was provided")
	}

	return nil
}

func validateVirtualClusterInput_FromVirtualClusterIdParameters(virtualClusterId *string) error {
	if virtualClusterId == nil {
		return fmt.Errorf("parameter virtualClusterId is required, but nil was provided")
	}

	return nil
}

