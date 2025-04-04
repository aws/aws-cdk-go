package awsec2


// Attributes for an imported LaunchTemplate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   launchTemplateAttributes := &LaunchTemplateAttributes{
//   	LaunchTemplateId: jsii.String("launchTemplateId"),
//   	LaunchTemplateName: jsii.String("launchTemplateName"),
//   	VersionNumber: jsii.String("versionNumber"),
//   }
//
type LaunchTemplateAttributes struct {
	// The identifier of the Launch Template.
	//
	// Exactly one of `launchTemplateId` and `launchTemplateName` may be set.
	// Default: None.
	//
	LaunchTemplateId *string `field:"optional" json:"launchTemplateId" yaml:"launchTemplateId"`
	// The name of the Launch Template.
	//
	// Exactly one of `launchTemplateId` and `launchTemplateName` may be set.
	// Default: None.
	//
	LaunchTemplateName *string `field:"optional" json:"launchTemplateName" yaml:"launchTemplateName"`
	// The version number of this launch template to use.
	// Default: Version: "$Default".
	//
	VersionNumber *string `field:"optional" json:"versionNumber" yaml:"versionNumber"`
}

