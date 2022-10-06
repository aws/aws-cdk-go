//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awscodedeploy

// Building without runtime type checking enabled, so all the below just return nil

func validateLambdaDeploymentConfig_ImportParameters(_scope constructs.Construct, _id *string, props *LambdaDeploymentConfigImportProps) error {
	return nil
}

