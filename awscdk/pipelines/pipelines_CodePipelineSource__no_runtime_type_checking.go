//go:build no_runtime_type_checking

package pipelines

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CodePipelineSource) validateAddDependencyFileSetParameters(fs FileSet) error {
	return nil
}

func (c *jsiiProxy_CodePipelineSource) validateAddStepDependencyParameters(step Step) error {
	return nil
}

func (c *jsiiProxy_CodePipelineSource) validateConfigurePrimaryOutputParameters(fs FileSet) error {
	return nil
}

func (c *jsiiProxy_CodePipelineSource) validateDiscoverReferencedOutputsParameters(structure interface{}) error {
	return nil
}

func (c *jsiiProxy_CodePipelineSource) validateGetActionParameters(output awscodepipeline.Artifact, actionName *string, runOrder *float64) error {
	return nil
}

func (c *jsiiProxy_CodePipelineSource) validateProduceActionParameters(stage awscodepipeline.IStage, options *ProduceActionOptions) error {
	return nil
}

func (c *jsiiProxy_CodePipelineSource) validateSourceAttributeParameters(name *string) error {
	return nil
}

func validateCodePipelineSource_CodeCommitParameters(repository awscodecommit.IRepository, branch *string, props *CodeCommitSourceOptions) error {
	return nil
}

func validateCodePipelineSource_ConnectionParameters(repoString *string, branch *string, props *ConnectionSourceOptions) error {
	return nil
}

func validateCodePipelineSource_EcrParameters(repository awsecr.IRepository, props *ECRSourceOptions) error {
	return nil
}

func validateCodePipelineSource_GitHubParameters(repoString *string, branch *string, props *GitHubSourceOptions) error {
	return nil
}

func validateCodePipelineSource_S3Parameters(bucket awss3.IBucket, objectKey *string, props *S3SourceOptions) error {
	return nil
}

func validateCodePipelineSource_SequenceParameters(steps *[]Step) error {
	return nil
}

func validateNewCodePipelineSourceParameters(id *string) error {
	return nil
}

