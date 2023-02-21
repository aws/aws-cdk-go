package awsappmesh


// An object that represents a gRPC gateway route.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   grpcGatewayRouteProperty := &grpcGatewayRouteProperty{
//   	action: &grpcGatewayRouteActionProperty{
//   		target: &gatewayRouteTargetProperty{
//   			virtualService: &gatewayRouteVirtualServiceProperty{
//   				virtualServiceName: jsii.String("virtualServiceName"),
//   			},
//
//   			// the properties below are optional
//   			port: jsii.Number(123),
//   		},
//
//   		// the properties below are optional
//   		rewrite: &grpcGatewayRouteRewriteProperty{
//   			hostname: &gatewayRouteHostnameRewriteProperty{
//   				defaultTargetHostname: jsii.String("defaultTargetHostname"),
//   			},
//   		},
//   	},
//   	match: &grpcGatewayRouteMatchProperty{
//   		hostname: &gatewayRouteHostnameMatchProperty{
//   			exact: jsii.String("exact"),
//   			suffix: jsii.String("suffix"),
//   		},
//   		metadata: []interface{}{
//   			&grpcGatewayRouteMetadataProperty{
//   				name: jsii.String("name"),
//
//   				// the properties below are optional
//   				invert: jsii.Boolean(false),
//   				match: &gatewayRouteMetadataMatchProperty{
//   					exact: jsii.String("exact"),
//   					prefix: jsii.String("prefix"),
//   					range: &gatewayRouteRangeMatchProperty{
//   						end: jsii.Number(123),
//   						start: jsii.Number(123),
//   					},
//   					regex: jsii.String("regex"),
//   					suffix: jsii.String("suffix"),
//   				},
//   			},
//   		},
//   		port: jsii.Number(123),
//   		serviceName: jsii.String("serviceName"),
//   	},
//   }
//
type CfnGatewayRoute_GrpcGatewayRouteProperty struct {
	// An object that represents the action to take if a match is determined.
	Action interface{} `field:"required" json:"action" yaml:"action"`
	// An object that represents the criteria for determining a request match.
	Match interface{} `field:"required" json:"match" yaml:"match"`
}

