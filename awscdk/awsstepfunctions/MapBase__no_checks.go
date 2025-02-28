//go:build no_runtime_type_checking

package awsstepfunctions

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MapBase) validateAddBranchParameters(branch StateGraph) error {
	return nil
}

func (m *jsiiProxy_MapBase) validateAddChoiceParameters(condition Condition, next State, options *ChoiceTransitionOptions) error {
	return nil
}

func (m *jsiiProxy_MapBase) validateAddItemProcessorParameters(processor StateGraph, config *ProcessorConfig) error {
	return nil
}

func (m *jsiiProxy_MapBase) validateAddIteratorParameters(iteration StateGraph) error {
	return nil
}

func (m *jsiiProxy_MapBase) validateAddPrefixParameters(x *string) error {
	return nil
}

func (m *jsiiProxy_MapBase) validateBindToGraphParameters(graph StateGraph) error {
	return nil
}

func (m *jsiiProxy_MapBase) validateMakeDefaultParameters(def State) error {
	return nil
}

func (m *jsiiProxy_MapBase) validateMakeNextParameters(next State) error {
	return nil
}

func (m *jsiiProxy_MapBase) validateNextParameters(next IChainable) error {
	return nil
}

func (m *jsiiProxy_MapBase) validateWhenBoundToGraphParameters(graph StateGraph) error {
	return nil
}

func validateMapBase_FilterNextablesParameters(states *[]State) error {
	return nil
}

func validateMapBase_FindReachableEndStatesParameters(start State, options *FindStateOptions) error {
	return nil
}

func validateMapBase_FindReachableStatesParameters(start State, options *FindStateOptions) error {
	return nil
}

func validateMapBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateMapBase_PrefixStatesParameters(root constructs.IConstruct, prefix *string) error {
	return nil
}

func (j *jsiiProxy_MapBase) validateSetProcessorConfigParameters(val *ProcessorConfig) error {
	return nil
}

func validateNewMapBaseParameters(scope constructs.Construct, id *string, props *MapBaseProps) error {
	return nil
}

