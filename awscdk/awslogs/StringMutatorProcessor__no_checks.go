//go:build no_runtime_type_checking

package awslogs

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_StringMutatorProcessor) validateSetTypeParameters(val StringMutatorType) error {
	return nil
}

func validateNewStringMutatorProcessorParameters(props *StringMutatorProps) error {
	return nil
}

