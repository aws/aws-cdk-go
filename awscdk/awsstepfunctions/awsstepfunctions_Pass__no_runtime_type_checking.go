//go:build no_runtime_type_checking

package awsstepfunctions

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_Pass) validateAddBranchParameters(branch StateGraph) error {
	return nil
}

func (p *jsiiProxy_Pass) validateAddChoiceParameters(condition Condition, next State) error {
	return nil
}

func (p *jsiiProxy_Pass) validateAddIteratorParameters(iteration StateGraph) error {
	return nil
}

func (p *jsiiProxy_Pass) validateAddPrefixParameters(x *string) error {
	return nil
}

func (p *jsiiProxy_Pass) validateBindToGraphParameters(graph StateGraph) error {
	return nil
}

func (p *jsiiProxy_Pass) validateMakeDefaultParameters(def State) error {
	return nil
}

func (p *jsiiProxy_Pass) validateMakeNextParameters(next State) error {
	return nil
}

func (p *jsiiProxy_Pass) validateNextParameters(next IChainable) error {
	return nil
}

func (p *jsiiProxy_Pass) validateWhenBoundToGraphParameters(graph StateGraph) error {
	return nil
}

func validatePass_FilterNextablesParameters(states *[]State) error {
	return nil
}

func validatePass_FindReachableEndStatesParameters(start State, options *FindStateOptions) error {
	return nil
}

func validatePass_FindReachableStatesParameters(start State, options *FindStateOptions) error {
	return nil
}

func validatePass_IsConstructParameters(x interface{}) error {
	return nil
}

func validatePass_PrefixStatesParameters(root constructs.IConstruct, prefix *string) error {
	return nil
}

func validateNewPassParameters(scope constructs.Construct, id *string, props *PassProps) error {
	return nil
}

