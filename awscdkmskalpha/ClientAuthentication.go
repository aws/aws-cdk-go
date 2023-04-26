// The CDK Construct Library for AWS::MSK
package awscdkmskalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmskalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Configuration properties for client authentication.
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
type ClientAuthentication interface {
	// - properties for SASL authentication.
	// Experimental.
	SaslProps() *SaslAuthProps
	// - properties for TLS authentication.
	// Experimental.
	TlsProps() *TlsAuthProps
}

// The jsii proxy struct for ClientAuthentication
type jsiiProxy_ClientAuthentication struct {
	_ byte // padding
}

func (j *jsiiProxy_ClientAuthentication) SaslProps() *SaslAuthProps {
	var returns *SaslAuthProps
	_jsii_.Get(
		j,
		"saslProps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClientAuthentication) TlsProps() *TlsAuthProps {
	var returns *TlsAuthProps
	_jsii_.Get(
		j,
		"tlsProps",
		&returns,
	)
	return returns
}


// SASL authentication.
// Experimental.
func ClientAuthentication_Sasl(props *SaslAuthProps) ClientAuthentication {
	_init_.Initialize()

	if err := validateClientAuthentication_SaslParameters(props); err != nil {
		panic(err)
	}
	var returns ClientAuthentication

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-msk-alpha.ClientAuthentication",
		"sasl",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// SASL + TLS authentication.
// Experimental.
func ClientAuthentication_SaslTls(saslTlsProps *SaslTlsAuthProps) ClientAuthentication {
	_init_.Initialize()

	if err := validateClientAuthentication_SaslTlsParameters(saslTlsProps); err != nil {
		panic(err)
	}
	var returns ClientAuthentication

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-msk-alpha.ClientAuthentication",
		"saslTls",
		[]interface{}{saslTlsProps},
		&returns,
	)

	return returns
}

// TLS authentication.
// Experimental.
func ClientAuthentication_Tls(props *TlsAuthProps) ClientAuthentication {
	_init_.Initialize()

	if err := validateClientAuthentication_TlsParameters(props); err != nil {
		panic(err)
	}
	var returns ClientAuthentication

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-msk-alpha.ClientAuthentication",
		"tls",
		[]interface{}{props},
		&returns,
	)

	return returns
}

