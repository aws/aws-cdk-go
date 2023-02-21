package awsssm


// The `MaintenanceWindowLambdaParameters` property type specifies the parameters for a `LAMBDA` task type for a maintenance window task in AWS Systems Manager .
//
// `MaintenanceWindowLambdaParameters` is a property of the [TaskInvocationParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   maintenanceWindowLambdaParametersProperty := &maintenanceWindowLambdaParametersProperty{
//   	clientContext: jsii.String("clientContext"),
//   	payload: jsii.String("payload"),
//   	qualifier: jsii.String("qualifier"),
//   }
//
type CfnMaintenanceWindowTask_MaintenanceWindowLambdaParametersProperty struct {
	// Client-specific information to pass to the AWS Lambda function that you're invoking.
	//
	// You can then use the `context` variable to process the client information in your AWS Lambda function.
	ClientContext *string `field:"optional" json:"clientContext" yaml:"clientContext"`
	// JSON to provide to your AWS Lambda function as input.
	//
	// > Although `Type` is listed as "String" for this property, the payload content must be formatted as a Base64-encoded binary data object.
	//
	// *Length Constraint:* 4096.
	Payload *string `field:"optional" json:"payload" yaml:"payload"`
	// An AWS Lambda function version or alias name.
	//
	// If you specify a function version, the action uses the qualified function Amazon Resource Name (ARN) to invoke a specific Lambda function. If you specify an alias name, the action uses the alias ARN to invoke the Lambda function version that the alias points to.
	Qualifier *string `field:"optional" json:"qualifier" yaml:"qualifier"`
}

