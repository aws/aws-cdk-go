package awstransfer

import (
	"github.com/aws/aws-cdk-go/awscdk"
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
type CfnCertificateProps struct {
	// The file name for the certificate.
	Certificate *string `field:"required" json:"certificate" yaml:"certificate"`
	// Specifies whether this certificate is used for signing or encryption.
	Usage *string `field:"required" json:"usage" yaml:"usage"`
	// An optional date that specifies when the certificate becomes active.
	ActiveDate *string `field:"optional" json:"activeDate" yaml:"activeDate"`
	// The list of certificates that make up the chain for the certificate.
	CertificateChain *string `field:"optional" json:"certificateChain" yaml:"certificateChain"`
	// The name or description that's used to identity the certificate.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An optional date that specifies when the certificate becomes inactive.
	InactiveDate *string `field:"optional" json:"inactiveDate" yaml:"inactiveDate"`
	// The file that contains the private key for the certificate that's being imported.
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// Key-value pairs that can be used to group and search for certificates.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

