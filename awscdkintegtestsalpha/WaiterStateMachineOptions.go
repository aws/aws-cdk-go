package awscdkintegtestsalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Options for creating a WaiterStateMachine.
//
// Example:
//   var testCase integTest
//   var start iApiCall
//
//
//   describe := testCase.Assertions.AwsApiCall(jsii.String("StepFunctions"), jsii.String("describeExecution"), map[string]*string{
//   	"executionArn": start.getAttString(jsii.String("executionArn")),
//   }).Expect(awscdkintegtestsalpha.ExpectedResult_ObjectLike(map[string]interface{}{
//   	"status": jsii.String("SUCCEEDED"),
//   })).WaitForAssertions(&WaiterStateMachineOptions{
//   	TotalTimeout: awscdk.Duration_Minutes(jsii.Number(5)),
//   	Interval: awscdk.Duration_Seconds(jsii.Number(15)),
//   	BackoffRate: jsii.Number(3),
//   })
//
// Experimental.
type WaiterStateMachineOptions struct {
	// Backoff between attempts.
	//
	// This is the multiplier by which the retry interval increases
	// after each retry attempt.
	//
	// By default there is no backoff. Each retry will wait the amount of time
	// specified by `interval`.
	// Default: 1 (no backoff).
	//
	// Experimental.
	BackoffRate *float64 `field:"optional" json:"backoffRate" yaml:"backoffRate"`
	// The interval (number of seconds) to wait between attempts.
	// Default: Duration.seconds(5)
	//
	// Experimental.
	Interval awscdk.Duration `field:"optional" json:"interval" yaml:"interval"`
	// The total time that the state machine will wait for a successful response.
	// Default: Duration.minutes(30)
	//
	// Experimental.
	TotalTimeout awscdk.Duration `field:"optional" json:"totalTimeout" yaml:"totalTimeout"`
}

