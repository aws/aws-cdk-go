package awseventsv2


// Parameters for starting an AWS Step Functions state machine execution target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   stepFunctionsParametersProperty := &StepFunctionsParametersProperty{
//   	InvocationTimeoutSeconds: jsii.String("invocationTimeoutSeconds"),
//   	InvocationType: jsii.String("invocationType"),
//   	Name: jsii.String("name"),
//   	TraceHeader: jsii.String("traceHeader"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-stepfunctionsparameters.html
//
type CfnSubscriberPropsMixin_StepFunctionsParametersProperty struct {
	// The timeout in seconds for each invocation of the target, written as a string.
	//
	// Accepts a literal value or a JSONata expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-stepfunctionsparameters.html#cfn-eventsv2-subscriber-stepfunctionsparameters-invocationtimeoutseconds
	//
	InvocationTimeoutSeconds *string `field:"optional" json:"invocationTimeoutSeconds" yaml:"invocationTimeoutSeconds"`
	// How the execution is started: EVENT (StartExecution, asynchronous) or REQUEST_RESPONSE (StartSyncExecution, synchronous).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-stepfunctionsparameters.html#cfn-eventsv2-subscriber-stepfunctionsparameters-invocationtype
	//
	InvocationType *string `field:"optional" json:"invocationType" yaml:"invocationType"`
	// A name for the execution.
	//
	// Must be unique for the account, Region, and state machine. Accepts a literal value or a JSONata expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-stepfunctionsparameters.html#cfn-eventsv2-subscriber-stepfunctionsparameters-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The AWS X-Ray trace header for distributed tracing.
	//
	// Accepts a literal value or a JSONata expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-stepfunctionsparameters.html#cfn-eventsv2-subscriber-stepfunctionsparameters-traceheader
	//
	TraceHeader *string `field:"optional" json:"traceHeader" yaml:"traceHeader"`
}

