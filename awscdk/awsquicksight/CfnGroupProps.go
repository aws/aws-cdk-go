package awsquicksight


// Properties for defining a `CfnGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGroupProps := &CfnGroupProps{
//   	GroupName: jsii.String("groupName"),
//   	Namespace: jsii.String("namespace"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-group.html
//
type CfnGroupProps struct {
	// A name for the group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-group.html#cfn-quicksight-group-groupname
	//
	GroupName *string `field:"required" json:"groupName" yaml:"groupName"`
	// The namespace that you want the group to be a part of.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-group.html#cfn-quicksight-group-namespace
	//
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// A description for the group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-group.html#cfn-quicksight-group-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

