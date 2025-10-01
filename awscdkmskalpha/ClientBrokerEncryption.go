package awscdkmskalpha


// Indicates the encryption setting for data in transit between clients and brokers.
//
// Example:
//   import acmpca "github.com/aws/aws-cdk-go/awscdk"
//
//   var vpc vpc
//
//   cluster := msk.NewCluster(this, jsii.String("Cluster"), &ClusterProps{
//   	ClusterName: jsii.String("myCluster"),
//   	KafkaVersion: msk.KafkaVersion_V4_0_X_KRAFT(),
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
type ClientBrokerEncryption string

const (
	// TLS means that client-broker communication is enabled with TLS only.
	// Experimental.
	ClientBrokerEncryption_TLS ClientBrokerEncryption = "TLS"
	// TLS_PLAINTEXT means that client-broker communication is enabled for both TLS-encrypted, as well as plaintext data.
	// Experimental.
	ClientBrokerEncryption_TLS_PLAINTEXT ClientBrokerEncryption = "TLS_PLAINTEXT"
	// PLAINTEXT means that client-broker communication is enabled in plaintext only.
	// Experimental.
	ClientBrokerEncryption_PLAINTEXT ClientBrokerEncryption = "PLAINTEXT"
)

