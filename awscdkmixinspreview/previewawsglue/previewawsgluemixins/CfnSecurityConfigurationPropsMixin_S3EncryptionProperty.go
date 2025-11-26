package previewawsgluemixins


// Specifies how Amazon Simple Storage Service (Amazon S3) data should be encrypted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3EncryptionProperty := &S3EncryptionProperty{
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	S3EncryptionMode: jsii.String("s3EncryptionMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-securityconfiguration-s3encryption.html
//
type CfnSecurityConfigurationPropsMixin_S3EncryptionProperty struct {
	// The Amazon Resource Name (ARN) of the KMS key to be used to encrypt the data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-securityconfiguration-s3encryption.html#cfn-glue-securityconfiguration-s3encryption-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The encryption mode to use for Amazon S3 data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-securityconfiguration-s3encryption.html#cfn-glue-securityconfiguration-s3encryption-s3encryptionmode
	//
	S3EncryptionMode *string `field:"optional" json:"s3EncryptionMode" yaml:"s3EncryptionMode"`
}

