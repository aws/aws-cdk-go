package awsappmesh


// Represents TLS properties for listener.
//
// Example:
//   // A Virtual Node with listener TLS from an ACM provided certificate
//   var cert certificate
//   var mesh mesh
//
//
//   node := appmesh.NewVirtualNode(this, jsii.String("node"), &VirtualNodeProps{
//   	Mesh: Mesh,
//   	ServiceDiscovery: appmesh.ServiceDiscovery_Dns(jsii.String("node")),
//   	Listeners: []virtualNodeListener{
//   		appmesh.*virtualNodeListener_Grpc(&GrpcVirtualNodeListenerOptions{
//   			Port: jsii.Number(80),
//   			Tls: &ListenerTlsOptions{
//   				Mode: appmesh.TlsMode_STRICT,
//   				Certificate: appmesh.TlsCertificate_Acm(cert),
//   			},
//   		}),
//   	},
//   })
//
//   // A Virtual Gateway with listener TLS from a customer provided file certificate
//   gateway := appmesh.NewVirtualGateway(this, jsii.String("gateway"), &VirtualGatewayProps{
//   	Mesh: Mesh,
//   	Listeners: []virtualGatewayListener{
//   		appmesh.*virtualGatewayListener_Grpc(&GrpcGatewayListenerOptions{
//   			Port: jsii.Number(8080),
//   			Tls: &ListenerTlsOptions{
//   				Mode: appmesh.TlsMode_STRICT,
//   				Certificate: appmesh.TlsCertificate_File(jsii.String("path/to/certChain"), jsii.String("path/to/privateKey")),
//   			},
//   		}),
//   	},
//   	VirtualGatewayName: jsii.String("gateway"),
//   })
//
//   // A Virtual Gateway with listener TLS from a SDS provided certificate
//   gateway2 := appmesh.NewVirtualGateway(this, jsii.String("gateway2"), &VirtualGatewayProps{
//   	Mesh: Mesh,
//   	Listeners: []*virtualGatewayListener{
//   		appmesh.*virtualGatewayListener_Http2(&Http2GatewayListenerOptions{
//   			Port: jsii.Number(8080),
//   			Tls: &ListenerTlsOptions{
//   				Mode: appmesh.TlsMode_STRICT,
//   				Certificate: appmesh.TlsCertificate_Sds(jsii.String("secrete_certificate")),
//   			},
//   		}),
//   	},
//   	VirtualGatewayName: jsii.String("gateway2"),
//   })
//
type ListenerTlsOptions struct {
	// Represents TLS certificate.
	Certificate TlsCertificate `field:"required" json:"certificate" yaml:"certificate"`
	// The TLS mode.
	Mode TlsMode `field:"required" json:"mode" yaml:"mode"`
	// Represents a listener's TLS validation context.
	//
	// The client certificate will only be validated if the client provides it, enabling mutual TLS.
	// Default: - client TLS certificate is not required.
	//
	MutualTlsValidation *MutualTlsValidation `field:"optional" json:"mutualTlsValidation" yaml:"mutualTlsValidation"`
}

