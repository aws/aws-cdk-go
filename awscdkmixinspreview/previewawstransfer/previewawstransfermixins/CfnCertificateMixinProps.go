package previewawstransfermixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnCertificatePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCertificateMixinProps := &CfnCertificateMixinProps{
//   	ActiveDate: jsii.String("activeDate"),
//   	Certificate: jsii.String("certificate"),
//   	CertificateChain: jsii.String("certificateChain"),
//   	Description: jsii.String("description"),
//   	InactiveDate: jsii.String("inactiveDate"),
//   	PrivateKey: jsii.String("privateKey"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Usage: jsii.String("usage"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-certificate.html
//
type CfnCertificateMixinProps struct {
	// An optional date that specifies when the certificate becomes active.
	//
	// If you do not specify a value, `ActiveDate` takes the same value as `NotBeforeDate` , which is specified by the CA.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-certificate.html#cfn-transfer-certificate-activedate
	//
	ActiveDate *string `field:"optional" json:"activeDate" yaml:"activeDate"`
	// The file name for the certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-certificate.html#cfn-transfer-certificate-certificate
	//
	Certificate *string `field:"optional" json:"certificate" yaml:"certificate"`
	// The list of certificates that make up the chain for the certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-certificate.html#cfn-transfer-certificate-certificatechain
	//
	CertificateChain *string `field:"optional" json:"certificateChain" yaml:"certificateChain"`
	// The name or description that's used to identity the certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-certificate.html#cfn-transfer-certificate-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An optional date that specifies when the certificate becomes inactive.
	//
	// If you do not specify a value, `InactiveDate` takes the same value as `NotAfterDate` , which is specified by the CA.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-certificate.html#cfn-transfer-certificate-inactivedate
	//
	InactiveDate *string `field:"optional" json:"inactiveDate" yaml:"inactiveDate"`
	// The file that contains the private key for the certificate that's being imported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-certificate.html#cfn-transfer-certificate-privatekey
	//
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// Key-value pairs that can be used to group and search for certificates.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-certificate.html#cfn-transfer-certificate-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Specifies how this certificate is used. It can be used in the following ways:.
	//
	// - `SIGNING` : For signing AS2 messages
	// - `ENCRYPTION` : For encrypting AS2 messages
	// - `TLS` : For securing AS2 communications sent over HTTPS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-certificate.html#cfn-transfer-certificate-usage
	//
	Usage *string `field:"optional" json:"usage" yaml:"usage"`
}

