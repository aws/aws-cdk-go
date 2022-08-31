//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsbatch

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_IMultiNodeProps) validateSetCountParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_IMultiNodeProps) validateSetMainNodeParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_IMultiNodeProps) validateSetRangePropsParameters(val *[]INodeRangeProps) error {
	return nil
}

