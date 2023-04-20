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

