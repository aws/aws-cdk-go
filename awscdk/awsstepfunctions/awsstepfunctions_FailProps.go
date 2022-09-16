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
//   submitJob := tasks.NewLambdaInvoke(this, jsii.String("Submit Job"), &lambdaInvokeProps{
//   	lambdaFunction: submitLambda,
//   	// Lambda's result is in the attribute `Payload`
//   	outputPath: jsii.String("$.Payload"),
//   })
//
//   waitX := sfn.NewWait(this, jsii.String("Wait X Seconds"), &waitProps{
//   	time: sfn.waitTime.secondsPath(jsii.String("$.waitSeconds")),
//   })
//
//   getStatus := tasks.NewLambdaInvoke(this, jsii.String("Get Job Status"), &lambdaInvokeProps{
//   	lambdaFunction: getStatusLambda,
//   	// Pass just the field named "guid" into the Lambda, put the
//   	// Lambda's result in a field called "status" in the response
//   	inputPath: jsii.String("$.guid"),
//   	outputPath: jsii.String("$.Payload"),
//   })
//
//   jobFailed := sfn.NewFail(this, jsii.String("Job Failed"), &failProps{
//   	cause: jsii.String("AWS Batch Job Failed"),
//   	error: jsii.String("DescribeJob returned FAILED"),
//   })
//
//   finalStatus := tasks.NewLambdaInvoke(this, jsii.String("Get Final Job Status"), &lambdaInvokeProps{
//   	lambdaFunction: getStatusLambda,
//   	// Use "guid" field as input
//   	inputPath: jsii.String("$.guid"),
//   	outputPath: jsii.String("$.Payload"),
//   })
//
//   definition := submitJob.next(waitX).next(getStatus).next(sfn.NewChoice(this, jsii.String("Job Complete?")).when(sfn.condition.stringEquals(jsii.String("$.status"), jsii.String("FAILED")), jobFailed).when(sfn.condition.stringEquals(jsii.String("$.status"), jsii.String("SUCCEEDED")), finalStatus).otherwise(waitX))
//
//   sfn.NewStateMachine(this, jsii.String("StateMachine"), &stateMachineProps{
//   	definition: definition,
//   	timeout: awscdk.Duration.minutes(jsii.Number(5)),
//   })
//
type FailProps struct {
	// A description for the cause of the failure.
	Cause *string `field:"optional" json:"cause" yaml:"cause"`
	// An optional description for this state.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Error code used to represent this failure.
	Error *string `field:"optional" json:"error" yaml:"error"`
}

