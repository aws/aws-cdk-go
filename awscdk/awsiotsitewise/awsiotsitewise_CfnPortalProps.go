package awsiotsitewise

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnPortal`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var alarms interface{}
//
//   cfnPortalProps := &cfnPortalProps{
//   	portalContactEmail: jsii.String("portalContactEmail"),
//   	portalName: jsii.String("portalName"),
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	alarms: alarms,
//   	notificationSenderEmail: jsii.String("notificationSenderEmail"),
//   	portalAuthMode: jsii.String("portalAuthMode"),
//   	portalDescription: jsii.String("portalDescription"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnPortalProps struct {
	// The AWS administrator's contact email address.
	PortalContactEmail *string `field:"required" json:"portalContactEmail" yaml:"portalContactEmail"`
	// A friendly name for the portal.
	PortalName *string `field:"required" json:"portalName" yaml:"portalName"`
	// The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of a service role that allows the portal's users to access your AWS IoT SiteWise resources on your behalf. For more information, see [Using service roles for AWS IoT SiteWise Monitor](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/monitor-service-role.html) in the *AWS IoT SiteWise User Guide* .
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Contains the configuration information of an alarm created in an AWS IoT SiteWise Monitor portal.
	//
	// You can use the alarm to monitor an asset property and get notified when the asset property value is outside a specified range. For more information, see [Monitoring with alarms](https://docs.aws.amazon.com/iot-sitewise/latest/appguide/monitor-alarms.html) in the *AWS IoT SiteWise Application Guide* .
	Alarms interface{} `field:"optional" json:"alarms" yaml:"alarms"`
	// The email address that sends alarm notifications.
	//
	// > If you use the [AWS IoT Events managed Lambda function](https://docs.aws.amazon.com/iotevents/latest/developerguide/lambda-support.html) to manage your emails, you must [verify the sender email address in Amazon SES](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/verify-email-addresses.html) .
	NotificationSenderEmail *string `field:"optional" json:"notificationSenderEmail" yaml:"notificationSenderEmail"`
	// The service to use to authenticate users to the portal. Choose from the following options:.
	//
	// - `SSO` – The portal uses AWS Single Sign-On to authenticate users and manage user permissions. Before you can create a portal that uses AWS SSO , you must enable AWS SSO . For more information, see [Enabling AWS SSO](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/monitor-get-started.html#mon-gs-sso) in the *AWS IoT SiteWise User Guide* . This option is only available in AWS Regions other than the China Regions.
	// - `IAM` – The portal uses AWS Identity and Access Management ( IAM ) to authenticate users and manage user permissions.
	//
	// You can't change this value after you create a portal.
	//
	// Default: `SSO`.
	PortalAuthMode *string `field:"optional" json:"portalAuthMode" yaml:"portalAuthMode"`
	// A description for the portal.
	PortalDescription *string `field:"optional" json:"portalDescription" yaml:"portalDescription"`
	// A list of key-value pairs that contain metadata for the portal.
	//
	// For more information, see [Tagging your AWS IoT SiteWise resources](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/tag-resources.html) in the *AWS IoT SiteWise User Guide* .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

