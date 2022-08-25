package awsappmesh

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a TLS certificate that is supported for mutual TLS authentication.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var mesh mesh
//
//
//   node1 := appmesh.NewVirtualNode(this, jsii.String("node1"), &virtualNodeProps{
//   	mesh: mesh,
//   	serviceDiscovery: appmesh.serviceDiscovery.dns(jsii.String("node")),
//   	listeners: []virtualNodeListener{
//   		appmesh.*virtualNodeListener.grpc(&grpcVirtualNodeListenerOptions{
//   			port: jsii.Number(80),
//   			tls: &listenerTlsOptions{
//   				mode: appmesh.tlsMode_STRICT,
//   				certificate: appmesh.tlsCertificate.file(jsii.String("path/to/certChain"), jsii.String("path/to/privateKey")),
//   				// Validate a file client certificates to enable mutual TLS authentication when a client provides a certificate.
//   				mutualTlsValidation: &mutualTlsValidation{
//   					trust: appmesh.tlsValidationTrust.file(jsii.String("path-to-certificate")),
//   				},
//   			},
//   		}),
//   	},
//   })
//
//   certificateAuthorityArn := "arn:aws:acm-pca:us-east-1:123456789012:certificate-authority/12345678-1234-1234-1234-123456789012"
//   node2 := appmesh.NewVirtualNode(this, jsii.String("node2"), &virtualNodeProps{
//   	mesh: mesh,
//   	serviceDiscovery: appmesh.*serviceDiscovery.dns(jsii.String("node2")),
//   	backendDefaults: &backendDefaults{
//   		tlsClientPolicy: &tlsClientPolicy{
//   			ports: []*f64{
//   				jsii.Number(8080),
//   				jsii.Number(8081),
//   			},
//   			validation: &tlsValidation{
//   				subjectAlternativeNames: appmesh.subjectAlternativeNames.matchingExactly(jsii.String("mesh-endpoint.apps.local")),
//   				trust: appmesh.*tlsValidationTrust.acm([]iCertificateAuthority{
//   					acmpca.certificateAuthority.fromCertificateAuthorityArn(this, jsii.String("certificate"), certificateAuthorityArn),
//   				}),
//   			},
//   			// Provide a SDS client certificate when a server requests it and enable mutual TLS authentication.
//   			mutualTlsCertificate: appmesh.*tlsCertificate.sds(jsii.String("secret_certificate")),
//   		},
//   	},
//   })
//
type MutualTlsCertificate interface {
	TlsCertificate
	Differentiator() *bool
	// Returns TLS certificate based provider.
	Bind(_scope constructs.Construct) *TlsCertificateConfig
}

// The jsii proxy struct for MutualTlsCertificate
type jsiiProxy_MutualTlsCertificate struct {
	jsiiProxy_TlsCertificate
}

func (j *jsiiProxy_MutualTlsCertificate) Differentiator() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"differentiator",
		&returns,
	)
	return returns
}


func NewMutualTlsCertificate_Override(m MutualTlsCertificate) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.MutualTlsCertificate",
		nil, // no parameters
		m,
	)
}

// Returns an ACM TLS Certificate.
func MutualTlsCertificate_Acm(certificate awscertificatemanager.ICertificate) TlsCertificate {
	_init_.Initialize()

	var returns TlsCertificate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.MutualTlsCertificate",
		"acm",
		[]interface{}{certificate},
		&returns,
	)

	return returns
}

// Returns an File TLS Certificate.
func MutualTlsCertificate_File(certificateChainPath *string, privateKeyPath *string) MutualTlsCertificate {
	_init_.Initialize()

	var returns MutualTlsCertificate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.MutualTlsCertificate",
		"file",
		[]interface{}{certificateChainPath, privateKeyPath},
		&returns,
	)

	return returns
}

// Returns an SDS TLS Certificate.
func MutualTlsCertificate_Sds(secretName *string) MutualTlsCertificate {
	_init_.Initialize()

	var returns MutualTlsCertificate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.MutualTlsCertificate",
		"sds",
		[]interface{}{secretName},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MutualTlsCertificate) Bind(_scope constructs.Construct) *TlsCertificateConfig {
	var returns *TlsCertificateConfig

	_jsii_.Invoke(
		m,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

