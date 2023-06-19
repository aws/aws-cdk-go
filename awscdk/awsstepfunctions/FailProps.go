package awsstepfunctions


// Properties for defining a Fail state.
//
// Example:
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//
//   var submitLambda function
//   var getStatusLambda function
//
//
//   submitJob := tasks.NewLambdaInvoke(this, jsii.String("Submit Job"), &LambdaInvokeProps{
//   	LambdaFunction: submitLambda,
//   	// Lambda's result is in the attribute `Payload`
//   	OutputPath: jsii.String("$.Payload"),
//   })
//
//   waitX := sfn.NewWait(this, jsii.String("Wait X Seconds"), &WaitProps{
//   	Time: sfn.WaitTime_SecondsPath(jsii.String("$.waitSeconds")),
//   })
//
//   getStatus := tasks.NewLambdaInvoke(this, jsii.String("Get Job Status"), &LambdaInvokeProps{
//   	LambdaFunction: getStatusLambda,
//   	// Pass just the field named "guid" into the Lambda, put the
//   	// Lambda's result in a field called "status" in the response
//   	InputPath: jsii.String("$.guid"),
//   	OutputPath: jsii.String("$.Payload"),
//   })
//
//   jobFailed := sfn.NewFail(this, jsii.String("Job Failed"), &FailProps{
//   	Cause: jsii.String("AWS Batch Job Failed"),
//   	Error: jsii.String("DescribeJob returned FAILED"),
//   })
//
//   finalStatus := tasks.NewLambdaInvoke(this, jsii.String("Get Final Job Status"), &LambdaInvokeProps{
//   	LambdaFunction: getStatusLambda,
//   	// Use "guid" field as input
//   	InputPath: jsii.String("$.guid"),
//   	OutputPath: jsii.String("$.Payload"),
//   })
//
//   definition := submitJob.Next(waitX).Next(getStatus).Next(sfn.NewChoice(this, jsii.String("Job Complete?")).When(sfn.Condition_StringEquals(jsii.String("$.status"), jsii.String("FAILED")), jobFailed).When(sfn.Condition_StringEquals(jsii.String("$.status"), jsii.String("SUCCEEDED")), finalStatus).Otherwise(waitX))
//
//   sfn.NewStateMachine(this, jsii.String("StateMachine"), &StateMachineProps{
//   	Definition: Definition,
//   	Timeout: awscdk.Duration_Minutes(jsii.Number(5)),
//   })
//
// Experimental.
type FailProps struct {
	// A description for the cause of the failure.
	// Experimental.
	Cause *string `field:"optional" json:"cause" yaml:"cause"`
	// An optional description for this state.
	// Experimental.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Error code used to represent this failure.
	// Experimental.
	Error *string `field:"optional" json:"error" yaml:"error"`
}

