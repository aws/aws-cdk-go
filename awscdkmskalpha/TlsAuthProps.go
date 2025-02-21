package awscdkmskalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsacmpca"
)

// TLS authentication properties.
//
// Example:
//   import acmpca "github.com/aws/aws-cdk-go/awscdk"
//
//   var vpc vpc
//
//   cluster := msk.NewCluster(this, jsii.String("Cluster"), &ClusterProps{
//   	ClusterName: jsii.String("myCluster"),
//   	KafkaVersion: msk.KafkaVersion_V2_8_1(),
//   	Vpc: Vpc,
//   	EncryptionInTransit: &EncryptionInTransitConfig{
//   		ClientBroker: msk.ClientBrokerEncryption_TLS,
//   	},
//   	ClientAuthentication: msk.ClientAuthentication_Tls(&TlsAuthProps{
//   		CertificateAuthorities: []iCertificateAuthority{
//   			acmpca.CertificateAuthority_FromCertificateAuthorityArn(this, jsii.String("CertificateAuthority"), jsii.String("arn:aws:acm-pca:us-west-2:1234567890:certificate-authority/11111111-1111-1111-1111-111111111111")),
//   		},
//   	}),
//   })
//
// Experimental.
type TlsAuthProps struct {
	// List of ACM Certificate Authorities to enable TLS authentication.
	// Default: - none.
	//
	// Experimental.
	CertificateAuthorities *[]awsacmpca.ICertificateAuthority `field:"optional" json:"certificateAuthorities" yaml:"certificateAuthorities"`
}

