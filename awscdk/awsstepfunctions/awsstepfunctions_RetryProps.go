package awsstepfunctions

import (
	"github.com/aws/aws-cdk-go/awscdk"
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
//   parallel.branch(shipItem)
//   parallel.branch(sendInvoice)
//   parallel.branch(restock)
//
//   // Retry the whole workflow if something goes wrong
//   parallel.addRetry(&retryProps{
//   	maxAttempts: jsii.Number(1),
//   })
//
//   // How to recover from errors
//   sendFailureNotification := sfn.NewPass(this, jsii.String("SendFailureNotification"))
//   parallel.addCatch(sendFailureNotification)
//
//   // What to do in case everything succeeded
//   closeOrder := sfn.NewPass(this, jsii.String("CloseOrder"))
//   parallel.next(closeOrder)
//
// Experimental.
type RetryProps struct {
	// Multiplication for how much longer the wait interval gets on every retry.
	// Experimental.
	BackoffRate *float64 `field:"optional" json:"backoffRate" yaml:"backoffRate"`
	// Errors to retry.
	//
	// A list of error strings to retry, which can be either predefined errors
	// (for example Errors.NoChoiceMatched) or a self-defined error.
	// Experimental.
	Errors *[]*string `field:"optional" json:"errors" yaml:"errors"`
	// How many seconds to wait initially before retrying.
	// Experimental.
	Interval awscdk.Duration `field:"optional" json:"interval" yaml:"interval"`
	// How many times to retry this particular error.
	//
	// May be 0 to disable retry for specific errors (in case you have
	// a catch-all retry policy).
	// Experimental.
	MaxAttempts *float64 `field:"optional" json:"maxAttempts" yaml:"maxAttempts"`
}

