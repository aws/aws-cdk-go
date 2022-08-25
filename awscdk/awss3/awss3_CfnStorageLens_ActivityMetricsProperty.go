package awss3


// This resource contains the details of the activity metrics for Amazon S3 Storage Lens.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   activityMetricsProperty := &activityMetricsProperty{
//   	isEnabled: jsii.Boolean(false),
//   }
//
type CfnStorageLens_ActivityMetricsProperty struct {
	// A property that indicates whether the activity metrics is enabled.
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
}

