package awsappmesh

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsacmpca"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a TLS Validation Context Trust that is supported for mutual TLS authentication.
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
type MutualTlsValidationTrust interface {
	TlsValidationTrust
	Differentiator() *bool
	// Returns Trust context based on trust type.
	Bind(scope constructs.Construct) *TlsValidationTrustConfig
}

// The jsii proxy struct for MutualTlsValidationTrust
type jsiiProxy_MutualTlsValidationTrust struct {
	jsiiProxy_TlsValidationTrust
}

func (j *jsiiProxy_MutualTlsValidationTrust) Differentiator() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"differentiator",
		&returns,
	)
	return returns
}


func NewMutualTlsValidationTrust_Override(m MutualTlsValidationTrust) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.MutualTlsValidationTrust",
		nil, // no parameters
		m,
	)
}

// TLS Validation Context Trust for ACM Private Certificate Authority (CA).
func MutualTlsValidationTrust_Acm(certificateAuthorities *[]awsacmpca.ICertificateAuthority) TlsValidationTrust {
	_init_.Initialize()

	if err := validateMutualTlsValidationTrust_AcmParameters(certificateAuthorities); err != nil {
		panic(err)
	}
	var returns TlsValidationTrust

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.MutualTlsValidationTrust",
		"acm",
		[]interface{}{certificateAuthorities},
		&returns,
	)

	return returns
}

// Tells envoy where to fetch the validation context from.
func MutualTlsValidationTrust_File(certificateChain *string) MutualTlsValidationTrust {
	_init_.Initialize()

	if err := validateMutualTlsValidationTrust_FileParameters(certificateChain); err != nil {
		panic(err)
	}
	var returns MutualTlsValidationTrust

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.MutualTlsValidationTrust",
		"file",
		[]interface{}{certificateChain},
		&returns,
	)

	return returns
}

// TLS Validation Context Trust for Envoy' service discovery service.
func MutualTlsValidationTrust_Sds(secretName *string) MutualTlsValidationTrust {
	_init_.Initialize()

	if err := validateMutualTlsValidationTrust_SdsParameters(secretName); err != nil {
		panic(err)
	}
	var returns MutualTlsValidationTrust

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.MutualTlsValidationTrust",
		"sds",
		[]interface{}{secretName},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MutualTlsValidationTrust) Bind(scope constructs.Construct) *TlsValidationTrustConfig {
	if err := m.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *TlsValidationTrustConfig

	_jsii_.Invoke(
		m,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

