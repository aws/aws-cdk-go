package awsamplifyuibuilder


// Describes the configuration settings for the form's style properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   formStyleConfigProperty := &formStyleConfigProperty{
//   	tokenReference: jsii.String("tokenReference"),
//   	value: jsii.String("value"),
//   }
//
type CfnForm_FormStyleConfigProperty struct {
	// A reference to a design token to use to bind the form's style properties to an existing theme.
	TokenReference *string `field:"optional" json:"tokenReference" yaml:"tokenReference"`
	// The value of the style setting.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

