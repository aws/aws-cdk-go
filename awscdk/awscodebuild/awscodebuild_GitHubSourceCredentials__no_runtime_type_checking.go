//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awscodebuild

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GitHubSourceCredentials) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (g *jsiiProxy_GitHubSourceCredentials) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (g *jsiiProxy_GitHubSourceCredentials) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (g *jsiiProxy_GitHubSourceCredentials) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (g *jsiiProxy_GitHubSourceCredentials) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateGitHubSourceCredentials_IsConstructParameters(x interface{}) error {
	return nil
}

func validateGitHubSourceCredentials_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewGitHubSourceCredentialsParameters(scope constructs.Construct, id *string, props *GitHubSourceCredentialsProps) error {
	return nil
}

