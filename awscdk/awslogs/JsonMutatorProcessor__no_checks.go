//go:build no_runtime_type_checking

package awslogs

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_JsonMutatorProcessor) validateSetTypeParameters(val JsonMutatorType) error {
	return nil
}

func validateNewJsonMutatorProcessorParameters(props *JsonMutatorProps) error {
	return nil
}

