package awscdkmskalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsacmpca"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskms"
)

// SASL + TLS authentication properties.
//
// Example:
//   import acmpca "github.com/aws/aws-cdk-go/awscdk"
//
//   var vpc Vpc
//
//   cluster := msk.NewCluster(this, jsii.String("Cluster"), &ClusterProps{
//   	ClusterName: jsii.String("myCluster"),
//   	KafkaVersion: msk.KafkaVersion_V4_1_X_KRAFT(),
//   	Vpc: Vpc,
//   	EncryptionInTransit: &EncryptionInTransitConfig{
//   		ClientBroker: msk.ClientBrokerEncryption_TLS,
//   	},
//   	ClientAuthentication: msk.ClientAuthentication_SaslTls(&SaslTlsAuthProps{
//   		Iam: jsii.Boolean(true),
//   		CertificateAuthorities: []ICertificateAuthority{
//   			acmpca.CertificateAuthority_FromCertificateAuthorityArn(this, jsii.String("CertificateAuthority"), jsii.String("arn:aws:acm-pca:us-west-2:1234567890:certificate-authority/11111111-1111-1111-1111-111111111111")),
//   		},
//   	}),
//   })
//
// Experimental.
type SaslTlsAuthProps struct {
	// Enable IAM access control.
	// Default: false.
	//
	// Experimental.
	Iam *bool `field:"optional" json:"iam" yaml:"iam"`
	// KMS Key to encrypt SASL/SCRAM secrets.
	//
	// You must use a customer master key (CMK) when creating users in secrets manager.
	// You cannot use a Secret with Amazon MSK that uses the default Secrets Manager encryption key.
	// Default: - CMK will be created with alias msk/{clusterName}/sasl/scram.
	//
	// Experimental.
	Key interfacesawskms.IKeyRef `field:"optional" json:"key" yaml:"key"`
	// Enable SASL/SCRAM authentication.
	// Default: false.
	//
	// Experimental.
	Scram *bool `field:"optional" json:"scram" yaml:"scram"`
	// List of ACM Certificate Authorities to enable TLS authentication.
	// Default: - none.
	//
	// Experimental.
	CertificateAuthorities *[]awsacmpca.ICertificateAuthority `field:"optional" json:"certificateAuthorities" yaml:"certificateAuthorities"`
}

