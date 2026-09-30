package awscloudfront


// A public key, its provider identity, and the field patterns CloudFront encrypts with that key.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionEntityProperty := &EncryptionEntityProperty{
//   	FieldPatterns: []*string{
//   		jsii.String("fieldPatterns"),
//   	},
//   	ProviderId: jsii.String("providerId"),
//   	PublicKeyId: jsii.String("publicKeyId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-fieldlevelencryptionprofile-encryptionentity.html
//
type CfnFieldLevelEncryptionProfile_EncryptionEntityProperty struct {
	// The request-body field names to encrypt.
	//
	// A pattern is either a full field name or leading characters followed by a wildcard (*). Patterns are case-sensitive and must not overlap.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-fieldlevelencryptionprofile-encryptionentity.html#cfn-cloudfront-fieldlevelencryptionprofile-encryptionentity-fieldpatterns
	//
	FieldPatterns *[]*string `field:"required" json:"fieldPatterns" yaml:"fieldPatterns"`
	// The provider associated with the public key.
	//
	// The same value must be supplied with the private key for an application to decrypt the data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-fieldlevelencryptionprofile-encryptionentity.html#cfn-cloudfront-fieldlevelencryptionprofile-encryptionentity-providerid
	//
	ProviderId *string `field:"required" json:"providerId" yaml:"providerId"`
	// The identifier of the CloudFront public key used to encrypt the fields that match the patterns.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-fieldlevelencryptionprofile-encryptionentity.html#cfn-cloudfront-fieldlevelencryptionprofile-encryptionentity-publickeyid
	//
	PublicKeyId *string `field:"required" json:"publicKeyId" yaml:"publicKeyId"`
}

