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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-formstyleconfig.html
//
type CfnForm_FormStyleConfigProperty struct {
	// A reference to a design token to use to bind the form's style properties to an existing theme.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-formstyleconfig.html#cfn-amplifyuibuilder-form-formstyleconfig-tokenreference
	//
	TokenReference *string `field:"optional" json:"tokenReference" yaml:"tokenReference"`
	// The value of the style setting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-formstyleconfig.html#cfn-amplifyuibuilder-form-formstyleconfig-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

