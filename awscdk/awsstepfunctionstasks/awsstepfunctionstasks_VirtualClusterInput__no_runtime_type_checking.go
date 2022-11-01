//go:build no_runtime_type_checking

package awsstepfunctionstasks

// Building without runtime type checking enabled, so all the below just return nil

func validateVirtualClusterInput_FromTaskInputParameters(taskInput awsstepfunctions.TaskInput) error {
	return nil
}

func validateVirtualClusterInput_FromVirtualClusterIdParameters(virtualClusterId *string) error {
	return nil
}

