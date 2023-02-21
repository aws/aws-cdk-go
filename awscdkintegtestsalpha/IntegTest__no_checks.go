//go:build no_runtime_type_checking

// CDK Integration Testing Constructs
package awscdkintegtestsalpha

// Building without runtime type checking enabled, so all the below just return nil

func validateIntegTest_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewIntegTestParameters(scope constructs.Construct, id *string, props *IntegTestProps) error {
	return nil
}

