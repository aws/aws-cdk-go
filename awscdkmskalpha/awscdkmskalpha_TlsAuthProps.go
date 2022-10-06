// The CDK Construct Library for AWS::MSK
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
//   cluster := msk.NewCluster(this, jsii.String("Cluster"), &clusterProps{
//   	clusterName: jsii.String("myCluster"),
//   	kafkaVersion: msk.kafkaVersion_V2_8_1(),
//   	vpc: vpc,
//   	encryptionInTransit: &encryptionInTransitConfig{
//   		clientBroker: msk.clientBrokerEncryption_TLS,
//   	},
//   	clientAuthentication: msk.clientAuthentication.tls(&tlsAuthProps{
//   		certificateAuthorities: []iCertificateAuthority{
//   			acmpca.certificateAuthority.fromCertificateAuthorityArn(this, jsii.String("CertificateAuthority"), jsii.String("arn:aws:acm-pca:us-west-2:1234567890:certificate-authority/11111111-1111-1111-1111-111111111111")),
//   		},
//   	}),
//   })
//
// Experimental.
type TlsAuthProps struct {
	// List of ACM Certificate Authorities to enable TLS authentication.
	// Experimental.
	CertificateAuthorities *[]awsacmpca.ICertificateAuthority `field:"optional" json:"certificateAuthorities" yaml:"certificateAuthorities"`
}

