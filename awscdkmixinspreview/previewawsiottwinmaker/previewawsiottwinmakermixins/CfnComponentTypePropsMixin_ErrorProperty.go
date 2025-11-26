package previewawsiottwinmakermixins


// The component type error.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   errorProperty := &ErrorProperty{
//   	Code: jsii.String("code"),
//   	Message: jsii.String("message"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-componenttype-error.html
//
type CfnComponentTypePropsMixin_ErrorProperty struct {
	// The component type error code.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-componenttype-error.html#cfn-iottwinmaker-componenttype-error-code
	//
	Code *string `field:"optional" json:"code" yaml:"code"`
	// The component type error message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-componenttype-error.html#cfn-iottwinmaker-componenttype-error-message
	//
	Message *string `field:"optional" json:"message" yaml:"message"`
}

