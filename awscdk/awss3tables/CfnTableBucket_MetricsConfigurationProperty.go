package awss3tables


// Settings governing the Metric configuration for the table bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricsConfigurationProperty := &MetricsConfigurationProperty{
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-tablebucket-metricsconfiguration.html
//
type CfnTableBucket_MetricsConfigurationProperty struct {
	// Indicates whether Metrics are enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-tablebucket-metricsconfiguration.html#cfn-s3tables-tablebucket-metricsconfiguration-status
	//
	// Default: - "Disabled".
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

