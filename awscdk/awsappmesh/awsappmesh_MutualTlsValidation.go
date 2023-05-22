package awsappmesh


// Represents the properties needed to define TLS Validation context that is supported for mutual TLS authentication.
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
// Experimental.
type MutualTlsValidation struct {
	// Reference to where to retrieve the trust chain.
	// Experimental.
	Trust MutualTlsValidationTrust `field:"required" json:"trust" yaml:"trust"`
	// Represents the subject alternative names (SANs) secured by the certificate.
	//
	// SANs must be in the FQDN or URI format.
	// Experimental.
	SubjectAlternativeNames SubjectAlternativeNames `field:"optional" json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
}

