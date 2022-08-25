package awsappmesh


// Connection pool properties for HTTP listeners.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   // A Virtual Node with a gRPC listener with a connection pool set
//   var mesh mesh
//
//   node := appmesh.NewVirtualNode(this, jsii.String("node"), &virtualNodeProps{
//   	mesh: mesh,
//   	// DNS service discovery can optionally specify the DNS response type as either LOAD_BALANCER or ENDPOINTS.
//   	// LOAD_BALANCER means that the DNS resolver returns a loadbalanced set of endpoints,
//   	// whereas ENDPOINTS means that the DNS resolver is returning all the endpoints.
//   	// By default, the response type is assumed to be LOAD_BALANCER
//   	serviceDiscovery: appmesh.serviceDiscovery.dns(jsii.String("node"), appmesh.dnsResponseType_ENDPOINTS),
//   	listeners: []virtualNodeListener{
//   		appmesh.*virtualNodeListener.http(&httpVirtualNodeListenerOptions{
//   			port: jsii.Number(80),
//   			connectionPool: &httpConnectionPool{
//   				maxConnections: jsii.Number(100),
//   				maxPendingRequests: jsii.Number(10),
//   			},
//   		}),
//   	},
//   })
//
//   // A Virtual Gateway with a gRPC listener with a connection pool set
//   gateway := appmesh.NewVirtualGateway(this, jsii.String("gateway"), &virtualGatewayProps{
//   	mesh: mesh,
//   	listeners: []virtualGatewayListener{
//   		appmesh.*virtualGatewayListener.grpc(&grpcGatewayListenerOptions{
//   			port: jsii.Number(8080),
//   			connectionPool: &grpcConnectionPool{
//   				maxRequests: jsii.Number(10),
//   			},
//   		}),
//   	},
//   	virtualGatewayName: jsii.String("gateway"),
//   })
//
type HttpConnectionPool struct {
	// The maximum connections in the pool.
	MaxConnections *float64 `field:"required" json:"maxConnections" yaml:"maxConnections"`
	// The maximum pending requests in the pool.
	MaxPendingRequests *float64 `field:"required" json:"maxPendingRequests" yaml:"maxPendingRequests"`
}

