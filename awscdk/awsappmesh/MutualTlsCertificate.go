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
//   var mesh mesh
//
//
//   node1 := appmesh.NewVirtualNode(this, jsii.String("node1"), &VirtualNodeProps{
//   	Mesh: Mesh,
//   	ServiceDiscovery: appmesh.ServiceDiscovery_Dns(jsii.String("node")),
//   	Listeners: []virtualNodeListener{
//   		appmesh.*virtualNodeListener_Grpc(&GrpcVirtualNodeListenerOptions{
//   			Port: jsii.Number(80),
//   			Tls: &ListenerTlsOptions{
//   				Mode: appmesh.TlsMode_STRICT,
//   				Certificate: appmesh.TlsCertificate_File(jsii.String("path/to/certChain"), jsii.String("path/to/privateKey")),
//   				// Validate a file client certificates to enable mutual TLS authentication when a client provides a certificate.
//   				MutualTlsValidation: &MutualTlsValidation{
//   					Trust: appmesh.TlsValidationTrust_File(jsii.String("path-to-certificate")),
//   				},
//   			},
//   		}),
//   	},
//   })
//
//   certificateAuthorityArn := "arn:aws:acm-pca:us-east-1:123456789012:certificate-authority/12345678-1234-1234-1234-123456789012"
//   node2 := appmesh.NewVirtualNode(this, jsii.String("node2"), &VirtualNodeProps{
//   	Mesh: Mesh,
//   	ServiceDiscovery: appmesh.ServiceDiscovery_*Dns(jsii.String("node2")),
//   	BackendDefaults: &BackendDefaults{
//   		TlsClientPolicy: &TlsClientPolicy{
//   			Ports: []*f64{
//   				jsii.Number(8080),
//   				jsii.Number(8081),
//   			},
//   			Validation: &TlsValidation{
//   				SubjectAlternativeNames: appmesh.SubjectAlternativeNames_MatchingExactly(jsii.String("mesh-endpoint.apps.local")),
//   				Trust: appmesh.TlsValidationTrust_Acm([]iCertificateAuthority{
//   					acmpca.CertificateAuthority_FromCertificateAuthorityArn(this, jsii.String("certificate"), certificateAuthorityArn),
//   				}),
//   			},
//   			// Provide a SDS client certificate when a server requests it and enable mutual TLS authentication.
//   			MutualTlsCertificate: appmesh.TlsCertificate_Sds(jsii.String("secret_certificate")),
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

	if err := validateMutualTlsCertificate_AcmParameters(certificate); err != nil {
		panic(err)
	}
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

	if err := validateMutualTlsCertificate_FileParameters(certificateChainPath, privateKeyPath); err != nil {
		panic(err)
	}
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

	if err := validateMutualTlsCertificate_SdsParameters(secretName); err != nil {
		panic(err)
	}
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
	if err := m.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *TlsCertificateConfig

	_jsii_.Invoke(
		m,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

