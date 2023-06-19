package awsiotanalytics

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnChannel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var serviceManagedS3 interface{}
//
//   cfnChannelProps := &CfnChannelProps{
//   	ChannelName: jsii.String("channelName"),
//   	ChannelStorage: &ChannelStorageProperty{
//   		CustomerManagedS3: &CustomerManagedS3Property{
//   			Bucket: jsii.String("bucket"),
//   			RoleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			KeyPrefix: jsii.String("keyPrefix"),
//   		},
//   		ServiceManagedS3: serviceManagedS3,
//   	},
//   	RetentionPeriod: &RetentionPeriodProperty{
//   		NumberOfDays: jsii.Number(123),
//   		Unlimited: jsii.Boolean(false),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnChannelProps struct {
	// The name of the channel.
	ChannelName *string `field:"optional" json:"channelName" yaml:"channelName"`
	// Where channel data is stored.
	ChannelStorage interface{} `field:"optional" json:"channelStorage" yaml:"channelStorage"`
	// How long, in days, message data is kept for the channel.
	RetentionPeriod interface{} `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
	// Metadata which can be used to manage the channel.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

