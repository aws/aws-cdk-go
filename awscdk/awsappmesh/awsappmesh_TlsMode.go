package awsappmesh


// Enum of supported TLS modes.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   // A Virtual Node with listener TLS from an ACM provided certificate
//   var cert certificate
//   var mesh mesh
//
//
//   node := appmesh.NewVirtualNode(this, jsii.String("node"), &virtualNodeProps{
//   	mesh: mesh,
//   	serviceDiscovery: appmesh.serviceDiscovery.dns(jsii.String("node")),
//   	listeners: []virtualNodeListener{
//   		appmesh.*virtualNodeListener.grpc(&grpcVirtualNodeListenerOptions{
//   			port: jsii.Number(80),
//   			tls: &listenerTlsOptions{
//   				mode: appmesh.tlsMode_STRICT,
//   				certificate: appmesh.tlsCertificate.acm(cert),
//   			},
//   		}),
//   	},
//   })
//
//   // A Virtual Gateway with listener TLS from a customer provided file certificate
//   gateway := appmesh.NewVirtualGateway(this, jsii.String("gateway"), &virtualGatewayProps{
//   	mesh: mesh,
//   	listeners: []virtualGatewayListener{
//   		appmesh.*virtualGatewayListener.grpc(&grpcGatewayListenerOptions{
//   			port: jsii.Number(8080),
//   			tls: &listenerTlsOptions{
//   				mode: appmesh.*tlsMode_STRICT,
//   				certificate: appmesh.*tlsCertificate.file(jsii.String("path/to/certChain"), jsii.String("path/to/privateKey")),
//   			},
//   		}),
//   	},
//   	virtualGatewayName: jsii.String("gateway"),
//   })
//
//   // A Virtual Gateway with listener TLS from a SDS provided certificate
//   gateway2 := appmesh.NewVirtualGateway(this, jsii.String("gateway2"), &virtualGatewayProps{
//   	mesh: mesh,
//   	listeners: []*virtualGatewayListener{
//   		appmesh.*virtualGatewayListener.http2(&http2GatewayListenerOptions{
//   			port: jsii.Number(8080),
//   			tls: &listenerTlsOptions{
//   				mode: appmesh.*tlsMode_STRICT,
//   				certificate: appmesh.*tlsCertificate.sds(jsii.String("secrete_certificate")),
//   			},
//   		}),
//   	},
//   	virtualGatewayName: jsii.String("gateway2"),
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

