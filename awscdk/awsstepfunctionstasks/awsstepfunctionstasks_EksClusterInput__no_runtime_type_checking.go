//go:build no_runtime_type_checking

package awsstepfunctionstasks

// Building without runtime type checking enabled, so all the below just return nil

func validateEksClusterInput_FromClusterParameters(cluster awseks.ICluster) error {
	return nil
}

func validateEksClusterInput_FromTaskInputParameters(taskInput awsstepfunctions.TaskInput) error {
	return nil
}

