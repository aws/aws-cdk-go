package awstranslate


// The encryption key used to encrypt this object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionKeyProperty := &EncryptionKeyProperty{
//   	Id: jsii.String("id"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-translate-paralleldata-encryptionkey.html
//
type CfnParallelData_EncryptionKeyProperty struct {
	// The Amazon Resource Name (ARN) of the encryption key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-translate-paralleldata-encryptionkey.html#cfn-translate-paralleldata-encryptionkey-id
	//
	Id *string `field:"required" json:"id" yaml:"id"`
	// The type of encryption key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-translate-paralleldata-encryptionkey.html#cfn-translate-paralleldata-encryptionkey-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}

