package previewawsiottwinmakermixins


// The entity error.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-entity-error.html
//
type CfnEntityPropsMixin_ErrorProperty struct {
	// The entity error code.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-entity-error.html#cfn-iottwinmaker-entity-error-code
	//
	Code *string `field:"optional" json:"code" yaml:"code"`
	// The entity error message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-entity-error.html#cfn-iottwinmaker-entity-error-message
	//
	Message *string `field:"optional" json:"message" yaml:"message"`
}

