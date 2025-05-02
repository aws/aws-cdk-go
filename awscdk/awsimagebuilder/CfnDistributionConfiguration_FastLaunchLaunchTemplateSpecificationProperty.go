package awsimagebuilder


// Identifies the launch template that the associated Windows AMI uses for launching an instance when faster launching is enabled.
//
// > You can specify either the `launchTemplateName` or the `launchTemplateId` , but not both.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fastLaunchLaunchTemplateSpecificationProperty := &FastLaunchLaunchTemplateSpecificationProperty{
//   	LaunchTemplateId: jsii.String("launchTemplateId"),
//   	LaunchTemplateName: jsii.String("launchTemplateName"),
//   	LaunchTemplateVersion: jsii.String("launchTemplateVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-fastlaunchlaunchtemplatespecification.html
//
type CfnDistributionConfiguration_FastLaunchLaunchTemplateSpecificationProperty struct {
	// The ID of the launch template to use for faster launching for a Windows AMI.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-fastlaunchlaunchtemplatespecification.html#cfn-imagebuilder-distributionconfiguration-fastlaunchlaunchtemplatespecification-launchtemplateid
	//
	LaunchTemplateId *string `field:"optional" json:"launchTemplateId" yaml:"launchTemplateId"`
	// The name of the launch template to use for faster launching for a Windows AMI.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-fastlaunchlaunchtemplatespecification.html#cfn-imagebuilder-distributionconfiguration-fastlaunchlaunchtemplatespecification-launchtemplatename
	//
	LaunchTemplateName *string `field:"optional" json:"launchTemplateName" yaml:"launchTemplateName"`
	// The version of the launch template to use for faster launching for a Windows AMI.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-fastlaunchlaunchtemplatespecification.html#cfn-imagebuilder-distributionconfiguration-fastlaunchlaunchtemplatespecification-launchtemplateversion
	//
	LaunchTemplateVersion *string `field:"optional" json:"launchTemplateVersion" yaml:"launchTemplateVersion"`
}

