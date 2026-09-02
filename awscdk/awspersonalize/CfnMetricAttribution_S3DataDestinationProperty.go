package awspersonalize


// The configuration details of an Amazon S3 output bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3DataDestinationProperty := &S3DataDestinationProperty{
//   	Path: jsii.String("path"),
//
//   	// the properties below are optional
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-metricattribution-s3datadestination.html
//
type CfnMetricAttribution_S3DataDestinationProperty struct {
	// The file path of the Amazon S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-metricattribution-s3datadestination.html#cfn-personalize-metricattribution-s3datadestination-path
	//
	Path *string `field:"required" json:"path" yaml:"path"`
	// The ARN of the KMS key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-metricattribution-s3datadestination.html#cfn-personalize-metricattribution-s3datadestination-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

