//go:build no_runtime_type_checking

package awsstepfunctions

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DefinitionBody) validateBindParameters(scope constructs.Construct, sfnPrincipal awsiam.IPrincipal, sfnProps *StateMachineProps) error {
	return nil
}

func validateDefinitionBody_FromChainableParameters(chainable IChainable) error {
	return nil
}

func validateDefinitionBody_FromFileParameters(path *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateDefinitionBody_FromStringParameters(definition *string) error {
	return nil
}

