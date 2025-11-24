package mixinsawsiotsitewise

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnPortalPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var alarms interface{}
//
//   cfnPortalMixinProps := &CfnPortalMixinProps{
//   	Alarms: alarms,
//   	NotificationSenderEmail: jsii.String("notificationSenderEmail"),
//   	PortalAuthMode: jsii.String("portalAuthMode"),
//   	PortalContactEmail: jsii.String("portalContactEmail"),
//   	PortalDescription: jsii.String("portalDescription"),
//   	PortalName: jsii.String("portalName"),
//   	PortalType: jsii.String("portalType"),
//   	PortalTypeConfiguration: map[string]interface{}{
//   		"portalTypeConfigurationKey": &PortalTypeEntryProperty{
//   			"portalTools": []*string{
//   				jsii.String("portalTools"),
//   			},
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-portal.html
//
type CfnPortalMixinProps struct {
	// Contains the configuration information of an alarm created in an AWS IoT SiteWise Monitor portal.
	//
	// You can use the alarm to monitor an asset property and get notified when the asset property value is outside a specified range. For more information, see [Monitoring with alarms](https://docs.aws.amazon.com/iot-sitewise/latest/appguide/monitor-alarms.html) in the *AWS IoT SiteWise Application Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-portal.html#cfn-iotsitewise-portal-alarms
	//
	Alarms interface{} `field:"optional" json:"alarms" yaml:"alarms"`
	// The email address that sends alarm notifications.
	//
	// > If you use the [AWS IoT Events managed Lambda function](https://docs.aws.amazon.com/iotevents/latest/developerguide/lambda-support.html) to manage your emails, you must [verify the sender email address in Amazon SES](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/verify-email-addresses.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-portal.html#cfn-iotsitewise-portal-notificationsenderemail
	//
	NotificationSenderEmail *string `field:"optional" json:"notificationSenderEmail" yaml:"notificationSenderEmail"`
	// The service to use to authenticate users to the portal. Choose from the following options:.
	//
	// - `SSO` – The portal uses AWS IAM Identity Center to authenticate users and manage user permissions. Before you can create a portal that uses IAM Identity Center, you must enable IAM Identity Center. For more information, see [Enabling IAM Identity Center](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/monitor-get-started.html#mon-gs-sso) in the *AWS IoT SiteWise User Guide* . This option is only available in AWS Regions other than the China Regions.
	// - `IAM` – The portal uses AWS Identity and Access Management to authenticate users and manage user permissions.
	//
	// You can't change this value after you create a portal.
	//
	// Default: `SSO`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-portal.html#cfn-iotsitewise-portal-portalauthmode
	//
	PortalAuthMode *string `field:"optional" json:"portalAuthMode" yaml:"portalAuthMode"`
	// The AWS administrator's contact email address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-portal.html#cfn-iotsitewise-portal-portalcontactemail
	//
	PortalContactEmail *string `field:"optional" json:"portalContactEmail" yaml:"portalContactEmail"`
	// A description for the portal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-portal.html#cfn-iotsitewise-portal-portaldescription
	//
	PortalDescription *string `field:"optional" json:"portalDescription" yaml:"portalDescription"`
	// A friendly name for the portal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-portal.html#cfn-iotsitewise-portal-portalname
	//
	PortalName *string `field:"optional" json:"portalName" yaml:"portalName"`
	// Define the type of portal.
	//
	// The value for AWS IoT SiteWise Monitor (Classic) is `SITEWISE_PORTAL_V1` . The value for AWS IoT SiteWise Monitor (AI-aware) is `SITEWISE_PORTAL_V2` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-portal.html#cfn-iotsitewise-portal-portaltype
	//
	PortalType *string `field:"optional" json:"portalType" yaml:"portalType"`
	// Map to associate detail of configuration related with a PortalType.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-portal.html#cfn-iotsitewise-portal-portaltypeconfiguration
	//
	PortalTypeConfiguration interface{} `field:"optional" json:"portalTypeConfiguration" yaml:"portalTypeConfiguration"`
	// The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of a service role that allows the portal's users to access your AWS IoT SiteWise resources on your behalf. For more information, see [Using service roles for AWS IoT SiteWise Monitor](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/monitor-service-role.html) in the *AWS IoT SiteWise User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-portal.html#cfn-iotsitewise-portal-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// A list of key-value pairs that contain metadata for the portal.
	//
	// For more information, see [Tagging your AWS IoT SiteWise resources](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/tag-resources.html) in the *AWS IoT SiteWise User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-portal.html#cfn-iotsitewise-portal-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

