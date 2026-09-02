package awspersonalize


// The configuration details of an Amazon S3 output bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   s3DataDestinationProperty := &S3DataDestinationProperty{
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	Path: jsii.String("path"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-metricattribution-s3datadestination.html
//
type CfnMetricAttributionPropsMixin_S3DataDestinationProperty struct {
	// The ARN of the KMS key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-metricattribution-s3datadestination.html#cfn-personalize-metricattribution-s3datadestination-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The file path of the Amazon S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-metricattribution-s3datadestination.html#cfn-personalize-metricattribution-s3datadestination-path
	//
	Path *string `field:"optional" json:"path" yaml:"path"`
}

