package awslogs

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Retry options for all AWS API calls.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//
//   logRetentionRetryOptions := &logRetentionRetryOptions{
//   	base: duration,
//   	maxRetries: jsii.Number(123),
//   }
//
// Experimental.
type LogRetentionRetryOptions struct {
	// The base duration to use in the exponential backoff for operation retries.
	// Experimental.
	Base awscdk.Duration `field:"optional" json:"base" yaml:"base"`
	// The maximum amount of retries.
	// Experimental.
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
}

