package awsappmesh


// Represents the properties needed to define TLS Validation context.
//
// Example:
//   var mesh mesh
//   var service service
//
//
//   node := appmesh.NewVirtualNode(this, jsii.String("node"), &VirtualNodeProps{
//   	Mesh: Mesh,
//   	ServiceDiscovery: appmesh.ServiceDiscovery_CloudMap(service),
//   	Listeners: []virtualNodeListener{
//   		appmesh.*virtualNodeListener_Http(&HttpVirtualNodeListenerOptions{
//   			Port: jsii.Number(8080),
//   			HealthCheck: appmesh.HealthCheck_Http(&HttpHealthCheckOptions{
//   				HealthyThreshold: jsii.Number(3),
//   				Interval: awscdk.Duration_Seconds(jsii.Number(5)),
//   				Path: jsii.String("/ping"),
//   				Timeout: awscdk.Duration_*Seconds(jsii.Number(2)),
//   				UnhealthyThreshold: jsii.Number(2),
//   			}),
//   			Timeout: &HttpTimeout{
//   				Idle: awscdk.Duration_*Seconds(jsii.Number(5)),
//   			},
//   		}),
//   	},
//   	BackendDefaults: &BackendDefaults{
//   		TlsClientPolicy: &TlsClientPolicy{
//   			Validation: &TlsValidation{
//   				Trust: appmesh.TlsValidationTrust_File(jsii.String("/keys/local_cert_chain.pem")),
//   			},
//   		},
//   	},
//   	AccessLog: appmesh.AccessLog_FromFilePath(jsii.String("/dev/stdout")),
//   })
//
//   cdk.Tags_Of(node).Add(jsii.String("Environment"), jsii.String("Dev"))
//
type TlsValidation struct {
	// Reference to where to retrieve the trust chain.
	Trust TlsValidationTrust `field:"required" json:"trust" yaml:"trust"`
	// Represents the subject alternative names (SANs) secured by the certificate.
	//
	// SANs must be in the FQDN or URI format.
	// Default: - If you don't specify SANs on the terminating mesh endpoint,
	// the Envoy proxy for that node doesn't verify the SAN on a peer client certificate.
	// If you don't specify SANs on the originating mesh endpoint,
	// the SAN on the certificate provided by the terminating endpoint must match the mesh endpoint service discovery configuration.
	//
	SubjectAlternativeNames SubjectAlternativeNames `field:"optional" json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
}

