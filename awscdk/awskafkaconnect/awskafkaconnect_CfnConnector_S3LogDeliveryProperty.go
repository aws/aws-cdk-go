package awskafkaconnect


// Details about delivering logs to Amazon S3.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3LogDeliveryProperty := &s3LogDeliveryProperty{
//   	enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	bucket: jsii.String("bucket"),
//   	prefix: jsii.String("prefix"),
//   }
//
type CfnConnector_S3LogDeliveryProperty struct {
	// Specifies whether connector logs get sent to the specified Amazon S3 destination.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The name of the S3 bucket that is the destination for log delivery.
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The S3 prefix that is the destination for log delivery.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

