package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Retry options for all AWS API calls.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logRetentionRetryOptions := &LogRetentionRetryOptions{
//   	Base: cdk.Duration_Minutes(jsii.Number(30)),
//   	MaxRetries: jsii.Number(123),
//   }
//
type LogRetentionRetryOptions struct {
	// The base duration to use in the exponential backoff for operation retries.
	// Default: - none, not used anymore.
	//
	// Deprecated: Unused since the upgrade to AWS SDK v3, which uses a different retry strategy.
	Base awscdk.Duration `field:"optional" json:"base" yaml:"base"`
	// The maximum amount of retries.
	// Default: 5.
	//
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
}

