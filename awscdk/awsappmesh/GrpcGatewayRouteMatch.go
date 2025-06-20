package awsappmesh


// The criterion for determining a request match for this GatewayRoute.
//
// Example:
//   var gateway virtualGateway
//   var virtualService virtualService
//
//
//   gateway.addGatewayRoute(jsii.String("gateway-route-grpc"), &GatewayRouteBaseProps{
//   	RouteSpec: appmesh.GatewayRouteSpec_Grpc(&GrpcGatewayRouteSpecOptions{
//   		RouteTarget: virtualService,
//   		Match: &GrpcGatewayRouteMatch{
//   			Hostname: appmesh.GatewayRouteHostnameMatch_EndsWith(jsii.String(".example.com")),
//   		},
//   	}),
//   })
//
type GrpcGatewayRouteMatch struct {
	// Create host name based gRPC gateway route match.
	// Default: - no matching on host name.
	//
	Hostname GatewayRouteHostnameMatch `field:"optional" json:"hostname" yaml:"hostname"`
	// Create metadata based gRPC gateway route match.
	//
	// All specified metadata must match for the route to match.
	// Default: - no matching on metadata.
	//
	Metadata *[]HeaderMatch `field:"optional" json:"metadata" yaml:"metadata"`
	// The port to match from the request.
	// Default: - do not match on port.
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// When `true`, rewrites the original request received at the Virtual Gateway to the destination Virtual Service name.
	//
	// When `false`, retains the original hostname from the request.
	// Default: true.
	//
	RewriteRequestHostname *bool `field:"optional" json:"rewriteRequestHostname" yaml:"rewriteRequestHostname"`
	// Create service name based gRPC gateway route match.
	// Default: - no matching on service name.
	//
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
}

