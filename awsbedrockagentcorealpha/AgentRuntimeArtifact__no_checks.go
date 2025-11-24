//go:build no_runtime_type_checking

package awsbedrockagentcorealpha

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AgentRuntimeArtifact) validateBindParameters(scope constructs.Construct, runtime Runtime) error {
	return nil
}

func validateAgentRuntimeArtifact_FromAssetParameters(directory *string, options *awsecrassets.DockerImageAssetOptions) error {
	return nil
}

func validateAgentRuntimeArtifact_FromEcrRepositoryParameters(repository awsecr.IRepository) error {
	return nil
}

func validateAgentRuntimeArtifact_FromS3Parameters(s3Location *awss3.Location, runtime AgentCoreRuntime, entrypoint *[]*string) error {
	return nil
}

