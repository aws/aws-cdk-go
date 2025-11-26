package previewawsevidentlymixins


// A structure containing the percentage of launch traffic to allocate to one launch group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   groupToWeightProperty := &GroupToWeightProperty{
//   	GroupName: jsii.String("groupName"),
//   	SplitWeight: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-launch-grouptoweight.html
//
type CfnLaunchPropsMixin_GroupToWeightProperty struct {
	// The name of the launch group.
	//
	// It can include up to 127 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-launch-grouptoweight.html#cfn-evidently-launch-grouptoweight-groupname
	//
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// The portion of launch traffic to allocate to this launch group.
	//
	// This is represented in thousandths of a percent. For example, specify 20,000 to allocate 20% of the launch audience to this launch group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-launch-grouptoweight.html#cfn-evidently-launch-grouptoweight-splitweight
	//
	SplitWeight *float64 `field:"optional" json:"splitWeight" yaml:"splitWeight"`
}

