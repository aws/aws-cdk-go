//go:build no_runtime_type_checking

package awsstepfunctions

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DistributedMap) validateAddBranchParameters(branch StateGraph) error {
	return nil
}

func (d *jsiiProxy_DistributedMap) validateAddCatchParameters(handler IChainable, props *CatchProps) error {
	return nil
}

func (d *jsiiProxy_DistributedMap) validateAddChoiceParameters(condition Condition, next State, options *ChoiceTransitionOptions) error {
	return nil
}

func (d *jsiiProxy_DistributedMap) validateAddItemProcessorParameters(processor StateGraph, config *ProcessorConfig) error {
	return nil
}

func (d *jsiiProxy_DistributedMap) validateAddIteratorParameters(iteration StateGraph) error {
	return nil
}

func (d *jsiiProxy_DistributedMap) validateAddPrefixParameters(x *string) error {
	return nil
}

func (d *jsiiProxy_DistributedMap) validateAddRetryParameters(props *RetryProps) error {
	return nil
}

func (d *jsiiProxy_DistributedMap) validateBindToGraphParameters(graph StateGraph) error {
	return nil
}

func (d *jsiiProxy_DistributedMap) validateItemProcessorParameters(processor IChainable, config *ProcessorConfig) error {
	return nil
}

func (d *jsiiProxy_DistributedMap) validateMakeDefaultParameters(def State) error {
	return nil
}

func (d *jsiiProxy_DistributedMap) validateMakeNextParameters(next State) error {
	return nil
}

func (d *jsiiProxy_DistributedMap) validateNextParameters(next IChainable) error {
	return nil
}

func (d *jsiiProxy_DistributedMap) validateWhenBoundToGraphParameters(graph StateGraph) error {
	return nil
}

func validateDistributedMap_FilterNextablesParameters(states *[]State) error {
	return nil
}

func validateDistributedMap_FindReachableEndStatesParameters(start State, options *FindStateOptions) error {
	return nil
}

func validateDistributedMap_FindReachableStatesParameters(start State, options *FindStateOptions) error {
	return nil
}

func validateDistributedMap_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDistributedMap_IsDistributedMapParameters(x interface{}) error {
	return nil
}

func validateDistributedMap_PrefixStatesParameters(root constructs.IConstruct, prefix *string) error {
	return nil
}

func (j *jsiiProxy_DistributedMap) validateSetProcessorConfigParameters(val *ProcessorConfig) error {
	return nil
}

func validateNewDistributedMapParameters(scope constructs.Construct, id *string, props *DistributedMapProps) error {
	return nil
}

