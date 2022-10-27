//go:build !no_runtime_type_checking

package pipelines

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/cxapi"
)

func (s *jsiiProxy_StackDeployment) validateAddStackDependencyParameters(stackDeployment StackDeployment) error {
	if stackDeployment == nil {
		return fmt.Errorf("parameter stackDeployment is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_StackDeployment) validateAddStackStepsParameters(pre *[]Step, changeSet *[]Step, post *[]Step) error {
	if pre == nil {
		return fmt.Errorf("parameter pre is required, but nil was provided")
	}

	if changeSet == nil {
		return fmt.Errorf("parameter changeSet is required, but nil was provided")
	}

	if post == nil {
		return fmt.Errorf("parameter post is required, but nil was provided")
	}

	return nil
}

func validateStackDeployment_FromArtifactParameters(stackArtifact cxapi.CloudFormationStackArtifact) error {
	if stackArtifact == nil {
		return fmt.Errorf("parameter stackArtifact is required, but nil was provided")
	}

	return nil
}

