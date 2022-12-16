//go:build no_runtime_type_checking

package awsstepfunctions

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CustomState) validateAddBranchParameters(branch StateGraph) error {
	return nil
}

func (c *jsiiProxy_CustomState) validateAddChoiceParameters(condition Condition, next State) error {
	return nil
}

func (c *jsiiProxy_CustomState) validateAddIteratorParameters(iteration StateGraph) error {
	return nil
}

func (c *jsiiProxy_CustomState) validateAddPrefixParameters(x *string) error {
	return nil
}

func (c *jsiiProxy_CustomState) validateBindToGraphParameters(graph StateGraph) error {
	return nil
}

func (c *jsiiProxy_CustomState) validateMakeDefaultParameters(def State) error {
	return nil
}

func (c *jsiiProxy_CustomState) validateMakeNextParameters(next State) error {
	return nil
}

func (c *jsiiProxy_CustomState) validateNextParameters(next IChainable) error {
	return nil
}

func (c *jsiiProxy_CustomState) validateWhenBoundToGraphParameters(graph StateGraph) error {
	return nil
}

func validateCustomState_FilterNextablesParameters(states *[]State) error {
	return nil
}

func validateCustomState_FindReachableEndStatesParameters(start State, options *FindStateOptions) error {
	return nil
}

func validateCustomState_FindReachableStatesParameters(start State, options *FindStateOptions) error {
	return nil
}

func validateCustomState_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCustomState_PrefixStatesParameters(root constructs.IConstruct, prefix *string) error {
	return nil
}

func validateNewCustomStateParameters(scope constructs.Construct, id *string, props *CustomStateProps) error {
	return nil
}

