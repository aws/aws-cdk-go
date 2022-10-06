//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package pipelines

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CdkStage) validateAddApplicationParameters(appStage awscdk.Stage, options *AddStageOptions) error {
	return nil
}

func (c *jsiiProxy_CdkStage) validateAddManualApprovalActionParameters(options *AddManualApprovalOptions) error {
	return nil
}

func (c *jsiiProxy_CdkStage) validateAddStackArtifactDeploymentParameters(stackArtifact cxapi.CloudFormationStackArtifact, options *AddStackOptions) error {
	return nil
}

func (c *jsiiProxy_CdkStage) validateDeploysStackParameters(artifactId *string) error {
	return nil
}

func (c *jsiiProxy_CdkStage) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (c *jsiiProxy_CdkStage) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateCdkStage_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewCdkStageParameters(scope constructs.Construct, id *string, props *CdkStageProps) error {
	return nil
}

