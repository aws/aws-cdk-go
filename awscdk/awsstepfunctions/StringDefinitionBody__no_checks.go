//go:build no_runtime_type_checking

package awsstepfunctions

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StringDefinitionBody) validateBindParameters(_scope constructs.Construct, _sfnPrincipal awsiam.IPrincipal, _sfnProps *StateMachineProps) error {
	return nil
}

func validateStringDefinitionBody_FromChainableParameters(chainable IChainable) error {
	return nil
}

func validateStringDefinitionBody_FromFileParameters(path *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateStringDefinitionBody_FromStringParameters(definition *string) error {
	return nil
}

func validateNewStringDefinitionBodyParameters(body *string) error {
	return nil
}

