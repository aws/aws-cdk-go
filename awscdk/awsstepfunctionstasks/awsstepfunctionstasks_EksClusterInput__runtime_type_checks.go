//go:build !no_runtime_type_checking

package awsstepfunctionstasks

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/awseks"
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctions"
)

func validateEksClusterInput_FromClusterParameters(cluster awseks.ICluster) error {
	if cluster == nil {
		return fmt.Errorf("parameter cluster is required, but nil was provided")
	}

	return nil
}

func validateEksClusterInput_FromTaskInputParameters(taskInput awsstepfunctions.TaskInput) error {
	if taskInput == nil {
		return fmt.Errorf("parameter taskInput is required, but nil was provided")
	}

	return nil
}

