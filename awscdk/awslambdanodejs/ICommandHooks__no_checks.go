//go:build no_runtime_type_checking

package awslambdanodejs

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ICommandHooks) validateAfterBundlingParameters(inputDir *string, outputDir *string) error {
	return nil
}

func (i *jsiiProxy_ICommandHooks) validateBeforeBundlingParameters(inputDir *string, outputDir *string) error {
	return nil
}

func (i *jsiiProxy_ICommandHooks) validateBeforeInstallParameters(inputDir *string, outputDir *string) error {
	return nil
}

