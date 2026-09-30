package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetlayoutgroup.html
//
type CfnTemplatePropsMixin_SheetLayoutGroupProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetlayoutgroup.html#cfn-quicksight-template-sheetlayoutgroup-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetlayoutgroup.html#cfn-quicksight-template-sheetlayoutgroup-members
	//
	Members interface{} `field:"optional" json:"members" yaml:"members"`
}

