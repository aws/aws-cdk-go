package awss3


// Advanced Performance Metrics.
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
	// Specifies whether the Advanced Performance Metrics is enabled or disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-advancedperformancemetrics.html#cfn-s3-storagelens-advancedperformancemetrics-isenabled
	//
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
}

