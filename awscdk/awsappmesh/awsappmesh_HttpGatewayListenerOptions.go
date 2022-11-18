package awsappmesh


// Represents the properties needed to define HTTP Listeners for a VirtualGateway.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var mesh mesh
//
//   certificateAuthorityArn := "arn:aws:acm-pca:us-east-1:123456789012:certificate-authority/12345678-1234-1234-1234-123456789012"
//
//   gateway := appmesh.NewVirtualGateway(this, jsii.String("gateway"), &virtualGatewayProps{
//   	mesh: mesh,
//   	listeners: []virtualGatewayListener{
//   		appmesh.*virtualGatewayListener.http(&httpGatewayListenerOptions{
//   			port: jsii.Number(443),
//   			healthCheck: appmesh.healthCheck.http(&httpHealthCheckOptions{
//   				interval: cdk.duration.seconds(jsii.Number(10)),
//   			}),
//   		}),
//   	},
//   	backendDefaults: &backendDefaults{
//   		tlsClientPolicy: &tlsClientPolicy{
//   			ports: []*f64{
//   				jsii.Number(8080),
//   				jsii.Number(8081),
//   			},
//   			validation: &tlsValidation{
//   				trust: appmesh.tlsValidationTrust.acm([]iCertificateAuthority{
//   					acmpca.certificateAuthority.fromCertificateAuthorityArn(this, jsii.String("certificate"), certificateAuthorityArn),
//   				}),
//   			},
//   		},
//   	},
//   	accessLog: appmesh.accessLog.fromFilePath(jsii.String("/dev/stdout")),
//   	virtualGatewayName: jsii.String("virtualGateway"),
//   })
//
type HttpGatewayListenerOptions struct {
	// Connection pool for http listeners.
	ConnectionPool *HttpConnectionPool `field:"optional" json:"connectionPool" yaml:"connectionPool"`
	// The health check information for the listener.
	HealthCheck HealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// Port to listen for connections on.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Represents the configuration for enabling TLS on a listener.
	Tls *ListenerTlsOptions `field:"optional" json:"tls" yaml:"tls"`
}

