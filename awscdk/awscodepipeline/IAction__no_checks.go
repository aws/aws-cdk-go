//go:build no_runtime_type_checking

package awscodepipeline

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IAction) validateBindParameters(scope constructs.Construct, stage IStage, options *ActionBindOptions) error {
	return nil
}

func (i *jsiiProxy_IAction) validateOnStateChangeParameters(name *string, options *awsevents.RuleProps) error {
	return nil
}

