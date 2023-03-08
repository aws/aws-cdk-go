//go:build no_runtime_type_checking

package awsstepfunctions

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_Fail) validateAddBranchParameters(branch StateGraph) error {
	return nil
}

func (f *jsiiProxy_Fail) validateAddChoiceParameters(condition Condition, next State) error {
	return nil
}

func (f *jsiiProxy_Fail) validateAddIteratorParameters(iteration StateGraph) error {
	return nil
}

func (f *jsiiProxy_Fail) validateAddPrefixParameters(x *string) error {
	return nil
}

func (f *jsiiProxy_Fail) validateBindToGraphParameters(graph StateGraph) error {
	return nil
}

func (f *jsiiProxy_Fail) validateMakeDefaultParameters(def State) error {
	return nil
}

func (f *jsiiProxy_Fail) validateMakeNextParameters(next State) error {
	return nil
}

func (f *jsiiProxy_Fail) validateWhenBoundToGraphParameters(graph StateGraph) error {
	return nil
}

func validateFail_FilterNextablesParameters(states *[]State) error {
	return nil
}

func validateFail_FindReachableEndStatesParameters(start State, options *FindStateOptions) error {
	return nil
}

func validateFail_FindReachableStatesParameters(start State, options *FindStateOptions) error {
	return nil
}

func validateFail_IsConstructParameters(x interface{}) error {
	return nil
}

func validateFail_PrefixStatesParameters(root constructs.IConstruct, prefix *string) error {
	return nil
}

func validateNewFailParameters(scope constructs.Construct, id *string, props *FailProps) error {
	return nil
}

