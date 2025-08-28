package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Options for attaching a CloudFormationInit to a resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cfnResource cfnResource
//   var role role
//   var userData userData
//
//   attachInitOptions := &AttachInitOptions{
//   	InstanceRole: role,
//   	Platform: awscdk.Aws_ec2.OperatingSystemType_LINUX,
//   	UserData: userData,
//
//   	// the properties below are optional
//   	ConfigSets: []*string{
//   		jsii.String("configSets"),
//   	},
//   	EmbedFingerprint: jsii.Boolean(false),
//   	IgnoreFailures: jsii.Boolean(false),
//   	IncludeRole: jsii.Boolean(false),
//   	IncludeUrl: jsii.Boolean(false),
//   	PrintLog: jsii.Boolean(false),
//   	SignalResource: cfnResource,
//   }
//
type AttachInitOptions struct {
	// Instance role of the consuming instance or fleet.
	InstanceRole awsiam.IRole `field:"required" json:"instanceRole" yaml:"instanceRole"`
	// OS Platform the init config will be used for.
	Platform OperatingSystemType `field:"required" json:"platform" yaml:"platform"`
	// UserData to add commands to.
	UserData UserData `field:"required" json:"userData" yaml:"userData"`
	// ConfigSet to activate.
	// Default: ['default'].
	//
	ConfigSets *[]*string `field:"optional" json:"configSets" yaml:"configSets"`
	// Whether to embed a hash into the userData.
	//
	// If `true` (the default), a hash of the config will be embedded into the
	// UserData, so that if the config changes, the UserData changes and
	// the instance will be replaced.
	//
	// If `false`, no such hash will be embedded, and if the CloudFormation Init
	// config changes nothing will happen to the running instance.
	// Default: true.
	//
	EmbedFingerprint *bool `field:"optional" json:"embedFingerprint" yaml:"embedFingerprint"`
	// Don't fail the instance creation when cfn-init fails.
	//
	// You can use this to prevent CloudFormation from rolling back when
	// instances fail to start up, to help in debugging.
	// Default: false.
	//
	IgnoreFailures *bool `field:"optional" json:"ignoreFailures" yaml:"ignoreFailures"`
	// Include --role argument when running cfn-init and cfn-signal commands.
	//
	// This will be the IAM instance profile attached to the EC2 instance.
	// Default: false.
	//
	IncludeRole *bool `field:"optional" json:"includeRole" yaml:"includeRole"`
	// Include --url argument when running cfn-init and cfn-signal commands.
	//
	// This will be the cloudformation endpoint in the deployed region
	// e.g. https://cloudformation.us-east-1.amazonaws.com
	// Default: false.
	//
	IncludeUrl *bool `field:"optional" json:"includeUrl" yaml:"includeUrl"`
	// Print the results of running cfn-init to the Instance System Log.
	//
	// By default, the output of running cfn-init is written to a log file
	// on the instance. Set this to `true` to print it to the System Log
	// (visible from the EC2 Console), `false` to not print it.
	//
	// (Be aware that the system log is refreshed at certain points in
	// time of the instance life cycle, and successful execution may
	// not always show up).
	// Default: true.
	//
	PrintLog *bool `field:"optional" json:"printLog" yaml:"printLog"`
	// When provided, signals this resource instead of the attached resource.
	//
	// You can use this to support signaling LaunchTemplate while attaching AutoScalingGroup.
	// Default: - if this property is undefined cfn-signal signals the attached resource.
	//
	SignalResource awscdk.CfnResource `field:"optional" json:"signalResource" yaml:"signalResource"`
}

