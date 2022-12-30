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
//   cfnCertificateProps := &cfnCertificateProps{
//   	certificate: jsii.String("certificate"),
//   	usage: jsii.String("usage"),
//
//   	// the properties below are optional
//   	activeDate: jsii.String("activeDate"),
//   	certificateChain: jsii.String("certificateChain"),
//   	description: jsii.String("description"),
//   	inactiveDate: jsii.String("inactiveDate"),
//   	privateKey: jsii.String("privateKey"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
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

