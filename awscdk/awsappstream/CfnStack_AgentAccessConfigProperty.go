package awsappstream


// The configuration for agent access on a stack.
//
// Agent access enables AI agents to interact with desktop applications during streaming sessions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   agentAccessConfigProperty := &AgentAccessConfigProperty{
//   	ScreenImageFormat: jsii.String("screenImageFormat"),
//   	ScreenResolution: jsii.String("screenResolution"),
//   	Settings: []interface{}{
//   		&AgentAccessSettingProperty{
//   			AgentAction: jsii.String("agentAction"),
//   			Permission: jsii.String("permission"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	S3BucketArn: jsii.String("s3BucketArn"),
//   	ScreenshotsUploadEnabled: jsii.Boolean(false),
//   	UserControlMode: jsii.String("userControlMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-agentaccessconfig.html
//
type CfnStack_AgentAccessConfigProperty struct {
	// The image format for agent screen captures.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-agentaccessconfig.html#cfn-appstream-stack-agentaccessconfig-screenimageformat
	//
	ScreenImageFormat *string `field:"required" json:"screenImageFormat" yaml:"screenImageFormat"`
	// The screen resolution for the agent streaming environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-agentaccessconfig.html#cfn-appstream-stack-agentaccessconfig-screenresolution
	//
	ScreenResolution *string `field:"required" json:"screenResolution" yaml:"screenResolution"`
	// The list of agent access settings that define permissions for each agent action.
	//
	// You must specify at least one setting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-agentaccessconfig.html#cfn-appstream-stack-agentaccessconfig-settings
	//
	Settings interface{} `field:"required" json:"settings" yaml:"settings"`
	// The Amazon Resource Name (ARN) of the Amazon S3 bucket where agent screenshots are stored.
	//
	// Required when ScreenshotsUploadEnabled is true.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-agentaccessconfig.html#cfn-appstream-stack-agentaccessconfig-s3bucketarn
	//
	S3BucketArn *string `field:"optional" json:"s3BucketArn" yaml:"s3BucketArn"`
	// Indicates whether screenshot uploads to Amazon S3 are enabled for agent sessions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-agentaccessconfig.html#cfn-appstream-stack-agentaccessconfig-screenshotsuploadenabled
	//
	ScreenshotsUploadEnabled interface{} `field:"optional" json:"screenshotsUploadEnabled" yaml:"screenshotsUploadEnabled"`
	// The user control mode for agent sessions.
	//
	// This setting determines how users can interact with agent sessions. Valid values are VIEW_ONLY, VIEW_STOP, and DISABLED.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-agentaccessconfig.html#cfn-appstream-stack-agentaccessconfig-usercontrolmode
	//
	UserControlMode *string `field:"optional" json:"userControlMode" yaml:"userControlMode"`
}

