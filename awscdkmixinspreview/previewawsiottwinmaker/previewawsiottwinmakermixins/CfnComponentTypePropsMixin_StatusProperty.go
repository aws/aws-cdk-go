package previewawsiottwinmakermixins


// The component type status.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   statusProperty := &StatusProperty{
//   	Error: &ErrorProperty{
//   		Code: jsii.String("code"),
//   		Message: jsii.String("message"),
//   	},
//   	State: jsii.String("state"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-componenttype-status.html
//
type CfnComponentTypePropsMixin_StatusProperty struct {
	// The component type error.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-componenttype-status.html#cfn-iottwinmaker-componenttype-status-error
	//
	Error interface{} `field:"optional" json:"error" yaml:"error"`
	// The component type status state.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-componenttype-status.html#cfn-iottwinmaker-componenttype-status-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
}

