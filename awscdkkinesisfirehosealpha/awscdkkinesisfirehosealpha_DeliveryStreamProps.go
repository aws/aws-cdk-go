// The CDK Construct Library for AWS::KinesisFirehose
package awscdkkinesisfirehosealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesis"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties for a new delivery stream.
//
// Example:
//   // Specify the roles created above when defining the destination and delivery stream.
//   var bucket bucket
//   // Create service roles for the delivery stream and destination.
//   // These can be used for other purposes and granted access to different resources.
//   // They must include the Kinesis Data Firehose service principal in their trust policies.
//   // Two separate roles are shown below, but the same role can be used for both purposes.
//   deliveryStreamRole := iam.NewRole(this, jsii.String("Delivery Stream Role"), &roleProps{
//   	assumedBy: iam.NewServicePrincipal(jsii.String("firehose.amazonaws.com")),
//   })
//   destinationRole := iam.NewRole(this, jsii.String("Destination Role"), &roleProps{
//   	assumedBy: iam.NewServicePrincipal(jsii.String("firehose.amazonaws.com")),
//   })
//   destination := destinations.NewS3Bucket(bucket, &s3BucketProps{
//   	role: destinationRole,
//   })
//   firehose.NewDeliveryStream(this, jsii.String("Delivery Stream"), &deliveryStreamProps{
//   	destinations: []iDestination{
//   		destination,
//   	},
//   	role: deliveryStreamRole,
//   })
//
// Experimental.
type DeliveryStreamProps struct {
	// The destinations that this delivery stream will deliver data to.
	//
	// Only a singleton array is supported at this time.
	// Experimental.
	Destinations *[]IDestination `field:"required" json:"destinations" yaml:"destinations"`
	// A name for the delivery stream.
	// Experimental.
	DeliveryStreamName *string `field:"optional" json:"deliveryStreamName" yaml:"deliveryStreamName"`
	// Indicates the type of customer master key (CMK) to use for server-side encryption, if any.
	// Experimental.
	Encryption StreamEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// Customer managed key to server-side encrypt data in the stream.
	// Experimental.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// The IAM role associated with this delivery stream.
	//
	// Assumed by Kinesis Data Firehose to read from sources and encrypt data server-side.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The Kinesis data stream to use as a source for this delivery stream.
	// Experimental.
	SourceStream awskinesis.IStream `field:"optional" json:"sourceStream" yaml:"sourceStream"`
}

