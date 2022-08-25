package awsappmesh

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Used to generate Subject Alternative Names Matchers.
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
	var returns *SubjectAlternativeNamesMatcherConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

