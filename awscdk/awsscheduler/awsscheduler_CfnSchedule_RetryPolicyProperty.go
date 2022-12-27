package awsscheduler


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
type CfnSchedule_RetryPolicyProperty struct {
	// `CfnSchedule.RetryPolicyProperty.MaximumEventAgeInSeconds`.
	MaximumEventAgeInSeconds *float64 `field:"optional" json:"maximumEventAgeInSeconds" yaml:"maximumEventAgeInSeconds"`
	// `CfnSchedule.RetryPolicyProperty.MaximumRetryAttempts`.
	MaximumRetryAttempts *float64 `field:"optional" json:"maximumRetryAttempts" yaml:"maximumRetryAttempts"`
}

