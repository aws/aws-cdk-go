package awsamplifyuibuilder


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   formCTAProperty := &FormCTAProperty{
//   	Cancel: &FormButtonProperty{
//   		Children: jsii.String("children"),
//   		Excluded: jsii.Boolean(false),
//   		Position: &FieldPositionProperty{
//   			Below: jsii.String("below"),
//   			Fixed: jsii.String("fixed"),
//   			RightOf: jsii.String("rightOf"),
//   		},
//   	},
//   	Clear: &FormButtonProperty{
//   		Children: jsii.String("children"),
//   		Excluded: jsii.Boolean(false),
//   		Position: &FieldPositionProperty{
//   			Below: jsii.String("below"),
//   			Fixed: jsii.String("fixed"),
//   			RightOf: jsii.String("rightOf"),
//   		},
//   	},
//   	Position: jsii.String("position"),
//   	Submit: &FormButtonProperty{
//   		Children: jsii.String("children"),
//   		Excluded: jsii.Boolean(false),
//   		Position: &FieldPositionProperty{
//   			Below: jsii.String("below"),
//   			Fixed: jsii.String("fixed"),
//   			RightOf: jsii.String("rightOf"),
//   		},
//   	},
//   }
//
type CfnForm_FormCTAProperty struct {
	// `CfnForm.FormCTAProperty.Cancel`.
	Cancel interface{} `field:"optional" json:"cancel" yaml:"cancel"`
	// `CfnForm.FormCTAProperty.Clear`.
	Clear interface{} `field:"optional" json:"clear" yaml:"clear"`
	// `CfnForm.FormCTAProperty.Position`.
	Position *string `field:"optional" json:"position" yaml:"position"`
	// `CfnForm.FormCTAProperty.Submit`.
	Submit interface{} `field:"optional" json:"submit" yaml:"submit"`
}

