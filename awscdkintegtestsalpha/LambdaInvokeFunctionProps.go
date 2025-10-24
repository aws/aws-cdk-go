package awscdkintegtestsalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Options to pass to the Lambda invokeFunction API call.
//
// Example:
//   var lambdaFunction IFunction
//   var app App
//
//
//   stack := awscdk.NewStack(app, jsii.String("cdk-integ-lambda-bundling"))
//
//   integ := awscdkintegtestsalpha.NewIntegTest(app, jsii.String("IntegTest"), &IntegTestProps{
//   	TestCases: []Stack{
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
	// Default: InvocationType.REQUEST_RESPONSE
	//
	// Experimental.
	InvocationType InvocationType `field:"optional" json:"invocationType" yaml:"invocationType"`
	// How long, in days, the log contents will be retained.
	// Default: - no retention days specified.
	//
	// Experimental.
	LogRetention awslogs.RetentionDays `field:"optional" json:"logRetention" yaml:"logRetention"`
	// Whether to return the logs as part of the response.
	// Default: LogType.NONE
	//
	// Experimental.
	LogType LogType `field:"optional" json:"logType" yaml:"logType"`
	// Payload to send as part of the invoke.
	// Default: - no payload.
	//
	// Experimental.
	Payload *string `field:"optional" json:"payload" yaml:"payload"`
}

