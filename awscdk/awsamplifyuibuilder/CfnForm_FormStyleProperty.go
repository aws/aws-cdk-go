package awsamplifyuibuilder


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
type CfnForm_FormStyleProperty struct {
	// `CfnForm.FormStyleProperty.HorizontalGap`.
	HorizontalGap interface{} `field:"optional" json:"horizontalGap" yaml:"horizontalGap"`
	// `CfnForm.FormStyleProperty.OuterPadding`.
	OuterPadding interface{} `field:"optional" json:"outerPadding" yaml:"outerPadding"`
	// `CfnForm.FormStyleProperty.VerticalGap`.
	VerticalGap interface{} `field:"optional" json:"verticalGap" yaml:"verticalGap"`
}

