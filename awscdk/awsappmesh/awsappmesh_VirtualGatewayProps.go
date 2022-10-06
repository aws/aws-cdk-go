package awsappmesh


// Properties used when creating a new VirtualGateway.
//
// Example:
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
// Experimental.
type VirtualGatewayProps struct {
	// Access Logging Configuration for the VirtualGateway.
	// Experimental.
	AccessLog AccessLog `field:"optional" json:"accessLog" yaml:"accessLog"`
	// Default Configuration Virtual Node uses to communicate with Virtual Service.
	// Experimental.
	BackendDefaults *BackendDefaults `field:"optional" json:"backendDefaults" yaml:"backendDefaults"`
	// Listeners for the VirtualGateway.
	//
	// Only one is supported.
	// Experimental.
	Listeners *[]VirtualGatewayListener `field:"optional" json:"listeners" yaml:"listeners"`
	// Name of the VirtualGateway.
	// Experimental.
	VirtualGatewayName *string `field:"optional" json:"virtualGatewayName" yaml:"virtualGatewayName"`
	// The Mesh which the VirtualGateway belongs to.
	// Experimental.
	Mesh IMesh `field:"required" json:"mesh" yaml:"mesh"`
}

