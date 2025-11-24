package mixinsawsevidently


// A structure that defines one launch group in a launch.
//
// A launch group is a variation of the feature that you are including in the launch.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   launchGroupObjectProperty := &LaunchGroupObjectProperty{
//   	Description: jsii.String("description"),
//   	Feature: jsii.String("feature"),
//   	GroupName: jsii.String("groupName"),
//   	Variation: jsii.String("variation"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-launch-launchgroupobject.html
//
type CfnLaunchPropsMixin_LaunchGroupObjectProperty struct {
	// A description of the launch group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-launch-launchgroupobject.html#cfn-evidently-launch-launchgroupobject-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The feature that this launch is using.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-launch-launchgroupobject.html#cfn-evidently-launch-launchgroupobject-feature
	//
	Feature *string `field:"optional" json:"feature" yaml:"feature"`
	// A name for this launch group.
	//
	// It can include up to 127 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-launch-launchgroupobject.html#cfn-evidently-launch-launchgroupobject-groupname
	//
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// The feature variation to use for this launch group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-launch-launchgroupobject.html#cfn-evidently-launch-launchgroupobject-variation
	//
	Variation *string `field:"optional" json:"variation" yaml:"variation"`
}

