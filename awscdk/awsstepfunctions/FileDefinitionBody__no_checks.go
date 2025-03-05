//go:build no_runtime_type_checking

package awsstepfunctions

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FileDefinitionBody) validateBindParameters(scope constructs.Construct, _sfnPrincipal awsiam.IPrincipal, _sfnProps *StateMachineProps) error {
	return nil
}

func validateFileDefinitionBody_FromChainableParameters(chainable IChainable) error {
	return nil
}

func validateFileDefinitionBody_FromFileParameters(path *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateFileDefinitionBody_FromStringParameters(definition *string) error {
	return nil
}

func validateNewFileDefinitionBodyParameters(path *string, options *awss3assets.AssetOptions) error {
	return nil
}

