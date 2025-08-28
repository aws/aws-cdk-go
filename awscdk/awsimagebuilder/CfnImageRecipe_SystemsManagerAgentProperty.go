package awsimagebuilder


// Contains settings for the Systems Manager agent on your build instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   systemsManagerAgentProperty := &SystemsManagerAgentProperty{
//   	UninstallAfterBuild: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-systemsmanageragent.html
//
type CfnImageRecipe_SystemsManagerAgentProperty struct {
	// Controls whether the Systems Manager agent is removed from your final build image, prior to creating the new AMI.
	//
	// If this is set to true, then the agent is removed from the final image. If it's set to false, then the agent is left in, so that it is included in the new AMI. default value is false.
	//
	// The default behavior of uninstallAfterBuild is to remove the SSM Agent if it was installed by EC2 Image Builder.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-systemsmanageragent.html#cfn-imagebuilder-imagerecipe-systemsmanageragent-uninstallafterbuild
	//
	UninstallAfterBuild interface{} `field:"optional" json:"uninstallAfterBuild" yaml:"uninstallAfterBuild"`
}

