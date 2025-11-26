package previewawssammixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   keySAMPTProperty := &KeySAMPTProperty{
//   	KeyId: jsii.String("keyId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-keysampt.html
//
type CfnFunctionPropsMixin_KeySAMPTProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-keysampt.html#cfn-serverless-function-keysampt-keyid
	//
	KeyId *string `field:"optional" json:"keyId" yaml:"keyId"`
}

