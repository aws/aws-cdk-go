package awsappmesh

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Used to generate Subject Alternative Names Matchers.
//
// Example:
//   var mesh Mesh
//
//
//   node1 := appmesh.NewVirtualNode(this, jsii.String("node1"), &VirtualNodeProps{
//   	Mesh: Mesh,
//   	ServiceDiscovery: appmesh.ServiceDiscovery_Dns(jsii.String("node")),
//   	Listeners: []VirtualNodeListener{
//   		appmesh.VirtualNodeListener_Grpc(&GrpcVirtualNodeListenerOptions{
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
//   				Trust: appmesh.TlsValidationTrust_Acm([]ICertificateAuthority{
//   					acmpca.CertificateAuthority_FromCertificateAuthorityArn(this, jsii.String("certificate"), certificateAuthorityArn),
//   				}),
//   			},
//   			// Provide a SDS client certificate when a server requests it and enable mutual TLS authentication.
//   			MutualTlsCertificate: appmesh.TlsCertificate_Sds(jsii.String("secret_certificate")),
//   		},
//   	},
//   })
//
type SubjectAlternativeNames interface {
	// Returns Subject Alternative Names Matcher based on method type.
	Bind(scope constructs.Construct) *SubjectAlternativeNamesMatcherConfig
}

// The jsii proxy struct for SubjectAlternativeNames
type jsiiProxy_SubjectAlternativeNames struct {
	_ byte // padding
}

func NewSubjectAlternativeNames_Override(s SubjectAlternativeNames) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.SubjectAlternativeNames",
		nil, // no parameters
		s,
	)
}

// The values of the SAN must match the specified values exactly.
func SubjectAlternativeNames_MatchingExactly(names ...*string) SubjectAlternativeNames {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range names {
		args = append(args, a)
	}

	var returns SubjectAlternativeNames

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.SubjectAlternativeNames",
		"matchingExactly",
		args,
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubjectAlternativeNames) Bind(scope constructs.Construct) *SubjectAlternativeNamesMatcherConfig {
	if err := s.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *SubjectAlternativeNamesMatcherConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

