package mixinsawsamplifyuibuilder


// The `FormCTA` property specifies the call to action button configuration for the form.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-formcta.html
//
type CfnFormPropsMixin_FormCTAProperty struct {
	// Displays a cancel button.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-formcta.html#cfn-amplifyuibuilder-form-formcta-cancel
	//
	Cancel interface{} `field:"optional" json:"cancel" yaml:"cancel"`
	// Displays a clear button.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-formcta.html#cfn-amplifyuibuilder-form-formcta-clear
	//
	Clear interface{} `field:"optional" json:"clear" yaml:"clear"`
	// The position of the button.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-formcta.html#cfn-amplifyuibuilder-form-formcta-position
	//
	Position *string `field:"optional" json:"position" yaml:"position"`
	// Displays a submit button.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-formcta.html#cfn-amplifyuibuilder-form-formcta-submit
	//
	Submit interface{} `field:"optional" json:"submit" yaml:"submit"`
}

