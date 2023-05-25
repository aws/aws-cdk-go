package awsappmesh


// An object that represents a gRPC gateway route.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   grpcGatewayRouteProperty := &GrpcGatewayRouteProperty{
//   	Action: &GrpcGatewayRouteActionProperty{
//   		Target: &GatewayRouteTargetProperty{
//   			VirtualService: &GatewayRouteVirtualServiceProperty{
//   				VirtualServiceName: jsii.String("virtualServiceName"),
//   			},
//
//   			// the properties below are optional
//   			Port: jsii.Number(123),
//   		},
//
//   		// the properties below are optional
//   		Rewrite: &GrpcGatewayRouteRewriteProperty{
//   			Hostname: &GatewayRouteHostnameRewriteProperty{
//   				DefaultTargetHostname: jsii.String("defaultTargetHostname"),
//   			},
//   		},
//   	},
//   	Match: &GrpcGatewayRouteMatchProperty{
//   		Hostname: &GatewayRouteHostnameMatchProperty{
//   			Exact: jsii.String("exact"),
//   			Suffix: jsii.String("suffix"),
//   		},
//   		Metadata: []interface{}{
//   			&GrpcGatewayRouteMetadataProperty{
//   				Name: jsii.String("name"),
//
//   				// the properties below are optional
//   				Invert: jsii.Boolean(false),
//   				Match: &GatewayRouteMetadataMatchProperty{
//   					Exact: jsii.String("exact"),
//   					Prefix: jsii.String("prefix"),
//   					Range: &GatewayRouteRangeMatchProperty{
//   						End: jsii.Number(123),
//   						Start: jsii.Number(123),
//   					},
//   					Regex: jsii.String("regex"),
//   					Suffix: jsii.String("suffix"),
//   				},
//   			},
//   		},
//   		Port: jsii.Number(123),
//   		ServiceName: jsii.String("serviceName"),
//   	},
//   }
//
type CfnGatewayRoute_GrpcGatewayRouteProperty struct {
	// An object that represents the action to take if a match is determined.
	Action interface{} `field:"required" json:"action" yaml:"action"`
	// An object that represents the criteria for determining a request match.
	Match interface{} `field:"required" json:"match" yaml:"match"`
}

