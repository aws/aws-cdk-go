//go:build no_runtime_type_checking

package awsstepfunctions

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_Choice) validateAddBranchParameters(branch StateGraph) error {
	return nil
}

func (c *jsiiProxy_Choice) validateAddChoiceParameters(condition Condition, next State) error {
	return nil
}

func (c *jsiiProxy_Choice) validateAddIteratorParameters(iteration StateGraph) error {
	return nil
}

func (c *jsiiProxy_Choice) validateAddPrefixParameters(x *string) error {
	return nil
}

func (c *jsiiProxy_Choice) validateAfterwardsParameters(options *AfterwardsOptions) error {
	return nil
}

func (c *jsiiProxy_Choice) validateBindToGraphParameters(graph StateGraph) error {
	return nil
}

func (c *jsiiProxy_Choice) validateMakeDefaultParameters(def State) error {
	return nil
}

func (c *jsiiProxy_Choice) validateMakeNextParameters(next State) error {
	return nil
}

func (c *jsiiProxy_Choice) validateOtherwiseParameters(def IChainable) error {
	return nil
}

func (c *jsiiProxy_Choice) validateWhenParameters(condition Condition, next IChainable) error {
	return nil
}

func (c *jsiiProxy_Choice) validateWhenBoundToGraphParameters(graph StateGraph) error {
	return nil
}

func validateChoice_FilterNextablesParameters(states *[]State) error {
	return nil
}

func validateChoice_FindReachableEndStatesParameters(start State, options *FindStateOptions) error {
	return nil
}

func validateChoice_FindReachableStatesParameters(start State, options *FindStateOptions) error {
	return nil
}

func validateChoice_IsConstructParameters(x interface{}) error {
	return nil
}

func validateChoice_PrefixStatesParameters(root constructs.IConstruct, prefix *string) error {
	return nil
}

func validateNewChoiceParameters(scope constructs.Construct, id *string, props *ChoiceProps) error {
	return nil
}

