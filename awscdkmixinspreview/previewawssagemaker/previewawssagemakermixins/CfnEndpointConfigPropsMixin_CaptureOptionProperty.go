package previewawssagemakermixins


// Specifies whether the endpoint captures input data or output data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   captureOptionProperty := &CaptureOptionProperty{
//   	CaptureMode: jsii.String("captureMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-captureoption.html
//
type CfnEndpointConfigPropsMixin_CaptureOptionProperty struct {
	// Specifies whether the endpoint captures input data or output data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-captureoption.html#cfn-sagemaker-endpointconfig-captureoption-capturemode
	//
	CaptureMode *string `field:"optional" json:"captureMode" yaml:"captureMode"`
}

