//go:build no_runtime_type_checking

package awslambdamixins

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FunctionRuntimeConfig) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (f *jsiiProxy_FunctionRuntimeConfig) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateFunctionRuntimeConfig_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewFunctionRuntimeConfigParameters(props *FunctionRuntimeConfigProps) error {
	return nil
}

