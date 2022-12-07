// CDK Integration Testing Constructs
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
//   describe := testCase.assertions.awsApiCall(jsii.String("StepFunctions"), jsii.String("describeExecution"), map[string]*string{
//   	"executionArn": start.getAttString(jsii.String("executionArn")),
//   }).expect(awscdkintegtestsalpha.ExpectedResult.objectLike(map[string]interface{}{
//   	"status": jsii.String("SUCCEEDED"),
//   })).waitForAssertions(&waiterStateMachineOptions{
//   	totalTimeout: awscdk.Duration.minutes(jsii.Number(5)),
//   	interval: awscdk.Duration.seconds(jsii.Number(15)),
//   	backoffRate: jsii.Number(3),
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
	// Experimental.
	BackoffRate *float64 `field:"optional" json:"backoffRate" yaml:"backoffRate"`
	// The interval (number of seconds) to wait between attempts.
	// Experimental.
	Interval awscdk.Duration `field:"optional" json:"interval" yaml:"interval"`
	// The total time that the state machine will wait for a successful response.
	// Experimental.
	TotalTimeout awscdk.Duration `field:"optional" json:"totalTimeout" yaml:"totalTimeout"`
}

