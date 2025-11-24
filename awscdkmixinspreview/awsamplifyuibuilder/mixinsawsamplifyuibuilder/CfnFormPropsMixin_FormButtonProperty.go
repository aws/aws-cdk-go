package mixinsawsamplifyuibuilder


// The `FormButton` property specifies the configuration for a button UI element that is a part of a form.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   formButtonProperty := &FormButtonProperty{
//   	Children: jsii.String("children"),
//   	Excluded: jsii.Boolean(false),
//   	Position: &FieldPositionProperty{
//   		Below: jsii.String("below"),
//   		Fixed: jsii.String("fixed"),
//   		RightOf: jsii.String("rightOf"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-formbutton.html
//
type CfnFormPropsMixin_FormButtonProperty struct {
	// Describes the button's properties.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-formbutton.html#cfn-amplifyuibuilder-form-formbutton-children
	//
	Children *string `field:"optional" json:"children" yaml:"children"`
	// Specifies whether the button is visible on the form.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-formbutton.html#cfn-amplifyuibuilder-form-formbutton-excluded
	//
	Excluded interface{} `field:"optional" json:"excluded" yaml:"excluded"`
	// The position of the button.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-formbutton.html#cfn-amplifyuibuilder-form-formbutton-position
	//
	Position interface{} `field:"optional" json:"position" yaml:"position"`
}

