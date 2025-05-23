package awssynthetics


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   retryConfigProperty := &RetryConfigProperty{
//   	MaxRetries: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-retryconfig.html
//
type CfnCanary_RetryConfigProperty struct {
	// maximum times the canary will be retried upon the scheduled run failure.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-retryconfig.html#cfn-synthetics-canary-retryconfig-maxretries
	//
	MaxRetries *float64 `field:"required" json:"maxRetries" yaml:"maxRetries"`
}

