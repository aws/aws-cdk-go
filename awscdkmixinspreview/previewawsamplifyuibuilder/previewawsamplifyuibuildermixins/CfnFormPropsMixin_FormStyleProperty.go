package previewawsamplifyuibuildermixins


// The `FormStyle` property specifies the configuration for the form's style.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   formStyleProperty := &FormStyleProperty{
//   	HorizontalGap: &FormStyleConfigProperty{
//   		TokenReference: jsii.String("tokenReference"),
//   		Value: jsii.String("value"),
//   	},
//   	OuterPadding: &FormStyleConfigProperty{
//   		TokenReference: jsii.String("tokenReference"),
//   		Value: jsii.String("value"),
//   	},
//   	VerticalGap: &FormStyleConfigProperty{
//   		TokenReference: jsii.String("tokenReference"),
//   		Value: jsii.String("value"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-formstyle.html
//
type CfnFormPropsMixin_FormStyleProperty struct {
	// The spacing for the horizontal gap.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-formstyle.html#cfn-amplifyuibuilder-form-formstyle-horizontalgap
	//
	HorizontalGap interface{} `field:"optional" json:"horizontalGap" yaml:"horizontalGap"`
	// The size of the outer padding for the form.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-formstyle.html#cfn-amplifyuibuilder-form-formstyle-outerpadding
	//
	OuterPadding interface{} `field:"optional" json:"outerPadding" yaml:"outerPadding"`
	// The spacing for the vertical gap.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-formstyle.html#cfn-amplifyuibuilder-form-formstyle-verticalgap
	//
	VerticalGap interface{} `field:"optional" json:"verticalGap" yaml:"verticalGap"`
}

