package awsevents


// A `RetryPolicy` object that includes information about the retry policy settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   retryPolicyProperty := &retryPolicyProperty{
//   	maximumEventAgeInSeconds: jsii.Number(123),
//   	maximumRetryAttempts: jsii.Number(123),
//   }
//
type CfnRule_RetryPolicyProperty struct {
	// The maximum amount of time, in seconds, to continue to make retry attempts.
	MaximumEventAgeInSeconds *float64 `field:"optional" json:"maximumEventAgeInSeconds" yaml:"maximumEventAgeInSeconds"`
	// The maximum number of retry attempts to make before the request fails.
	//
	// Retry attempts continue until either the maximum number of attempts is made or until the duration of the `MaximumEventAgeInSeconds` is met.
	MaximumRetryAttempts *float64 `field:"optional" json:"maximumRetryAttempts" yaml:"maximumRetryAttempts"`
}

