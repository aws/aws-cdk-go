package awsamplifyuibuilder


// Describes the call to action button configuration for the form.
//
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
	// Displays a cancel button.
	Cancel interface{} `field:"optional" json:"cancel" yaml:"cancel"`
	// Displays a clear button.
	Clear interface{} `field:"optional" json:"clear" yaml:"clear"`
	// The position of the button.
	Position *string `field:"optional" json:"position" yaml:"position"`
	// Displays a submit button.
	Submit interface{} `field:"optional" json:"submit" yaml:"submit"`
}

