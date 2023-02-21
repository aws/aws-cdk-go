//go:build no_runtime_type_checking

package awsstepfunctions

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Succeed) validateAddBranchParameters(branch StateGraph) error {
	return nil
}

func (s *jsiiProxy_Succeed) validateAddChoiceParameters(condition Condition, next State) error {
	return nil
}

func (s *jsiiProxy_Succeed) validateAddIteratorParameters(iteration StateGraph) error {
	return nil
}

func (s *jsiiProxy_Succeed) validateAddPrefixParameters(x *string) error {
	return nil
}

func (s *jsiiProxy_Succeed) validateBindToGraphParameters(graph StateGraph) error {
	return nil
}

func (s *jsiiProxy_Succeed) validateMakeDefaultParameters(def State) error {
	return nil
}

func (s *jsiiProxy_Succeed) validateMakeNextParameters(next State) error {
	return nil
}

func (s *jsiiProxy_Succeed) validateWhenBoundToGraphParameters(graph StateGraph) error {
	return nil
}

func validateSucceed_FilterNextablesParameters(states *[]State) error {
	return nil
}

func validateSucceed_FindReachableEndStatesParameters(start State, options *FindStateOptions) error {
	return nil
}

func validateSucceed_FindReachableStatesParameters(start State, options *FindStateOptions) error {
	return nil
}

func validateSucceed_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSucceed_PrefixStatesParameters(root constructs.IConstruct, prefix *string) error {
	return nil
}

func validateNewSucceedParameters(scope constructs.Construct, id *string, props *SucceedProps) error {
	return nil
}

