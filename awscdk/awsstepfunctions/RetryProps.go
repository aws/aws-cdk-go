package awsstepfunctions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Retry details.
//
// Example:
//   parallel := sfn.NewParallel(this, jsii.String("Do the work in parallel"))
//
//   // Add branches to be executed in parallel
//   shipItem := sfn.NewPass(this, jsii.String("ShipItem"))
//   sendInvoice := sfn.NewPass(this, jsii.String("SendInvoice"))
//   restock := sfn.NewPass(this, jsii.String("Restock"))
//   parallel.Branch(shipItem)
//   parallel.Branch(sendInvoice)
//   parallel.Branch(restock)
//
//   // Retry the whole workflow if something goes wrong with exponential backoff
//   parallel.AddRetry(&RetryProps{
//   	MaxAttempts: jsii.Number(1),
//   	MaxDelay: awscdk.Duration_Seconds(jsii.Number(5)),
//   	JitterStrategy: sfn.JitterType_FULL,
//   })
//
//   // How to recover from errors
//   sendFailureNotification := sfn.NewPass(this, jsii.String("SendFailureNotification"))
//   parallel.AddCatch(sendFailureNotification)
//
//   // What to do in case everything succeeded
//   closeOrder := sfn.NewPass(this, jsii.String("CloseOrder"))
//   parallel.Next(closeOrder)
//
type RetryProps struct {
	// Multiplication for how much longer the wait interval gets on every retry.
	// Default: 2.
	//
	BackoffRate *float64 `field:"optional" json:"backoffRate" yaml:"backoffRate"`
	// Errors to retry.
	//
	// A list of error strings to retry, which can be either predefined errors
	// (for example Errors.NoChoiceMatched) or a self-defined error.
	// Default: All errors.
	//
	Errors *[]*string `field:"optional" json:"errors" yaml:"errors"`
	// How many seconds to wait initially before retrying.
	// Default: Duration.seconds(1)
	//
	Interval awscdk.Duration `field:"optional" json:"interval" yaml:"interval"`
	// Introduces a randomization over the retry interval.
	// Default: - No jitter strategy.
	//
	JitterStrategy JitterType `field:"optional" json:"jitterStrategy" yaml:"jitterStrategy"`
	// How many times to retry this particular error.
	//
	// May be 0 to disable retry for specific errors (in case you have
	// a catch-all retry policy).
	// Default: 3.
	//
	MaxAttempts *float64 `field:"optional" json:"maxAttempts" yaml:"maxAttempts"`
	// Maximum limit on retry interval growth during exponential backoff.
	// Default: - No max delay.
	//
	MaxDelay awscdk.Duration `field:"optional" json:"maxDelay" yaml:"maxDelay"`
}

