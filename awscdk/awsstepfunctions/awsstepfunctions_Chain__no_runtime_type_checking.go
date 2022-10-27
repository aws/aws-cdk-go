//go:build no_runtime_type_checking

package awsstepfunctions

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_Chain) validateNextParameters(next IChainable) error {
	return nil
}

func (c *jsiiProxy_Chain) validateToSingleStateParameters(id *string, props *ParallelProps) error {
	return nil
}

func validateChain_CustomParameters(startState State, endStates *[]INextable, lastAdded IChainable) error {
	return nil
}

func validateChain_SequenceParameters(start IChainable, next IChainable) error {
	return nil
}

func validateChain_StartParameters(state IChainable) error {
	return nil
}

