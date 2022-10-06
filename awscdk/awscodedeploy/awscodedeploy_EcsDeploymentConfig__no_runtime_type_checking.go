//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awscodedeploy

// Building without runtime type checking enabled, so all the below just return nil

func validateEcsDeploymentConfig_FromEcsDeploymentConfigNameParameters(_scope constructs.Construct, _id *string, ecsDeploymentConfigName *string) error {
	return nil
}

