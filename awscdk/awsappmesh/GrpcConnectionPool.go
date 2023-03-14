package awsappmesh


// Connection pool properties for gRPC listeners.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   // A Virtual Node with a gRPC listener with a connection pool set
//   var mesh mesh
//
//   node := appmesh.NewVirtualNode(this, jsii.String("node"), &VirtualNodeProps{
//   	Mesh: Mesh,
//   	// DNS service discovery can optionally specify the DNS response type as either LOAD_BALANCER or ENDPOINTS.
//   	// LOAD_BALANCER means that the DNS resolver returns a loadbalanced set of endpoints,
//   	// whereas ENDPOINTS means that the DNS resolver is returning all the endpoints.
//   	// By default, the response type is assumed to be LOAD_BALANCER
//   	ServiceDiscovery: appmesh.ServiceDiscovery_Dns(jsii.String("node"), appmesh.DnsResponseType_ENDPOINTS),
//   	Listeners: []virtualNodeListener{
//   		appmesh.*virtualNodeListener_Http(&HttpVirtualNodeListenerOptions{
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
//   	Listeners: []virtualGatewayListener{
//   		appmesh.*virtualGatewayListener_Grpc(&GrpcGatewayListenerOptions{
//   			Port: jsii.Number(8080),
//   			ConnectionPool: &GrpcConnectionPool{
//   				MaxRequests: jsii.Number(10),
//   			},
//   		}),
//   	},
//   	VirtualGatewayName: jsii.String("gateway"),
//   })
//
type GrpcConnectionPool struct {
	// The maximum requests in the pool.
	MaxRequests *float64 `field:"required" json:"maxRequests" yaml:"maxRequests"`
}

