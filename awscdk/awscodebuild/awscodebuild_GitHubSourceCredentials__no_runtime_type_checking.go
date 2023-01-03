//go:build no_runtime_type_checking

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

func validateGitHubSourceCredentials_IsConstructParameters(x interface{}) error {
	return nil
}

func validateGitHubSourceCredentials_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateGitHubSourceCredentials_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewGitHubSourceCredentialsParameters(scope constructs.Construct, id *string, props *GitHubSourceCredentialsProps) error {
	return nil
}

