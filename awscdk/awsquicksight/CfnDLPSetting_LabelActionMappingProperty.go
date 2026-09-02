package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   labelActionMappingProperty := &LabelActionMappingProperty{
//   	Action: jsii.String("action"),
//   	LabelId: jsii.String("labelId"),
//   	LabelName: jsii.String("labelName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dlpsetting-labelactionmapping.html
//
type CfnDLPSetting_LabelActionMappingProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dlpsetting-labelactionmapping.html#cfn-quicksight-dlpsetting-labelactionmapping-action
	//
	Action *string `field:"required" json:"action" yaml:"action"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dlpsetting-labelactionmapping.html#cfn-quicksight-dlpsetting-labelactionmapping-labelid
	//
	LabelId *string `field:"required" json:"labelId" yaml:"labelId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dlpsetting-labelactionmapping.html#cfn-quicksight-dlpsetting-labelactionmapping-labelname
	//
	LabelName *string `field:"required" json:"labelName" yaml:"labelName"`
}

