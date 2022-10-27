package awsappmesh


// The criterion for determining a request match for this GatewayRoute.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var gateway virtualGateway
//   var virtualService virtualService
//
//
//   gateway.addGatewayRoute(jsii.String("gateway-route-grpc"), &gatewayRouteBaseProps{
//   	routeSpec: appmesh.gatewayRouteSpec.grpc(&grpcGatewayRouteSpecOptions{
//   		routeTarget: virtualService,
//   		match: &grpcGatewayRouteMatch{
//   			hostname: appmesh.gatewayRouteHostnameMatch.endsWith(jsii.String(".example.com")),
//   		},
//   	}),
//   })
//
type GrpcGatewayRouteMatch struct {
	// Create host name based gRPC gateway route match.
	Hostname GatewayRouteHostnameMatch `field:"optional" json:"hostname" yaml:"hostname"`
	// Create metadata based gRPC gateway route match.
	//
	// All specified metadata must match for the route to match.
	Metadata *[]HeaderMatch `field:"optional" json:"metadata" yaml:"metadata"`
	// When `true`, rewrites the original request received at the Virtual Gateway to the destination Virtual Service name.
	//
	// When `false`, retains the original hostname from the request.
	RewriteRequestHostname *bool `field:"optional" json:"rewriteRequestHostname" yaml:"rewriteRequestHostname"`
	// Create service name based gRPC gateway route match.
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
}

