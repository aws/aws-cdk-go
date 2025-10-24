package awsappmesh


// Connection pool properties for HTTP listeners.
//
// Example:
//   // A Virtual Node with a gRPC listener with a connection pool set
//   var mesh Mesh
//
//   node := appmesh.NewVirtualNode(this, jsii.String("node"), &VirtualNodeProps{
//   	Mesh: Mesh,
//   	// DNS service discovery can optionally specify the DNS response type as either LOAD_BALANCER or ENDPOINTS.
//   	// LOAD_BALANCER means that the DNS resolver returns a loadbalanced set of endpoints,
//   	// whereas ENDPOINTS means that the DNS resolver is returning all the endpoints.
//   	// By default, the response type is assumed to be LOAD_BALANCER
//   	ServiceDiscovery: appmesh.ServiceDiscovery_Dns(jsii.String("node"), appmesh.DnsResponseType_ENDPOINTS),
//   	Listeners: []VirtualNodeListener{
//   		appmesh.VirtualNodeListener_Http(&HttpVirtualNodeListenerOptions{
//   			Port: jsii.Number(80),
//   			ConnectionPool: &HttpConnectionPool{
//   				MaxConnections: jsii.Number(100),
//   				MaxPendingRequests: jsii.Number(10),
//   			},
//   		}),
//   	},
//   })
//
//   // A Virtual Gateway with a gRPC listener with a connection pool set
//   gateway := appmesh.NewVirtualGateway(this, jsii.String("gateway"), &VirtualGatewayProps{
//   	Mesh: Mesh,
//   	Listeners: []VirtualGatewayListener{
//   		appmesh.VirtualGatewayListener_Grpc(&GrpcGatewayListenerOptions{
//   			Port: jsii.Number(8080),
//   			ConnectionPool: &GrpcConnectionPool{
//   				MaxRequests: jsii.Number(10),
//   			},
//   		}),
//   	},
//   	VirtualGatewayName: jsii.String("gateway"),
//   })
//
type HttpConnectionPool struct {
	// The maximum connections in the pool.
	// Default: - none.
	//
	MaxConnections *float64 `field:"required" json:"maxConnections" yaml:"maxConnections"`
	// The maximum pending requests in the pool.
	// Default: - none.
	//
	MaxPendingRequests *float64 `field:"required" json:"maxPendingRequests" yaml:"maxPendingRequests"`
}

