package awseventsv2


// Parameters for invoking an AWS Lambda function target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaParametersProperty := &LambdaParametersProperty{
//   	DurableExecutionName: jsii.String("durableExecutionName"),
//   	InvocationTimeoutSeconds: jsii.String("invocationTimeoutSeconds"),
//   	InvocationType: jsii.String("invocationType"),
//   	Qualifier: jsii.String("qualifier"),
//   	TenantId: jsii.String("tenantId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-lambdaparameters.html
//
type CfnSubscriber_LambdaParametersProperty struct {
	// A unique name for a durable function execution.
	//
	// Accepts a literal value or a JSONata expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-lambdaparameters.html#cfn-eventsv2-subscriber-lambdaparameters-durableexecutionname
	//
	DurableExecutionName *string `field:"optional" json:"durableExecutionName" yaml:"durableExecutionName"`
	// The timeout in seconds for each invocation of the target, written as a string.
	//
	// Accepts a literal value or a JSONata expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-lambdaparameters.html#cfn-eventsv2-subscriber-lambdaparameters-invocationtimeoutseconds
	//
	InvocationTimeoutSeconds *string `field:"optional" json:"invocationTimeoutSeconds" yaml:"invocationTimeoutSeconds"`
	// How the function is invoked: EVENT (asynchronous) or REQUEST_RESPONSE (synchronous).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-lambdaparameters.html#cfn-eventsv2-subscriber-lambdaparameters-invocationtype
	//
	InvocationType *string `field:"optional" json:"invocationType" yaml:"invocationType"`
	// The version or alias of the Lambda function to invoke.
	//
	// Accepts a literal value or a JSONata expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-lambdaparameters.html#cfn-eventsv2-subscriber-lambdaparameters-qualifier
	//
	Qualifier *string `field:"optional" json:"qualifier" yaml:"qualifier"`
	// The tenant identifier for multi-tenant Lambda functions.
	//
	// Accepts a literal value or a JSONata expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-lambdaparameters.html#cfn-eventsv2-subscriber-lambdaparameters-tenantid
	//
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

