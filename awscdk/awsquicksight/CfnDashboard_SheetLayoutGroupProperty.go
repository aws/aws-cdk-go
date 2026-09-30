package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var type interface{}
//
//   sheetLayoutGroupProperty := &SheetLayoutGroupProperty{
//   	Id: jsii.String("id"),
//   	Members: []interface{}{
//   		&SheetLayoutGroupMemberProperty{
//   			Id: jsii.String("id"),
//   			Type: type,
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sheetlayoutgroup.html
//
type CfnDashboard_SheetLayoutGroupProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sheetlayoutgroup.html#cfn-quicksight-dashboard-sheetlayoutgroup-id
	//
	Id *string `field:"required" json:"id" yaml:"id"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sheetlayoutgroup.html#cfn-quicksight-dashboard-sheetlayoutgroup-members
	//
	Members interface{} `field:"required" json:"members" yaml:"members"`
}

