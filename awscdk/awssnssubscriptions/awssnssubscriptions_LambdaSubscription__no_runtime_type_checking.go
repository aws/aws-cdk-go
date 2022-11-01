//go:build no_runtime_type_checking

package awssnssubscriptions

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LambdaSubscription) validateBindParameters(topic awssns.ITopic) error {
	return nil
}

func validateNewLambdaSubscriptionParameters(fn awslambda.IFunction, props *LambdaSubscriptionProps) error {
	return nil
}

