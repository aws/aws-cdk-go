package awsiot


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   codeSigningSignatureProperty := &CodeSigningSignatureProperty{
//   	InlineDocument: jsii.String("inlineDocument"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-codesigningsignature.html
//
type CfnOTAUpdatePropsMixin_CodeSigningSignatureProperty struct {
	// A base64 encoded binary representation of the code signing signature.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-codesigningsignature.html#cfn-iot-otaupdate-codesigningsignature-inlinedocument
	//
	InlineDocument *string `field:"optional" json:"inlineDocument" yaml:"inlineDocument"`
}

