package previewawsiotmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   commandPayloadProperty := &CommandPayloadProperty{
//   	Content: jsii.String("content"),
//   	ContentType: jsii.String("contentType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-command-commandpayload.html
//
type CfnCommandPropsMixin_CommandPayloadProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-command-commandpayload.html#cfn-iot-command-commandpayload-content
	//
	Content *string `field:"optional" json:"content" yaml:"content"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-command-commandpayload.html#cfn-iot-command-commandpayload-contenttype
	//
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
}

