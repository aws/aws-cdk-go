package previewawsrolesanywheremixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnTrustAnchorPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTrustAnchorMixinProps := &CfnTrustAnchorMixinProps{
//   	Enabled: jsii.Boolean(false),
//   	Name: jsii.String("name"),
//   	NotificationSettings: []interface{}{
//   		&NotificationSettingProperty{
//   			Channel: jsii.String("channel"),
//   			Enabled: jsii.Boolean(false),
//   			Event: jsii.String("event"),
//   			Threshold: jsii.Number(123),
//   		},
//   	},
//   	Source: &SourceProperty{
//   		SourceData: &SourceDataProperty{
//   			AcmPcaArn: jsii.String("acmPcaArn"),
//   			X509CertificateData: jsii.String("x509CertificateData"),
//   		},
//   		SourceType: jsii.String("sourceType"),
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
type CfnTrustAnchorMixinProps struct {
	// Indicates whether the trust anchor is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-trustanchor.html#cfn-rolesanywhere-trustanchor-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The name of the trust anchor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-trustanchor.html#cfn-rolesanywhere-trustanchor-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of notification settings to be associated to the trust anchor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-trustanchor.html#cfn-rolesanywhere-trustanchor-notificationsettings
	//
	NotificationSettings interface{} `field:"optional" json:"notificationSettings" yaml:"notificationSettings"`
	// The trust anchor type and its related certificate data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-trustanchor.html#cfn-rolesanywhere-trustanchor-source
	//
	Source interface{} `field:"optional" json:"source" yaml:"source"`
	// The tags to attach to the trust anchor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-trustanchor.html#cfn-rolesanywhere-trustanchor-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

