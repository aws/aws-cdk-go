//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package pipelines

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DeployCdkStackAction) validateBindParameters(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (d *jsiiProxy_DeployCdkStackAction) validateOnStateChangeParameters(name *string, options *awsevents.RuleProps) error {
	return nil
}

func validateDeployCdkStackAction_FromStackArtifactParameters(scope constructs.Construct, artifact cxapi.CloudFormationStackArtifact, options *CdkStackActionFromArtifactOptions) error {
	return nil
}

func validateNewDeployCdkStackActionParameters(props *DeployCdkStackActionProps) error {
	return nil
}

