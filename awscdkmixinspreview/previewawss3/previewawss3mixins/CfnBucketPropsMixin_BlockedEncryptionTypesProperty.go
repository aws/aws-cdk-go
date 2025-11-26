package previewawss3mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   blockedEncryptionTypesProperty := &BlockedEncryptionTypesProperty{
//   	EncryptionType: []*string{
//   		jsii.String("encryptionType"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-blockedencryptiontypes.html
//
type CfnBucketPropsMixin_BlockedEncryptionTypesProperty struct {
	// List of encryption types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-blockedencryptiontypes.html#cfn-s3-bucket-blockedencryptiontypes-encryptiontype
	//
	EncryptionType *[]*string `field:"optional" json:"encryptionType" yaml:"encryptionType"`
}

