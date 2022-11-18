//go:build no_runtime_type_checking

package awsstepfunctions

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_Map) validateAddBranchParameters(branch StateGraph) error {
	return nil
}

func (m *jsiiProxy_Map) validateAddCatchParameters(handler IChainable, props *CatchProps) error {
	return nil
}

func (m *jsiiProxy_Map) validateAddChoiceParameters(condition Condition, next State) error {
	return nil
}

func (m *jsiiProxy_Map) validateAddIteratorParameters(iteration StateGraph) error {
	return nil
}

func (m *jsiiProxy_Map) validateAddPrefixParameters(x *string) error {
	return nil
}

func (m *jsiiProxy_Map) validateAddRetryParameters(props *RetryProps) error {
	return nil
}

func (m *jsiiProxy_Map) validateBindToGraphParameters(graph StateGraph) error {
	return nil
}

func (m *jsiiProxy_Map) validateIteratorParameters(iterator IChainable) error {
	return nil
}

func (m *jsiiProxy_Map) validateMakeDefaultParameters(def State) error {
	return nil
}

func (m *jsiiProxy_Map) validateMakeNextParameters(next State) error {
	return nil
}

func (m *jsiiProxy_Map) validateNextParameters(next IChainable) error {
	return nil
}

func (m *jsiiProxy_Map) validateWhenBoundToGraphParameters(graph StateGraph) error {
	return nil
}

func validateMap_FilterNextablesParameters(states *[]State) error {
	return nil
}

func validateMap_FindReachableEndStatesParameters(start State, options *FindStateOptions) error {
	return nil
}

func validateMap_FindReachableStatesParameters(start State, options *FindStateOptions) error {
	return nil
}

func validateMap_IsConstructParameters(x interface{}) error {
	return nil
}

func validateMap_PrefixStatesParameters(root constructs.IConstruct, prefix *string) error {
	return nil
}

func validateNewMapParameters(scope constructs.Construct, id *string, props *MapProps) error {
	return nil
}

