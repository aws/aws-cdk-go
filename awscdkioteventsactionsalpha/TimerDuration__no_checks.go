//go:build no_runtime_type_checking

package awscdkioteventsactionsalpha

// Building without runtime type checking enabled, so all the below just return nil

func validateTimerDuration_FromDurationParameters(duration awscdk.Duration) error {
	return nil
}

func validateTimerDuration_FromExpressionParameters(expression awscdkioteventsalpha.Expression) error {
	return nil
}

