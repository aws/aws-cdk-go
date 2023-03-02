package awsimagebuilder


// Contains settings for the Systems Manager agent on your build instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   systemsManagerAgentProperty := &systemsManagerAgentProperty{
//   	uninstallAfterBuild: jsii.Boolean(false),
//   }
//
type CfnImageRecipe_SystemsManagerAgentProperty struct {
	// Controls whether the Systems Manager agent is removed from your final build image, prior to creating the new AMI.
	//
	// If this is set to true, then the agent is removed from the final image. If it's set to false, then the agent is left in, so that it is included in the new AMI. The default value is false.
	UninstallAfterBuild interface{} `field:"optional" json:"uninstallAfterBuild" yaml:"uninstallAfterBuild"`
}

