package awscdkpipestargetsalpha


// InvocationType for invoking the Lambda Function.
//
// Example:
//   var sourceQueue queue
//   var targetFunction iFunction
//
//
//   pipeTarget := targets.NewLambdaFunction(targetFunction, &LambdaFunctionParameters{
//   	InvocationType: targets.LambdaFunctionInvocationType_FIRE_AND_FORGET,
//   })
//
//   pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
//   	Source: NewSomeSource(sourceQueue),
//   	Target: pipeTarget,
//   })
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetlambdafunctionparameters.html
//
// Experimental.
type LambdaFunctionInvocationType string

const (
	// Invoke Lambda Function asynchronously (`Invoke`).
	//
	// `InvocationType` is set to `Event` on `Invoke`, see https://docs.aws.amazon.com/lambda/latest/api/API_Invoke.html for more details.
	// Experimental.
	LambdaFunctionInvocationType_FIRE_AND_FORGET LambdaFunctionInvocationType = "FIRE_AND_FORGET"
	// Invoke Lambda Function synchronously (`Invoke`) and wait for the response.
	//
	// `InvocationType` is set to `RequestResponse` on `Invoke`, see https://docs.aws.amazon.com/lambda/latest/api/API_Invoke.html for more details.
	// Experimental.
	LambdaFunctionInvocationType_REQUEST_RESPONSE LambdaFunctionInvocationType = "REQUEST_RESPONSE"
)

