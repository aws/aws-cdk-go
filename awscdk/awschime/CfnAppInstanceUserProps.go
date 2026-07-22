package awschime

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnAppInstanceUser`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAppInstanceUserProps := &CfnAppInstanceUserProps{
//   	AppInstanceArn: jsii.String("appInstanceArn"),
//   	AppInstanceUserId: jsii.String("appInstanceUserId"),
//
//   	// the properties below are optional
//   	ExpirationSettings: &ExpirationSettingsProperty{
//   		ExpirationCriterion: jsii.String("expirationCriterion"),
//   		ExpirationDays: jsii.Number(123),
//   	},
//   	Metadata: jsii.String("metadata"),
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-appinstanceuser.html
//
type CfnAppInstanceUserProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-appinstanceuser.html#cfn-chime-appinstanceuser-appinstancearn
	//
	AppInstanceArn *string `field:"required" json:"appInstanceArn" yaml:"appInstanceArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-appinstanceuser.html#cfn-chime-appinstanceuser-appinstanceuserid
	//
	AppInstanceUserId *string `field:"required" json:"appInstanceUserId" yaml:"appInstanceUserId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-appinstanceuser.html#cfn-chime-appinstanceuser-expirationsettings
	//
	ExpirationSettings interface{} `field:"optional" json:"expirationSettings" yaml:"expirationSettings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-appinstanceuser.html#cfn-chime-appinstanceuser-metadata
	//
	Metadata *string `field:"optional" json:"metadata" yaml:"metadata"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-appinstanceuser.html#cfn-chime-appinstanceuser-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-appinstanceuser.html#cfn-chime-appinstanceuser-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

