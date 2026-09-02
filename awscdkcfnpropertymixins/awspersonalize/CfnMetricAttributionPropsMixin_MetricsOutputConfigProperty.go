package awspersonalize


// The output configuration details for the metric attribution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   metricsOutputConfigProperty := &MetricsOutputConfigProperty{
//   	RoleArn: jsii.String("roleArn"),
//   	S3DataDestination: &S3DataDestinationProperty{
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   		Path: jsii.String("path"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-metricattribution-metricsoutputconfig.html
//
type CfnMetricAttributionPropsMixin_MetricsOutputConfigProperty struct {
	// The ARN of the IAM role for the metric attribution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-metricattribution-metricsoutputconfig.html#cfn-personalize-metricattribution-metricsoutputconfig-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The configuration details of an Amazon S3 output bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-metricattribution-metricsoutputconfig.html#cfn-personalize-metricattribution-metricsoutputconfig-s3datadestination
	//
	S3DataDestination interface{} `field:"optional" json:"s3DataDestination" yaml:"s3DataDestination"`
}

