//go:build no_runtime_type_checking

package awsstepfunctions

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_State) validateAddBranchParameters(branch StateGraph) error {
	return nil
}

func (s *jsiiProxy_State) validateAddChoiceParameters(condition Condition, next State) error {
	return nil
}

func (s *jsiiProxy_State) validateAddIteratorParameters(iteration StateGraph) error {
	return nil
}

func (s *jsiiProxy_State) validateAddPrefixParameters(x *string) error {
	return nil
}

func (s *jsiiProxy_State) validateBindToGraphParameters(graph StateGraph) error {
	return nil
}

func (s *jsiiProxy_State) validateMakeDefaultParameters(def State) error {
	return nil
}

func (s *jsiiProxy_State) validateMakeNextParameters(next State) error {
	return nil
}

func (s *jsiiProxy_State) validateWhenBoundToGraphParameters(graph StateGraph) error {
	return nil
}

func validateState_FilterNextablesParameters(states *[]State) error {
	return nil
}

func validateState_FindReachableEndStatesParameters(start State, options *FindStateOptions) error {
	return nil
}

func validateState_FindReachableStatesParameters(start State, options *FindStateOptions) error {
	return nil
}

func validateState_IsConstructParameters(x interface{}) error {
	return nil
}

func validateState_PrefixStatesParameters(root constructs.IConstruct, prefix *string) error {
	return nil
}

func validateNewStateParameters(scope constructs.Construct, id *string, props *StateProps) error {
	return nil
}

