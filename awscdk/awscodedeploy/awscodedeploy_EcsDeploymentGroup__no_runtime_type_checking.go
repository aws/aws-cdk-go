//go:build no_runtime_type_checking

package awscodedeploy

// Building without runtime type checking enabled, so all the below just return nil

func validateEcsDeploymentGroup_FromEcsDeploymentGroupAttributesParameters(scope constructs.Construct, id *string, attrs *EcsDeploymentGroupAttributes) error {
	return nil
}

