package awseventsv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnEventBus`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEventBusProps := &CfnEventBusProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	EncryptionConfiguration: &EncryptionConfigurationProperty{
//   		KmsKeyIdentifier: jsii.String("kmsKeyIdentifier"),
//   	},
//   	StorageConfiguration: &StorageConfigurationProperty{
//   		RetentionPeriodInDays: jsii.Number(123),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventsv2-eventbus.html
//
type CfnEventBusProps struct {
	// The name of the event bus.
	//
	// The first character must be alphanumeric; the remaining characters may also include '.', '-', and '_'.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventsv2-eventbus.html#cfn-eventsv2-eventbus-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description of the event bus.
	//
	// Control characters and Unicode line separators are not allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventsv2-eventbus.html#cfn-eventsv2-eventbus-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Encryption configuration for an event bus.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventsv2-eventbus.html#cfn-eventsv2-eventbus-encryptionconfiguration
	//
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// Storage (retention) configuration for the event bus.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventsv2-eventbus.html#cfn-eventsv2-eventbus-storageconfiguration
	//
	StorageConfiguration interface{} `field:"optional" json:"storageConfiguration" yaml:"storageConfiguration"`
	// The tags assigned to the event bus.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventsv2-eventbus.html#cfn-eventsv2-eventbus-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

