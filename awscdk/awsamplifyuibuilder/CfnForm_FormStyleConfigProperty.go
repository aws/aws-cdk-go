package awsamplifyuibuilder


// The `FormStyleConfig` property specifies the configuration settings for the form's style properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   formStyleConfigProperty := &FormStyleConfigProperty{
//   	TokenReference: jsii.String("tokenReference"),
//   	Value: jsii.String("value"),
//   }
//
type CfnForm_FormStyleConfigProperty struct {
	// `CfnForm.FormStyleConfigProperty.TokenReference`.
	TokenReference *string `field:"optional" json:"tokenReference" yaml:"tokenReference"`
	// `CfnForm.FormStyleConfigProperty.Value`.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

