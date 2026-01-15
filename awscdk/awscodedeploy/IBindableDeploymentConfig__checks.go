//go:build !no_runtime_type_checking

package awscodedeploy

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscodedeploy"
)

func (i *jsiiProxy_IBindableDeploymentConfig) validateBindEnvironmentParameters(deploymentGroup interfacesawscodedeploy.IDeploymentGroupRef) error {
	if deploymentGroup == nil {
		return fmt.Errorf("parameter deploymentGroup is required, but nil was provided")
	}

	return nil
}

