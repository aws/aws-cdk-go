package awsiam


// The private key metadata for the SAML provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sAMLPrivateKeyProperty := &SAMLPrivateKeyProperty{
//   	KeyId: jsii.String("keyId"),
//   	Timestamp: jsii.String("timestamp"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-samlprovider-samlprivatekey.html
//
type CfnSAMLProvider_SAMLPrivateKeyProperty struct {
	// The unique identifier for the SAML private key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-samlprovider-samlprivatekey.html#cfn-iam-samlprovider-samlprivatekey-keyid
	//
	KeyId *string `field:"required" json:"keyId" yaml:"keyId"`
	// The date and time, in <a href=\"http://www.iso.org/iso/iso8601\">ISO 8601 date-time </a> format, when the private key was uploaded.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-samlprovider-samlprivatekey.html#cfn-iam-samlprovider-samlprivatekey-timestamp
	//
	Timestamp *string `field:"required" json:"timestamp" yaml:"timestamp"`
}

