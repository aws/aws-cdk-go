package awsamplifyuibuilder


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   formCTAProperty := &formCTAProperty{
//   	cancel: &formButtonProperty{
//   		children: jsii.String("children"),
//   		excluded: jsii.Boolean(false),
//   		position: &fieldPositionProperty{
//   			below: jsii.String("below"),
//   			fixed: jsii.String("fixed"),
//   			rightOf: jsii.String("rightOf"),
//   		},
//   	},
//   	clear: &formButtonProperty{
//   		children: jsii.String("children"),
//   		excluded: jsii.Boolean(false),
//   		position: &fieldPositionProperty{
//   			below: jsii.String("below"),
//   			fixed: jsii.String("fixed"),
//   			rightOf: jsii.String("rightOf"),
//   		},
//   	},
//   	position: jsii.String("position"),
//   	submit: &formButtonProperty{
//   		children: jsii.String("children"),
//   		excluded: jsii.Boolean(false),
//   		position: &fieldPositionProperty{
//   			below: jsii.String("below"),
//   			fixed: jsii.String("fixed"),
//   			rightOf: jsii.String("rightOf"),
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

