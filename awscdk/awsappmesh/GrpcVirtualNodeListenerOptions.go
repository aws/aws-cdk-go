package awsappmesh


// Represent the GRPC Node Listener property.
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
type GrpcVirtualNodeListenerOptions struct {
	// Connection pool for http listeners.
	// Default: - None.
	//
	ConnectionPool *GrpcConnectionPool `field:"optional" json:"connectionPool" yaml:"connectionPool"`
	// The health check information for the listener.
	// Default: - no healthcheck.
	//
	HealthCheck HealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// Represents the configuration for enabling outlier detection.
	// Default: - none.
	//
	OutlierDetection *OutlierDetection `field:"optional" json:"outlierDetection" yaml:"outlierDetection"`
	// Port to listen for connections on.
	// Default: - 8080.
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Timeout for GRPC protocol.
	// Default: - None.
	//
	Timeout *GrpcTimeout `field:"optional" json:"timeout" yaml:"timeout"`
	// Represents the configuration for enabling TLS on a listener.
	// Default: - none.
	//
	Tls *ListenerTlsOptions `field:"optional" json:"tls" yaml:"tls"`
}

