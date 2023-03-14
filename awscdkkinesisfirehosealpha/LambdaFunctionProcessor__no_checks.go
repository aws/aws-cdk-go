//go:build no_runtime_type_checking

// The CDK Construct Library for AWS::KinesisFirehose
package awscdkkinesisfirehosealpha

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LambdaFunctionProcessor) validateBindParameters(_scope constructs.Construct, options *DataProcessorBindOptions) error {
	return nil
}

func validateNewLambdaFunctionProcessorParameters(lambdaFunction awslambda.IFunction, props *DataProcessorProps) error {
	return nil
}

