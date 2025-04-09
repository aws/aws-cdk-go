package awsappmesh


// Represents the properties needed to define client policy.
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
type TlsClientPolicy struct {
	// Represents the object for TLS validation context.
	Validation *TlsValidation `field:"required" json:"validation" yaml:"validation"`
	// Whether the policy is enforced.
	// Default: true.
	//
	Enforce *bool `field:"optional" json:"enforce" yaml:"enforce"`
	// Represents a client TLS certificate.
	//
	// The certificate will be sent only if the server requests it, enabling mutual TLS.
	// Default: - client TLS certificate is not provided.
	//
	MutualTlsCertificate MutualTlsCertificate `field:"optional" json:"mutualTlsCertificate" yaml:"mutualTlsCertificate"`
	// TLS is enforced on the ports specified here.
	//
	// If no ports are specified, TLS will be enforced on all the ports.
	// Default: - all ports.
	//
	Ports *[]*float64 `field:"optional" json:"ports" yaml:"ports"`
}

