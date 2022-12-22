package awsamplifyuibuilder


// Describes the configuration for the form's style.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   formStyleProperty := &formStyleProperty{
//   	horizontalGap: &formStyleConfigProperty{
//   		tokenReference: jsii.String("tokenReference"),
//   		value: jsii.String("value"),
//   	},
//   	outerPadding: &formStyleConfigProperty{
//   		tokenReference: jsii.String("tokenReference"),
//   		value: jsii.String("value"),
//   	},
//   	verticalGap: &formStyleConfigProperty{
//   		tokenReference: jsii.String("tokenReference"),
//   		value: jsii.String("value"),
//   	},
//   }
//
type CfnForm_FormStyleProperty struct {
	// The spacing for the horizontal gap.
	HorizontalGap interface{} `field:"optional" json:"horizontalGap" yaml:"horizontalGap"`
	// The size of the outer padding for the form.
	OuterPadding interface{} `field:"optional" json:"outerPadding" yaml:"outerPadding"`
	// The spacing for the vertical gap.
	VerticalGap interface{} `field:"optional" json:"verticalGap" yaml:"verticalGap"`
}

