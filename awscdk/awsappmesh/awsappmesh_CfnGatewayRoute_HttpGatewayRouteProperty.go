package awsappmesh


// An object that represents an HTTP gateway route.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpGatewayRouteProperty := &httpGatewayRouteProperty{
//   	action: &httpGatewayRouteActionProperty{
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
//   		rewrite: &httpGatewayRouteRewriteProperty{
//   			hostname: &gatewayRouteHostnameRewriteProperty{
//   				defaultTargetHostname: jsii.String("defaultTargetHostname"),
//   			},
//   			path: &httpGatewayRoutePathRewriteProperty{
//   				exact: jsii.String("exact"),
//   			},
//   			prefix: &httpGatewayRoutePrefixRewriteProperty{
//   				defaultPrefix: jsii.String("defaultPrefix"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	match: &httpGatewayRouteMatchProperty{
//   		headers: []interface{}{
//   			&httpGatewayRouteHeaderProperty{
//   				name: jsii.String("name"),
//
//   				// the properties below are optional
//   				invert: jsii.Boolean(false),
//   				match: &httpGatewayRouteHeaderMatchProperty{
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
//   		hostname: &gatewayRouteHostnameMatchProperty{
//   			exact: jsii.String("exact"),
//   			suffix: jsii.String("suffix"),
//   		},
//   		method: jsii.String("method"),
//   		path: &httpPathMatchProperty{
//   			exact: jsii.String("exact"),
//   			regex: jsii.String("regex"),
//   		},
//   		port: jsii.Number(123),
//   		prefix: jsii.String("prefix"),
//   		queryParameters: []interface{}{
//   			&queryParameterProperty{
//   				name: jsii.String("name"),
//
//   				// the properties below are optional
//   				match: &httpQueryParameterMatchProperty{
//   					exact: jsii.String("exact"),
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

