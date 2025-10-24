package awsappmesh


// Properties used when creating a new VirtualGateway.
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
type VirtualGatewayProps struct {
	// Access Logging Configuration for the VirtualGateway.
	// Default: - no access logging.
	//
	AccessLog AccessLog `field:"optional" json:"accessLog" yaml:"accessLog"`
	// Default Configuration Virtual Node uses to communicate with Virtual Service.
	// Default: - No Config.
	//
	BackendDefaults *BackendDefaults `field:"optional" json:"backendDefaults" yaml:"backendDefaults"`
	// Listeners for the VirtualGateway.
	//
	// Only one is supported.
	// Default: - Single HTTP listener on port 8080.
	//
	Listeners *[]VirtualGatewayListener `field:"optional" json:"listeners" yaml:"listeners"`
	// Name of the VirtualGateway.
	// Default: - A name is automatically determined.
	//
	VirtualGatewayName *string `field:"optional" json:"virtualGatewayName" yaml:"virtualGatewayName"`
	// The Mesh which the VirtualGateway belongs to.
	Mesh IMesh `field:"required" json:"mesh" yaml:"mesh"`
}

