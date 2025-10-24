package awsappmesh


// Represents the properties needed to define HTTP Listeners for a VirtualGateway.
//
// Example:
//   var mesh Mesh
//
//   certificateAuthorityArn := "arn:aws:acm-pca:us-east-1:123456789012:certificate-authority/12345678-1234-1234-1234-123456789012"
//
//   gateway := appmesh.NewVirtualGateway(this, jsii.String("gateway"), &VirtualGatewayProps{
//   	Mesh: mesh,
//   	Listeners: []VirtualGatewayListener{
//   		appmesh.VirtualGatewayListener_Http(&HttpGatewayListenerOptions{
//   			Port: jsii.Number(443),
//   			HealthCheck: appmesh.HealthCheck_Http(&HttpHealthCheckOptions{
//   				Interval: awscdk.Duration_Seconds(jsii.Number(10)),
//   			}),
//   		}),
//   	},
//   	BackendDefaults: &BackendDefaults{
//   		TlsClientPolicy: &TlsClientPolicy{
//   			Ports: []*f64{
//   				jsii.Number(8080),
//   				jsii.Number(8081),
//   			},
//   			Validation: &TlsValidation{
//   				Trust: appmesh.TlsValidationTrust_Acm([]ICertificateAuthority{
//   					acmpca.CertificateAuthority_FromCertificateAuthorityArn(this, jsii.String("certificate"), certificateAuthorityArn),
//   				}),
//   			},
//   		},
//   	},
//   	AccessLog: appmesh.AccessLog_FromFilePath(jsii.String("/dev/stdout")),
//   	VirtualGatewayName: jsii.String("virtualGateway"),
//   })
//
type HttpGatewayListenerOptions struct {
	// Connection pool for http listeners.
	// Default: - None.
	//
	ConnectionPool *HttpConnectionPool `field:"optional" json:"connectionPool" yaml:"connectionPool"`
	// The health check information for the listener.
	// Default: - no healthcheck.
	//
	HealthCheck HealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// Port to listen for connections on.
	// Default: - 8080.
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Represents the configuration for enabling TLS on a listener.
	// Default: - none.
	//
	Tls *ListenerTlsOptions `field:"optional" json:"tls" yaml:"tls"`
}

