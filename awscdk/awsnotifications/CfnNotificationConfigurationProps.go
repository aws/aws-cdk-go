package awsnotifications

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnNotificationConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnNotificationConfigurationProps := &CfnNotificationConfigurationProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	AggregationDuration: jsii.String("aggregationDuration"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notifications-notificationconfiguration.html
//
type CfnNotificationConfigurationProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notifications-notificationconfiguration.html#cfn-notifications-notificationconfiguration-description
	//
	Description *string `field:"required" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notifications-notificationconfiguration.html#cfn-notifications-notificationconfiguration-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notifications-notificationconfiguration.html#cfn-notifications-notificationconfiguration-aggregationduration
	//
	AggregationDuration *string `field:"optional" json:"aggregationDuration" yaml:"aggregationDuration"`
	// A list of tags that are attached to the role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notifications-notificationconfiguration.html#cfn-notifications-notificationconfiguration-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

