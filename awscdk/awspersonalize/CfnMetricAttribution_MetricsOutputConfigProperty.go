package awspersonalize


// The output configuration details for the metric attribution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricsOutputConfigProperty := &MetricsOutputConfigProperty{
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	S3DataDestination: &S3DataDestinationProperty{
//   		Path: jsii.String("path"),
//
//   		// the properties below are optional
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-metricattribution-metricsoutputconfig.html
//
type CfnMetricAttribution_MetricsOutputConfigProperty struct {
	// The ARN of the IAM role for the metric attribution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-metricattribution-metricsoutputconfig.html#cfn-personalize-metricattribution-metricsoutputconfig-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The configuration details of an Amazon S3 output bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-metricattribution-metricsoutputconfig.html#cfn-personalize-metricattribution-metricsoutputconfig-s3datadestination
	//
	S3DataDestination interface{} `field:"optional" json:"s3DataDestination" yaml:"s3DataDestination"`
}

