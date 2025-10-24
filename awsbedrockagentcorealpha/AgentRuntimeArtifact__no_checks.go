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

