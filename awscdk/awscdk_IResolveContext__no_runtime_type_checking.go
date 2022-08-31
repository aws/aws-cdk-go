//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// An experiment to bundle the entire CDK into a single module
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IResolveContext) validateRegisterPostProcessorParameters(postProcessor IPostProcessor) error {
	return nil
}

func (i *jsiiProxy_IResolveContext) validateResolveParameters(x interface{}, options *ResolveChangeContextOptions) error {
	return nil
}

