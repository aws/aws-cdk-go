//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awscloudformation

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CustomResourceProvider) validateBindParameters(_arg awscdk.Construct) error {
	return nil
}

func validateCustomResourceProvider_FromLambdaParameters(handler awslambda.IFunction) error {
	return nil
}

func validateCustomResourceProvider_FromTopicParameters(topic awssns.ITopic) error {
	return nil
}

func validateCustomResourceProvider_LambdaParameters(handler awslambda.IFunction) error {
	return nil
}

func validateCustomResourceProvider_TopicParameters(topic awssns.ITopic) error {
	return nil
}

