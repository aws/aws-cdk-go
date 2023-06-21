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
//   // Retry the whole workflow if something goes wrong
//   parallel.AddRetry(&RetryProps{
//   	MaxAttempts: jsii.Number(1),
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
	BackoffRate *float64 `field:"optional" json:"backoffRate" yaml:"backoffRate"`
	// Errors to retry.
	//
	// A list of error strings to retry, which can be either predefined errors
	// (for example Errors.NoChoiceMatched) or a self-defined error.
	Errors *[]*string `field:"optional" json:"errors" yaml:"errors"`
	// How many seconds to wait initially before retrying.
	Interval awscdk.Duration `field:"optional" json:"interval" yaml:"interval"`
	// How many times to retry this particular error.
	//
	// May be 0 to disable retry for specific errors (in case you have
	// a catch-all retry policy).
	MaxAttempts *float64 `field:"optional" json:"maxAttempts" yaml:"maxAttempts"`
}

