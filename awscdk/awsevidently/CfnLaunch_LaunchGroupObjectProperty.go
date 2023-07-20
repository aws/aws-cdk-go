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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-launch-launchgroupobject.html
//
type CfnLaunch_LaunchGroupObjectProperty struct {
	// The feature that this launch is using.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-launch-launchgroupobject.html#cfn-evidently-launch-launchgroupobject-feature
	//
	Feature *string `field:"required" json:"feature" yaml:"feature"`
	// A name for this launch group.
	//
	// It can include up to 127 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-launch-launchgroupobject.html#cfn-evidently-launch-launchgroupobject-groupname
	//
	GroupName *string `field:"required" json:"groupName" yaml:"groupName"`
	// The feature variation to use for this launch group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-launch-launchgroupobject.html#cfn-evidently-launch-launchgroupobject-variation
	//
	Variation *string `field:"required" json:"variation" yaml:"variation"`
	// A description of the launch group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-launch-launchgroupobject.html#cfn-evidently-launch-launchgroupobject-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

