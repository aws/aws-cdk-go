//go:build no_runtime_type_checking

package awscodedeploy

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TrafficRouting) validateBindParameters(scope constructs.Construct) error {
	return nil
}

func validateTrafficRouting_TimeBasedCanaryParameters(props *TimeBasedCanaryTrafficRoutingProps) error {
	return nil
}

func validateTrafficRouting_TimeBasedLinearParameters(props *TimeBasedLinearTrafficRoutingProps) error {
	return nil
}

