package awsrolesanywhere

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnTrustAnchor`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTrustAnchorProps := &CfnTrustAnchorProps{
//   	Name: jsii.String("name"),
//   	Source: &SourceProperty{
//   		SourceData: &SourceDataProperty{
//   			AcmPcaArn: jsii.String("acmPcaArn"),
//   			X509CertificateData: jsii.String("x509CertificateData"),
//   		},
//   		SourceType: jsii.String("sourceType"),
//   	},
//
//   	// the properties below are optional
//   	Enabled: jsii.Boolean(false),
//   	NotificationSettings: []interface{}{
//   		&NotificationSettingProperty{
//   			Enabled: jsii.Boolean(false),
//   			Event: jsii.String("event"),
//
//   			// the properties below are optional
//   			Channel: jsii.String("channel"),
//   			Threshold: jsii.Number(123),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-trustanchor.html
//
type CfnTrustAnchorProps struct {
	// The name of the trust anchor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-trustanchor.html#cfn-rolesanywhere-trustanchor-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The trust anchor type and its related certificate data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-trustanchor.html#cfn-rolesanywhere-trustanchor-source
	//
	Source interface{} `field:"required" json:"source" yaml:"source"`
	// Indicates whether the trust anchor is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-trustanchor.html#cfn-rolesanywhere-trustanchor-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// A list of notification settings to be associated to the trust anchor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-trustanchor.html#cfn-rolesanywhere-trustanchor-notificationsettings
	//
	NotificationSettings interface{} `field:"optional" json:"notificationSettings" yaml:"notificationSettings"`
	// The tags to attach to the trust anchor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-trustanchor.html#cfn-rolesanywhere-trustanchor-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

