package previewawsappmeshmixins


// An object that represents a gRPC gateway route.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   grpcGatewayRouteProperty := &GrpcGatewayRouteProperty{
//   	Action: &GrpcGatewayRouteActionProperty{
//   		Rewrite: &GrpcGatewayRouteRewriteProperty{
//   			Hostname: &GatewayRouteHostnameRewriteProperty{
//   				DefaultTargetHostname: jsii.String("defaultTargetHostname"),
//   			},
//   		},
//   		Target: &GatewayRouteTargetProperty{
//   			Port: jsii.Number(123),
//   			VirtualService: &GatewayRouteVirtualServiceProperty{
//   				VirtualServiceName: jsii.String("virtualServiceName"),
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
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		Port: jsii.Number(123),
//   		ServiceName: jsii.String("serviceName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-grpcgatewayroute.html
//
type CfnGatewayRoutePropsMixin_GrpcGatewayRouteProperty struct {
	// An object that represents the action to take if a match is determined.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-grpcgatewayroute.html#cfn-appmesh-gatewayroute-grpcgatewayroute-action
	//
	Action interface{} `field:"optional" json:"action" yaml:"action"`
	// An object that represents the criteria for determining a request match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-grpcgatewayroute.html#cfn-appmesh-gatewayroute-grpcgatewayroute-match
	//
	Match interface{} `field:"optional" json:"match" yaml:"match"`
}

