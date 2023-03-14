package awsappstream

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnFleet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFleetProps := &cfnFleetProps{
//   	instanceType: jsii.String("instanceType"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	computeCapacity: &computeCapacityProperty{
//   		desiredInstances: jsii.Number(123),
//   	},
//   	description: jsii.String("description"),
//   	disconnectTimeoutInSeconds: jsii.Number(123),
//   	displayName: jsii.String("displayName"),
//   	domainJoinInfo: &domainJoinInfoProperty{
//   		directoryName: jsii.String("directoryName"),
//   		organizationalUnitDistinguishedName: jsii.String("organizationalUnitDistinguishedName"),
//   	},
//   	enableDefaultInternetAccess: jsii.Boolean(false),
//   	fleetType: jsii.String("fleetType"),
//   	iamRoleArn: jsii.String("iamRoleArn"),
//   	idleDisconnectTimeoutInSeconds: jsii.Number(123),
//   	imageArn: jsii.String("imageArn"),
//   	imageName: jsii.String("imageName"),
//   	maxConcurrentSessions: jsii.Number(123),
//   	maxUserDurationInSeconds: jsii.Number(123),
//   	platform: jsii.String("platform"),
//   	sessionScriptS3Location: &s3LocationProperty{
//   		s3Bucket: jsii.String("s3Bucket"),
//   		s3Key: jsii.String("s3Key"),
//   	},
//   	streamView: jsii.String("streamView"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	usbDeviceFilterStrings: []*string{
//   		jsii.String("usbDeviceFilterStrings"),
//   	},
//   	vpcConfig: &vpcConfigProperty{
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		subnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   }
//
type CfnFleetProps struct {
	// The instance type to use when launching fleet instances. The following instance types are available for non-Elastic fleets:.
	//
	// - stream.standard.small
	// - stream.standard.medium
	// - stream.standard.large
	// - stream.compute.large
	// - stream.compute.xlarge
	// - stream.compute.2xlarge
	// - stream.compute.4xlarge
	// - stream.compute.8xlarge
	// - stream.memory.large
	// - stream.memory.xlarge
	// - stream.memory.2xlarge
	// - stream.memory.4xlarge
	// - stream.memory.8xlarge
	// - stream.memory.z1d.large
	// - stream.memory.z1d.xlarge
	// - stream.memory.z1d.2xlarge
	// - stream.memory.z1d.3xlarge
	// - stream.memory.z1d.6xlarge
	// - stream.memory.z1d.12xlarge
	// - stream.graphics-design.large
	// - stream.graphics-design.xlarge
	// - stream.graphics-design.2xlarge
	// - stream.graphics-design.4xlarge
	// - stream.graphics-desktop.2xlarge
	// - stream.graphics.g4dn.xlarge
	// - stream.graphics.g4dn.2xlarge
	// - stream.graphics.g4dn.4xlarge
	// - stream.graphics.g4dn.8xlarge
	// - stream.graphics.g4dn.12xlarge
	// - stream.graphics.g4dn.16xlarge
	// - stream.graphics-pro.4xlarge
	// - stream.graphics-pro.8xlarge
	// - stream.graphics-pro.16xlarge
	//
	// The following instance types are available for Elastic fleets:
	//
	// - stream.standard.small
	// - stream.standard.medium
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// A unique name for the fleet.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The desired capacity for the fleet.
	//
	// This is not allowed for Elastic fleets.
	ComputeCapacity interface{} `field:"optional" json:"computeCapacity" yaml:"computeCapacity"`
	// The description to display.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The amount of time that a streaming session remains active after users disconnect.
	//
	// If users try to reconnect to the streaming session after a disconnection or network interruption within this time interval, they are connected to their previous session. Otherwise, they are connected to a new session with a new streaming instance.
	//
	// Specify a value between 60 and 360000.
	DisconnectTimeoutInSeconds *float64 `field:"optional" json:"disconnectTimeoutInSeconds" yaml:"disconnectTimeoutInSeconds"`
	// The fleet name to display.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The name of the directory and organizational unit (OU) to use to join the fleet to a Microsoft Active Directory domain.
	//
	// This is not allowed for Elastic fleets.
	DomainJoinInfo interface{} `field:"optional" json:"domainJoinInfo" yaml:"domainJoinInfo"`
	// Enables or disables default internet access for the fleet.
	EnableDefaultInternetAccess interface{} `field:"optional" json:"enableDefaultInternetAccess" yaml:"enableDefaultInternetAccess"`
	// The fleet type.
	//
	// - **ALWAYS_ON** - Provides users with instant-on access to their apps. You are charged for all running instances in your fleet, even if no users are streaming apps.
	// - **ON_DEMAND** - Provide users with access to applications after they connect, which takes one to two minutes. You are charged for instance streaming when users are connected and a small hourly fee for instances that are not streaming apps.
	// - **ELASTIC** - The pool of streaming instances is managed by Amazon AppStream 2.0. When a user selects their application or desktop to launch, they will start streaming after the app block has been downloaded and mounted to a streaming instance.
	//
	// *Allowed Values* : `ALWAYS_ON` | `ELASTIC` | `ON_DEMAND`.
	FleetType *string `field:"optional" json:"fleetType" yaml:"fleetType"`
	// The ARN of the IAM role that is applied to the fleet.
	//
	// To assume a role, the fleet instance calls the AWS Security Token Service `AssumeRole` API operation and passes the ARN of the role to use. The operation creates a new session with temporary credentials. AppStream 2.0 retrieves the temporary credentials and creates the *appstream_machine_role* credential profile on the instance.
	//
	// For more information, see [Using an IAM Role to Grant Permissions to Applications and Scripts Running on AppStream 2.0 Streaming Instances](https://docs.aws.amazon.com/appstream2/latest/developerguide/using-iam-roles-to-grant-permissions-to-applications-scripts-streaming-instances.html) in the *Amazon AppStream 2.0 Administration Guide* .
	IamRoleArn *string `field:"optional" json:"iamRoleArn" yaml:"iamRoleArn"`
	// The amount of time that users can be idle (inactive) before they are disconnected from their streaming session and the `DisconnectTimeoutInSeconds` time interval begins.
	//
	// Users are notified before they are disconnected due to inactivity. If they try to reconnect to the streaming session before the time interval specified in `DisconnectTimeoutInSeconds` elapses, they are connected to their previous session. Users are considered idle when they stop providing keyboard or mouse input during their streaming session. File uploads and downloads, audio in, audio out, and pixels changing do not qualify as user activity. If users continue to be idle after the time interval in `IdleDisconnectTimeoutInSeconds` elapses, they are disconnected.
	//
	// To prevent users from being disconnected due to inactivity, specify a value of 0. Otherwise, specify a value between 60 and 3600.
	//
	// If you enable this feature, we recommend that you specify a value that corresponds exactly to a whole number of minutes (for example, 60, 120, and 180). If you don't do this, the value is rounded to the nearest minute. For example, if you specify a value of 70, users are disconnected after 1 minute of inactivity. If you specify a value that is at the midpoint between two different minutes, the value is rounded up. For example, if you specify a value of 90, users are disconnected after 2 minutes of inactivity.
	IdleDisconnectTimeoutInSeconds *float64 `field:"optional" json:"idleDisconnectTimeoutInSeconds" yaml:"idleDisconnectTimeoutInSeconds"`
	// The ARN of the public, private, or shared image to use.
	ImageArn *string `field:"optional" json:"imageArn" yaml:"imageArn"`
	// The name of the image used to create the fleet.
	ImageName *string `field:"optional" json:"imageName" yaml:"imageName"`
	// The maximum number of concurrent sessions that can be run on an Elastic fleet.
	//
	// This setting is required for Elastic fleets, but is not used for other fleet types.
	MaxConcurrentSessions *float64 `field:"optional" json:"maxConcurrentSessions" yaml:"maxConcurrentSessions"`
	// The maximum amount of time that a streaming session can remain active, in seconds.
	//
	// If users are still connected to a streaming instance five minutes before this limit is reached, they are prompted to save any open documents before being disconnected. After this time elapses, the instance is terminated and replaced by a new instance.
	//
	// Specify a value between 600 and 360000.
	MaxUserDurationInSeconds *float64 `field:"optional" json:"maxUserDurationInSeconds" yaml:"maxUserDurationInSeconds"`
	// The platform of the fleet.
	//
	// Platform is a required setting for Elastic fleets, and is not used for other fleet types.
	//
	// *Allowed Values* : `WINDOWS_SERVER_2019` | `AMAZON_LINUX2`.
	Platform *string `field:"optional" json:"platform" yaml:"platform"`
	// The S3 location of the session scripts configuration zip file.
	//
	// This only applies to Elastic fleets.
	SessionScriptS3Location interface{} `field:"optional" json:"sessionScriptS3Location" yaml:"sessionScriptS3Location"`
	// The AppStream 2.0 view that is displayed to your users when they stream from the fleet. When `APP` is specified, only the windows of applications opened by users display. When `DESKTOP` is specified, the standard desktop that is provided by the operating system displays.
	//
	// The default value is `APP` .
	StreamView *string `field:"optional" json:"streamView" yaml:"streamView"`
	// An array of key-value pairs.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The USB device filter strings that specify which USB devices a user can redirect to the fleet streaming session, when using the Windows native client.
	//
	// This is allowed but not required for Elastic fleets.
	UsbDeviceFilterStrings *[]*string `field:"optional" json:"usbDeviceFilterStrings" yaml:"usbDeviceFilterStrings"`
	// The VPC configuration for the fleet.
	//
	// This is required for Elastic fleets, but not required for other fleet types.
	VpcConfig interface{} `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

