package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sheetLayoutGroupMemberProperty := &SheetLayoutGroupMemberProperty{
//   	Id: jsii.String("id"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetlayoutgroupmember.html
//
type CfnAnalysis_SheetLayoutGroupMemberProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetlayoutgroupmember.html#cfn-quicksight-analysis-sheetlayoutgroupmember-id
	//
	Id *string `field:"required" json:"id" yaml:"id"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetlayoutgroupmember.html#cfn-quicksight-analysis-sheetlayoutgroupmember-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}

