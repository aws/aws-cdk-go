// CDK Integration Testing Constructs
package awscdkintegtestsalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props for creating a WaiterStateMachine.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import integ_tests_alpha "github.com/aws/aws-cdk-go/awscdkintegtestsalpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   waiterStateMachineProps := &waiterStateMachineProps{
//   	backoffRate: jsii.Number(123),
//   	interval: cdk.duration.minutes(jsii.Number(30)),
//   	totalTimeout: cdk.*duration.minutes(jsii.Number(30)),
//   }
//
// Experimental.
type WaiterStateMachineProps struct {
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

