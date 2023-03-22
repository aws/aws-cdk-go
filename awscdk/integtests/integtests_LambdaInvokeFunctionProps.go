package integtests


// Options to pass to the Lambda invokeFunction API call.
//
// Example:
//   var lambdaFunction iFunction
//   var app app
//
//
//   stack := awscdk.NewStack(app, jsii.String("cdk-integ-lambda-bundling"))
//
//   integ := awscdk.NewIntegTest(app, jsii.String("IntegTest"), &integTestProps{
//   	testCases: []stack{
//   		stack,
//   	},
//   })
//
//   invoke := integ.assertions.invokeFunction(&lambdaInvokeFunctionProps{
//   	functionName: lambdaFunction.functionName,
//   })
//   invoke.expect(awscdk.ExpectedResult.objectLike(map[string]interface{}{
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

