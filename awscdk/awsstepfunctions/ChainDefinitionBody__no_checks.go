//go:build no_runtime_type_checking

package awsstepfunctions

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ChainDefinitionBody) validateBindParameters(scope constructs.Construct, _sfnPrincipal awsiam.IPrincipal, sfnProps *StateMachineProps) error {
	return nil
}

func validateChainDefinitionBody_FromChainableParameters(chainable IChainable) error {
	return nil
}

func validateChainDefinitionBody_FromFileParameters(path *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateChainDefinitionBody_FromStringParameters(definition *string) error {
	return nil
}

func validateNewChainDefinitionBodyParameters(chainable IChainable) error {
	return nil
}

