package awsgamelift


// *This data type is used with the GameLift FleetIQ and game server groups.*.
//
// An Amazon EC2 launch template that contains configuration settings and game server code to be deployed to all instances in a game server group. The launch template is specified when creating a new game server group with `GameServerGroup` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   launchTemplateProperty := &launchTemplateProperty{
//   	launchTemplateId: jsii.String("launchTemplateId"),
//   	launchTemplateName: jsii.String("launchTemplateName"),
//   	version: jsii.String("version"),
//   }
//
type CfnGameServerGroup_LaunchTemplateProperty struct {
	// A unique identifier for an existing Amazon EC2 launch template.
	LaunchTemplateId *string `field:"optional" json:"launchTemplateId" yaml:"launchTemplateId"`
	// A readable identifier for an existing Amazon EC2 launch template.
	LaunchTemplateName *string `field:"optional" json:"launchTemplateName" yaml:"launchTemplateName"`
	// The version of the Amazon EC2 launch template to use.
	//
	// If no version is specified, the default version will be used. With Amazon EC2, you can specify a default version for a launch template. If none is set, the default is the first version created.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

