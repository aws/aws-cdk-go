package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var type interface{}
//
//   sheetLayoutGroupMemberProperty := &SheetLayoutGroupMemberProperty{
//   	Id: jsii.String("id"),
//   	Type: type,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sheetlayoutgroupmember.html
//
type CfnDashboardPropsMixin_SheetLayoutGroupMemberProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sheetlayoutgroupmember.html#cfn-quicksight-dashboard-sheetlayoutgroupmember-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sheetlayoutgroupmember.html#cfn-quicksight-dashboard-sheetlayoutgroupmember-type
	//
	Type interface{} `field:"optional" json:"type" yaml:"type"`
}

