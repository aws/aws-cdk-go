package mixinsawsiam


// Contains the private keys for the SAML provider.
//
// This data type is used as a response element in the [GetSAMLProvider](https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetSAMLProvider.html) operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sAMLPrivateKeyProperty := &SAMLPrivateKeyProperty{
//   	KeyId: jsii.String("keyId"),
//   	Timestamp: jsii.String("timestamp"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-samlprovider-samlprivatekey.html
//
type CfnSAMLProviderPropsMixin_SAMLPrivateKeyProperty struct {
	// The unique identifier for the SAML private key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-samlprovider-samlprivatekey.html#cfn-iam-samlprovider-samlprivatekey-keyid
	//
	KeyId *string `field:"optional" json:"keyId" yaml:"keyId"`
	// The date and time, in [ISO 8601 date-time](https://docs.aws.amazon.com/http://www.iso.org/iso/iso8601) format, when the private key was uploaded.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-samlprovider-samlprivatekey.html#cfn-iam-samlprovider-samlprivatekey-timestamp
	//
	Timestamp *string `field:"optional" json:"timestamp" yaml:"timestamp"`
}

