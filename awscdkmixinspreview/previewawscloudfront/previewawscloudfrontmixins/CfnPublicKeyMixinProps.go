package previewawscloudfrontmixins


// Properties for CfnPublicKeyPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPublicKeyMixinProps := &CfnPublicKeyMixinProps{
//   	PublicKeyConfig: &PublicKeyConfigProperty{
//   		CallerReference: jsii.String("callerReference"),
//   		Comment: jsii.String("comment"),
//   		EncodedKey: jsii.String("encodedKey"),
//   		Name: jsii.String("name"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-publickey.html
//
type CfnPublicKeyMixinProps struct {
	// Configuration information about a public key that you can use with [signed URLs and signed cookies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html) , or with [field-level encryption](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/field-level-encryption.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-publickey.html#cfn-cloudfront-publickey-publickeyconfig
	//
	PublicKeyConfig interface{} `field:"optional" json:"publicKeyConfig" yaml:"publicKeyConfig"`
}

