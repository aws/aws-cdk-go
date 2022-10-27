//go:build no_runtime_type_checking

package awsbatch

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_INodeRangeProps) validateSetContainerParameters(val *JobDefinitionContainer) error {
	return nil
}

