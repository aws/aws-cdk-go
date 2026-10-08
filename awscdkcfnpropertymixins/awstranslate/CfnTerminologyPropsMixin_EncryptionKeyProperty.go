package awstranslate


// The encryption key for the custom terminology.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   encryptionKeyProperty := &EncryptionKeyProperty{
//   	Id: jsii.String("id"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-translate-terminology-encryptionkey.html
//
type CfnTerminologyPropsMixin_EncryptionKeyProperty struct {
	// The Amazon Resource Name (ARN) of the encryption key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-translate-terminology-encryptionkey.html#cfn-translate-terminology-encryptionkey-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The type of encryption key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-translate-terminology-encryptionkey.html#cfn-translate-terminology-encryptionkey-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

