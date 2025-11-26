package previewawskafkaconnectmixins


// Details about delivering logs to Amazon S3.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3LogDeliveryProperty := &S3LogDeliveryProperty{
//   	Bucket: jsii.String("bucket"),
//   	Enabled: jsii.Boolean(false),
//   	Prefix: jsii.String("prefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-s3logdelivery.html
//
type CfnConnectorPropsMixin_S3LogDeliveryProperty struct {
	// The name of the S3 bucket that is the destination for log delivery.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-s3logdelivery.html#cfn-kafkaconnect-connector-s3logdelivery-bucket
	//
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// Specifies whether connector logs get sent to the specified Amazon S3 destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-s3logdelivery.html#cfn-kafkaconnect-connector-s3logdelivery-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The S3 prefix that is the destination for log delivery.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-s3logdelivery.html#cfn-kafkaconnect-connector-s3logdelivery-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

