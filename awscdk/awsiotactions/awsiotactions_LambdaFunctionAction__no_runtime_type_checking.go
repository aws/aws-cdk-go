//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsiotactions

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LambdaFunctionAction) validateBindParameters(topicRule awsiot.ITopicRule) error {
	return nil
}

func validateNewLambdaFunctionActionParameters(func_ awslambda.IFunction) error {
	return nil
}

