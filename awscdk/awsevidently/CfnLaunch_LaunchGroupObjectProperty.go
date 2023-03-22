package awsevidently


// A structure that defines one launch group in a launch.
//
// A launch group is a variation of the feature that you are including in the launch.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   launchGroupObjectProperty := &LaunchGroupObjectProperty{
//   	Feature: jsii.String("feature"),
//   	GroupName: jsii.String("groupName"),
//   	Variation: jsii.String("variation"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
type CfnLaunch_LaunchGroupObjectProperty struct {
	// The feature that this launch is using.
	Feature *string `field:"required" json:"feature" yaml:"feature"`
	// A name for this launch group.
	//
	// It can include up to 127 characters.
	GroupName *string `field:"required" json:"groupName" yaml:"groupName"`
	// The feature variation to use for this launch group.
	Variation *string `field:"required" json:"variation" yaml:"variation"`
	// A description of the launch group.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

