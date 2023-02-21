//go:build no_runtime_type_checking

package awsstepfunctions

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_Parallel) validateAddBranchParameters(branch StateGraph) error {
	return nil
}

func (p *jsiiProxy_Parallel) validateAddCatchParameters(handler IChainable, props *CatchProps) error {
	return nil
}

func (p *jsiiProxy_Parallel) validateAddChoiceParameters(condition Condition, next State) error {
	return nil
}

func (p *jsiiProxy_Parallel) validateAddIteratorParameters(iteration StateGraph) error {
	return nil
}

func (p *jsiiProxy_Parallel) validateAddPrefixParameters(x *string) error {
	return nil
}

func (p *jsiiProxy_Parallel) validateAddRetryParameters(props *RetryProps) error {
	return nil
}

func (p *jsiiProxy_Parallel) validateBindToGraphParameters(graph StateGraph) error {
	return nil
}

func (p *jsiiProxy_Parallel) validateMakeDefaultParameters(def State) error {
	return nil
}

func (p *jsiiProxy_Parallel) validateMakeNextParameters(next State) error {
	return nil
}

func (p *jsiiProxy_Parallel) validateNextParameters(next IChainable) error {
	return nil
}

func (p *jsiiProxy_Parallel) validateWhenBoundToGraphParameters(graph StateGraph) error {
	return nil
}

func validateParallel_FilterNextablesParameters(states *[]State) error {
	return nil
}

func validateParallel_FindReachableEndStatesParameters(start State, options *FindStateOptions) error {
	return nil
}

func validateParallel_FindReachableStatesParameters(start State, options *FindStateOptions) error {
	return nil
}

func validateParallel_IsConstructParameters(x interface{}) error {
	return nil
}

func validateParallel_PrefixStatesParameters(root constructs.IConstruct, prefix *string) error {
	return nil
}

func validateNewParallelParameters(scope constructs.Construct, id *string, props *ParallelProps) error {
	return nil
}

