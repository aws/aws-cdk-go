package awsappmesh


// Enum of supported TLS modes.
//
// Example:
//   // A Virtual Node with listener TLS from an ACM provided certificate
//   var cert Certificate
//   var mesh Mesh
//
//
//   node := appmesh.NewVirtualNode(this, jsii.String("node"), &VirtualNodeProps{
//   	Mesh: Mesh,
//   	ServiceDiscovery: appmesh.ServiceDiscovery_Dns(jsii.String("node")),
//   	Listeners: []VirtualNodeListener{
//   		appmesh.VirtualNodeListener_Grpc(&GrpcVirtualNodeListenerOptions{
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
//   	Listeners: []VirtualGatewayListener{
//   		appmesh.VirtualGatewayListener_Grpc(&GrpcGatewayListenerOptions{
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
//   	Listeners: []VirtualGatewayListener{
//   		appmesh.VirtualGatewayListener_Http2(&Http2GatewayListenerOptions{
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
type TlsMode string

const (
	// Only accept encrypted traffic.
	TlsMode_STRICT TlsMode = "STRICT"
	// Accept encrypted and plaintext traffic.
	TlsMode_PERMISSIVE TlsMode = "PERMISSIVE"
	// TLS is disabled, only accept plaintext traffic.
	TlsMode_DISABLED TlsMode = "DISABLED"
)

