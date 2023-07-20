package awskafkaconnect


// Details about delivering logs to Amazon S3.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3LogDeliveryProperty := &S3LogDeliveryProperty{
//   	Enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	Bucket: jsii.String("bucket"),
//   	Prefix: jsii.String("prefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-s3logdelivery.html
//
type CfnConnector_S3LogDeliveryProperty struct {
	// Specifies whether connector logs get sent to the specified Amazon S3 destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-s3logdelivery.html#cfn-kafkaconnect-connector-s3logdelivery-enabled
	//
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The name of the S3 bucket that is the destination for log delivery.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-s3logdelivery.html#cfn-kafkaconnect-connector-s3logdelivery-bucket
	//
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The S3 prefix that is the destination for log delivery.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-s3logdelivery.html#cfn-kafkaconnect-connector-s3logdelivery-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

