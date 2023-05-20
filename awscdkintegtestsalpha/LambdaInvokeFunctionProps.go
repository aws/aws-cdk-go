package awscdkintegtestsalpha


// Options to pass to the Lambda invokeFunction API call.
//
// Example:
//   var lambdaFunction iFunction
//   var app app
//
//
//   stack := awscdk.NewStack(app, jsii.String("cdk-integ-lambda-bundling"))
//
//   integ := awscdkintegtestsalpha.NewIntegTest(app, jsii.String("IntegTest"), &IntegTestProps{
//   	TestCases: []stack{
//   		stack,
//   	},
//   })
//
//   invoke := integ.Assertions.InvokeFunction(&LambdaInvokeFunctionProps{
//   	FunctionName: lambdaFunction.FunctionName,
//   })
//   invoke.Expect(awscdkintegtestsalpha.ExpectedResult_ObjectLike(map[string]interface{}{
//   	"Payload": jsii.String("200"),
//   }))
//
// Experimental.
type LambdaInvokeFunctionProps struct {
	// The name of the function to invoke.
	// Experimental.
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
	// The type of invocation to use.
	// Experimental.
	InvocationType InvocationType `field:"optional" json:"invocationType" yaml:"invocationType"`
	// Whether to return the logs as part of the response.
	// Experimental.
	LogType LogType `field:"optional" json:"logType" yaml:"logType"`
	// Payload to send as part of the invoke.
	// Experimental.
	Payload *string `field:"optional" json:"payload" yaml:"payload"`
}

