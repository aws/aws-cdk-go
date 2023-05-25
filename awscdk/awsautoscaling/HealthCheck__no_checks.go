//go:build no_runtime_type_checking

package awsautoscaling

// Building without runtime type checking enabled, so all the below just return nil

func validateHealthCheck_Ec2Parameters(options *Ec2HealthCheckOptions) error {
	return nil
}

func validateHealthCheck_ElbParameters(options *ElbHealthCheckOptions) error {
	return nil
}

