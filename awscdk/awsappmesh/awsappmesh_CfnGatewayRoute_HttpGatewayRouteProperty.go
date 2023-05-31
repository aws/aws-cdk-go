package awsappmesh


// An object that represents an HTTP gateway route.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpGatewayRouteProperty := &HttpGatewayRouteProperty{
//   	Action: &HttpGatewayRouteActionProperty{
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
//   		Rewrite: &HttpGatewayRouteRewriteProperty{
//   			Hostname: &GatewayRouteHostnameRewriteProperty{
//   				DefaultTargetHostname: jsii.String("defaultTargetHostname"),
//   			},
//   			Path: &HttpGatewayRoutePathRewriteProperty{
//   				Exact: jsii.String("exact"),
//   			},
//   			Prefix: &HttpGatewayRoutePrefixRewriteProperty{
//   				DefaultPrefix: jsii.String("defaultPrefix"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	Match: &HttpGatewayRouteMatchProperty{
//   		Headers: []interface{}{
//   			&HttpGatewayRouteHeaderProperty{
//   				Name: jsii.String("name"),
//
//   				// the properties below are optional
//   				Invert: jsii.Boolean(false),
//   				Match: &HttpGatewayRouteHeaderMatchProperty{
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
//   		Hostname: &GatewayRouteHostnameMatchProperty{
//   			Exact: jsii.String("exact"),
//   			Suffix: jsii.String("suffix"),
//   		},
//   		Method: jsii.String("method"),
//   		Path: &HttpPathMatchProperty{
//   			Exact: jsii.String("exact"),
//   			Regex: jsii.String("regex"),
//   		},
//   		Port: jsii.Number(123),
//   		Prefix: jsii.String("prefix"),
//   		QueryParameters: []interface{}{
//   			&QueryParameterProperty{
//   				Name: jsii.String("name"),
//
//   				// the properties below are optional
//   				Match: &HttpQueryParameterMatchProperty{
//   					Exact: jsii.String("exact"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnGatewayRoute_HttpGatewayRouteProperty struct {
	// An object that represents the action to take if a match is determined.
	Action interface{} `field:"required" json:"action" yaml:"action"`
	// An object that represents the criteria for determining a request match.
	Match interface{} `field:"required" json:"match" yaml:"match"`
}

