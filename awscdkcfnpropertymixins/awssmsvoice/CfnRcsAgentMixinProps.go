package awssmsvoice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnRcsAgentPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnRcsAgentMixinProps := &CfnRcsAgentMixinProps{
//   	DeletionProtectionEnabled: jsii.Boolean(false),
//   	OptOutListName: jsii.String("optOutListName"),
//   	SelfManagedOptOutsEnabled: jsii.Boolean(false),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TwoWayChannelArn: jsii.String("twoWayChannelArn"),
//   	TwoWayChannelRole: jsii.String("twoWayChannelRole"),
//   	TwoWayEnabled: jsii.Boolean(false),
//   	TwoWayMediaS3BucketName: jsii.String("twoWayMediaS3BucketName"),
//   	TwoWayMediaS3KeyPrefix: jsii.String("twoWayMediaS3KeyPrefix"),
//   	TwoWayMediaS3Role: jsii.String("twoWayMediaS3Role"),
//   	TwoWayRcsEventsEnabled: []*string{
//   		jsii.String("twoWayRcsEventsEnabled"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-rcsagent.html
//
type CfnRcsAgentMixinProps struct {
	// When set to true the RCS agent can't be deleted.
	//
	// By default this is false.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-rcsagent.html#cfn-smsvoice-rcsagent-deletionprotectionenabled
	//
	DeletionProtectionEnabled interface{} `field:"optional" json:"deletionProtectionEnabled" yaml:"deletionProtectionEnabled"`
	// The name of the opt-out list associated with the RCS agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-rcsagent.html#cfn-smsvoice-rcsagent-optoutlistname
	//
	OptOutListName *string `field:"optional" json:"optOutListName" yaml:"optOutListName"`
	// When set to true you're responsible for responding to HELP and STOP requests, and for tracking and honoring opt-out requests.
	//
	// By default this is false.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-rcsagent.html#cfn-smsvoice-rcsagent-selfmanagedoptoutsenabled
	//
	SelfManagedOptOutsEnabled interface{} `field:"optional" json:"selfManagedOptOutsEnabled" yaml:"selfManagedOptOutsEnabled"`
	// An array of key-value pairs to apply to the RCS agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-rcsagent.html#cfn-smsvoice-rcsagent-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The Amazon Resource Name (ARN) of the two way channel where inbound messages are delivered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-rcsagent.html#cfn-smsvoice-rcsagent-twowaychannelarn
	//
	TwoWayChannelArn *string `field:"optional" json:"twoWayChannelArn" yaml:"twoWayChannelArn"`
	// The Amazon Resource Name (ARN) of an IAM role for the service to assume in order to post inbound messages to the two way channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-rcsagent.html#cfn-smsvoice-rcsagent-twowaychannelrole
	//
	TwoWayChannelRole *string `field:"optional" json:"twoWayChannelRole" yaml:"twoWayChannelRole"`
	// When set to true two-way messaging is enabled for the RCS agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-rcsagent.html#cfn-smsvoice-rcsagent-twowayenabled
	//
	TwoWayEnabled interface{} `field:"optional" json:"twoWayEnabled" yaml:"twoWayEnabled"`
	// The name of the Amazon S3 bucket where inbound RCS media objects are written.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-rcsagent.html#cfn-smsvoice-rcsagent-twowaymedias3bucketname
	//
	TwoWayMediaS3BucketName *string `field:"optional" json:"twoWayMediaS3BucketName" yaml:"twoWayMediaS3BucketName"`
	// The key prefix used for inbound RCS media objects in the Amazon S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-rcsagent.html#cfn-smsvoice-rcsagent-twowaymedias3keyprefix
	//
	TwoWayMediaS3KeyPrefix *string `field:"optional" json:"twoWayMediaS3KeyPrefix" yaml:"twoWayMediaS3KeyPrefix"`
	// The Amazon Resource Name (ARN) of the IAM role used to write inbound RCS media files to the Amazon S3 bucket.
	//
	// The role must have s3:PutObject permission on the bucket and a trust policy allowing sms-voice.amazonaws.com to assume it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-rcsagent.html#cfn-smsvoice-rcsagent-twowaymedias3role
	//
	TwoWayMediaS3Role *string `field:"optional" json:"twoWayMediaS3Role" yaml:"twoWayMediaS3Role"`
	// The list of RCS event types enabled for two-way messaging.
	//
	// An empty list disables all event types. The special value ALL enables all current and future event types and must be the only element if used. Requires TwoWayEnabled to be true.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-rcsagent.html#cfn-smsvoice-rcsagent-twowayrcseventsenabled
	//
	TwoWayRcsEventsEnabled *[]*string `field:"optional" json:"twoWayRcsEventsEnabled" yaml:"twoWayRcsEventsEnabled"`
}

