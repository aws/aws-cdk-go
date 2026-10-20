package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sheetLayoutGroupProperty := &SheetLayoutGroupProperty{
//   	Id: jsii.String("id"),
//   	Members: []interface{}{
//   		&SheetLayoutGroupMemberProperty{
//   			Id: jsii.String("id"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetlayoutgroup.html
//
type CfnTemplate_SheetLayoutGroupProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetlayoutgroup.html#cfn-quicksight-template-sheetlayoutgroup-id
	//
	Id *string `field:"required" json:"id" yaml:"id"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetlayoutgroup.html#cfn-quicksight-template-sheetlayoutgroup-members
	//
	Members interface{} `field:"required" json:"members" yaml:"members"`
}

