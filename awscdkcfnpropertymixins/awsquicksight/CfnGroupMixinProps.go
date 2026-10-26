package awsquicksight


// Properties for CfnGroupPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnGroupMixinProps := &CfnGroupMixinProps{
//   	Description: jsii.String("description"),
//   	GroupName: jsii.String("groupName"),
//   	Namespace: jsii.String("namespace"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-group.html
//
type CfnGroupMixinProps struct {
	// A description for the group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-group.html#cfn-quicksight-group-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A name for the group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-group.html#cfn-quicksight-group-groupname
	//
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// The namespace that you want the group to be a part of.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-group.html#cfn-quicksight-group-namespace
	//
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

