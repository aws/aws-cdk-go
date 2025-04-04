package awss3


// This resource enables the Amazon CloudWatch publishing option for Amazon S3 Storage Lens metrics.
//
// For more information, see [Monitor S3 Storage Lens metrics in CloudWatch](https://docs.aws.amazon.com/AmazonS3/latest/userguide/storage_lens_view_metrics_cloudwatch.html) in the *Amazon S3 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchMetricsProperty := &CloudWatchMetricsProperty{
//   	IsEnabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-cloudwatchmetrics.html
//
type CfnStorageLens_CloudWatchMetricsProperty struct {
	// This property identifies whether the CloudWatch publishing option for S3 Storage Lens is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-cloudwatchmetrics.html#cfn-s3-storagelens-cloudwatchmetrics-isenabled
	//
	IsEnabled interface{} `field:"required" json:"isEnabled" yaml:"isEnabled"`
}

