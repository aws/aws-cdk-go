package awss3


// A property for S3 Storage Lens advanced performance metrics.
//
// Advanced performance metrics provide insights into application performance such as access patterns and network originality metrics.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   advancedPerformanceMetricsProperty := &AdvancedPerformanceMetricsProperty{
//   	IsEnabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-advancedperformancemetrics.html
//
type CfnStorageLens_AdvancedPerformanceMetricsProperty struct {
	// This property indicates whether the advanced performance metrics are enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-advancedperformancemetrics.html#cfn-s3-storagelens-advancedperformancemetrics-isenabled
	//
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
}

