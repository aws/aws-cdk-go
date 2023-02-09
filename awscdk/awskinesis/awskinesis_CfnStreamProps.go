package awskinesis

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnStream`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStreamProps := &cfnStreamProps{
//   	name: jsii.String("name"),
//   	retentionPeriodHours: jsii.Number(123),
//   	shardCount: jsii.Number(123),
//   	streamEncryption: &streamEncryptionProperty{
//   		encryptionType: jsii.String("encryptionType"),
//   		keyId: jsii.String("keyId"),
//   	},
//   	streamModeDetails: &streamModeDetailsProperty{
//   		streamMode: jsii.String("streamMode"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnStreamProps struct {
	// The name of the Kinesis stream.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the stream name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The number of hours for the data records that are stored in shards to remain accessible.
	//
	// The default value is 24. For more information about the stream retention period, see [Changing the Data Retention Period](https://docs.aws.amazon.com/streams/latest/dev/kinesis-extended-retention.html) in the Amazon Kinesis Developer Guide.
	RetentionPeriodHours *float64 `field:"optional" json:"retentionPeriodHours" yaml:"retentionPeriodHours"`
	// The number of shards that the stream uses.
	//
	// For greater provisioned throughput, increase the number of shards.
	ShardCount *float64 `field:"optional" json:"shardCount" yaml:"shardCount"`
	// When specified, enables or updates server-side encryption using an AWS KMS key for a specified stream.
	//
	// Removing this property from your stack template and updating your stack disables encryption.
	StreamEncryption interface{} `field:"optional" json:"streamEncryption" yaml:"streamEncryption"`
	// Specifies the capacity mode to which you want to set your data stream.
	//
	// Currently, in Kinesis Data Streams, you can choose between an *on-demand* capacity mode and a *provisioned* capacity mode for your data streams.
	StreamModeDetails interface{} `field:"optional" json:"streamModeDetails" yaml:"streamModeDetails"`
	// An arbitrary set of tags (key–value pairs) to associate with the Kinesis stream.
	//
	// For information about constraints for this property, see [Tag Restrictions](https://docs.aws.amazon.com/streams/latest/dev/tagging.html#tagging-restrictions) in the *Amazon Kinesis Developer Guide* .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

