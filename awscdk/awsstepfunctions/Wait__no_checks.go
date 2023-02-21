//go:build no_runtime_type_checking

package awsstepfunctions

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_Wait) validateAddBranchParameters(branch StateGraph) error {
	return nil
}

func (w *jsiiProxy_Wait) validateAddChoiceParameters(condition Condition, next State) error {
	return nil
}

func (w *jsiiProxy_Wait) validateAddIteratorParameters(iteration StateGraph) error {
	return nil
}

func (w *jsiiProxy_Wait) validateAddPrefixParameters(x *string) error {
	return nil
}

func (w *jsiiProxy_Wait) validateBindToGraphParameters(graph StateGraph) error {
	return nil
}

func (w *jsiiProxy_Wait) validateMakeDefaultParameters(def State) error {
	return nil
}

func (w *jsiiProxy_Wait) validateMakeNextParameters(next State) error {
	return nil
}

func (w *jsiiProxy_Wait) validateNextParameters(next IChainable) error {
	return nil
}

func (w *jsiiProxy_Wait) validateWhenBoundToGraphParameters(graph StateGraph) error {
	return nil
}

func validateWait_FilterNextablesParameters(states *[]State) error {
	return nil
}

func validateWait_FindReachableEndStatesParameters(start State, options *FindStateOptions) error {
	return nil
}

func validateWait_FindReachableStatesParameters(start State, options *FindStateOptions) error {
	return nil
}

func validateWait_IsConstructParameters(x interface{}) error {
	return nil
}

func validateWait_PrefixStatesParameters(root constructs.IConstruct, prefix *string) error {
	return nil
}

func validateNewWaitParameters(scope constructs.Construct, id *string, props *WaitProps) error {
	return nil
}

