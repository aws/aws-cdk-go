package awsamplifyuibuilder


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
	// `CfnForm.FormStyleProperty.HorizontalGap`.
	HorizontalGap interface{} `field:"optional" json:"horizontalGap" yaml:"horizontalGap"`
	// `CfnForm.FormStyleProperty.OuterPadding`.
	OuterPadding interface{} `field:"optional" json:"outerPadding" yaml:"outerPadding"`
	// `CfnForm.FormStyleProperty.VerticalGap`.
	VerticalGap interface{} `field:"optional" json:"verticalGap" yaml:"verticalGap"`
}

