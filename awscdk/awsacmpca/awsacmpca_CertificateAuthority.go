package awsacmpca

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Defines a Certificate for ACMPCA.
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
type CertificateAuthority interface {
}

// The jsii proxy struct for CertificateAuthority
type jsiiProxy_CertificateAuthority struct {
	_ byte // padding
}

// Import an existing Certificate given an ARN.
func CertificateAuthority_FromCertificateAuthorityArn(scope constructs.Construct, id *string, certificateAuthorityArn *string) ICertificateAuthority {
	_init_.Initialize()

	if err := validateCertificateAuthority_FromCertificateAuthorityArnParameters(scope, id, certificateAuthorityArn); err != nil {
		panic(err)
	}
	var returns ICertificateAuthority

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_acmpca.CertificateAuthority",
		"fromCertificateAuthorityArn",
		[]interface{}{scope, id, certificateAuthorityArn},
		&returns,
	)

	return returns
}

