package awscdkintegtestsalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Props for creating a WaiterStateMachine.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import integ_tests_alpha "github.com/aws/aws-cdk-go/awscdkintegtestsalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   waiterStateMachineProps := &WaiterStateMachineProps{
//   	BackoffRate: jsii.Number(123),
//   	Interval: cdk.Duration_Minutes(jsii.Number(30)),
//   	ProviderLogLevel: awscdk.Aws_lambda.ApplicationLogLevel_INFO,
//   	TotalTimeout: cdk.Duration_*Minutes(jsii.Number(30)),
//   }
//
// Experimental.
type WaiterStateMachineProps struct {
	// The log level of the provider lambda function.
	// Default: ApplicationLogLevel.FATAL
	//
	// Experimental.
	ProviderLogLevel awslambda.ApplicationLogLevel `field:"optional" json:"providerLogLevel" yaml:"providerLogLevel"`
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

