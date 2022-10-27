//go:build no_runtime_type_checking

package pipelines

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StackDeployment) validateAddStackDependencyParameters(stackDeployment StackDeployment) error {
	return nil
}

func (s *jsiiProxy_StackDeployment) validateAddStackStepsParameters(pre *[]Step, changeSet *[]Step, post *[]Step) error {
	return nil
}

func validateStackDeployment_FromArtifactParameters(stackArtifact cxapi.CloudFormationStackArtifact) error {
	return nil
}

