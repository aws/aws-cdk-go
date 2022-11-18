package awsappmesh


// Represents the properties needed to define client policy.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var mesh mesh
//   var service service
//
//
//   node := appmesh.NewVirtualNode(this, jsii.String("node"), &virtualNodeProps{
//   	mesh: mesh,
//   	serviceDiscovery: appmesh.serviceDiscovery.cloudMap(service),
//   	listeners: []virtualNodeListener{
//   		appmesh.*virtualNodeListener.http(&httpVirtualNodeListenerOptions{
//   			port: jsii.Number(8080),
//   			healthCheck: appmesh.healthCheck.http(&httpHealthCheckOptions{
//   				healthyThreshold: jsii.Number(3),
//   				interval: cdk.duration.seconds(jsii.Number(5)),
//   				path: jsii.String("/ping"),
//   				timeout: cdk.*duration.seconds(jsii.Number(2)),
//   				unhealthyThreshold: jsii.Number(2),
//   			}),
//   			timeout: &httpTimeout{
//   				idle: cdk.*duration.seconds(jsii.Number(5)),
//   			},
//   		}),
//   	},
//   	backendDefaults: &backendDefaults{
//   		tlsClientPolicy: &tlsClientPolicy{
//   			validation: &tlsValidation{
//   				trust: appmesh.tlsValidationTrust.file(jsii.String("/keys/local_cert_chain.pem")),
//   			},
//   		},
//   	},
//   	accessLog: appmesh.accessLog.fromFilePath(jsii.String("/dev/stdout")),
//   })
//
//   cdk.tags.of(node).add(jsii.String("Environment"), jsii.String("Dev"))
//
type TlsClientPolicy struct {
	// Represents the object for TLS validation context.
	Validation *TlsValidation `field:"required" json:"validation" yaml:"validation"`
	// Whether the policy is enforced.
	Enforce *bool `field:"optional" json:"enforce" yaml:"enforce"`
	// Represents a client TLS certificate.
	//
	// The certificate will be sent only if the server requests it, enabling mutual TLS.
	MutualTlsCertificate MutualTlsCertificate `field:"optional" json:"mutualTlsCertificate" yaml:"mutualTlsCertificate"`
	// TLS is enforced on the ports specified here.
	//
	// If no ports are specified, TLS will be enforced on all the ports.
	Ports *[]*float64 `field:"optional" json:"ports" yaml:"ports"`
}

