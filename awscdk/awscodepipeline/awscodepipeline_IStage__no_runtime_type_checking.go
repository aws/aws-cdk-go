//go:build no_runtime_type_checking

package awscodepipeline

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IStage) validateAddActionParameters(action IAction) error {
	return nil
}

func (i *jsiiProxy_IStage) validateOnStateChangeParameters(name *string, options *awsevents.RuleProps) error {
	return nil
}

