package awstransfer

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCertificate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCertificateProps := &CfnCertificateProps{
//   	Certificate: jsii.String("certificate"),
//   	Usage: jsii.String("usage"),
//
//   	// the properties below are optional
//   	ActiveDate: jsii.String("activeDate"),
//   	CertificateChain: jsii.String("certificateChain"),
//   	Description: jsii.String("description"),
//   	InactiveDate: jsii.String("inactiveDate"),
//   	PrivateKey: jsii.String("privateKey"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-certificate.html
//
type CfnCertificateProps struct {
	// The file name for the certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-certificate.html#cfn-transfer-certificate-certificate
	//
	Certificate *string `field:"required" json:"certificate" yaml:"certificate"`
	// Specifies whether this certificate is used for signing or encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-certificate.html#cfn-transfer-certificate-usage
	//
	Usage *string `field:"required" json:"usage" yaml:"usage"`
	// An optional date that specifies when the certificate becomes active.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-certificate.html#cfn-transfer-certificate-activedate
	//
	ActiveDate *string `field:"optional" json:"activeDate" yaml:"activeDate"`
	// The list of certificates that make up the chain for the certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-certificate.html#cfn-transfer-certificate-certificatechain
	//
	CertificateChain *string `field:"optional" json:"certificateChain" yaml:"certificateChain"`
	// The name or description that's used to identity the certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-certificate.html#cfn-transfer-certificate-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An optional date that specifies when the certificate becomes inactive.
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
}

