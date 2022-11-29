//go:build no_runtime_type_checking

package awscodestar

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GitHubRepository) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (g *jsiiProxy_GitHubRepository) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (g *jsiiProxy_GitHubRepository) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (g *jsiiProxy_GitHubRepository) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (g *jsiiProxy_GitHubRepository) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateGitHubRepository_IsConstructParameters(x interface{}) error {
	return nil
}

func validateGitHubRepository_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewGitHubRepositoryParameters(scope constructs.Construct, id *string, props *GitHubRepositoryProps) error {
	return nil
}

