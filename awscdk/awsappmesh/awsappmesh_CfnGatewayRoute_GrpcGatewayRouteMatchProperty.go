package awsappmesh


// An object that represents the criteria for determining a request match.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   grpcGatewayRouteMatchProperty := &grpcGatewayRouteMatchProperty{
//   	hostname: &gatewayRouteHostnameMatchProperty{
//   		exact: jsii.String("exact"),
//   		suffix: jsii.String("suffix"),
//   	},
//   	metadata: []interface{}{
//   		&grpcGatewayRouteMetadataProperty{
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			invert: jsii.Boolean(false),
//   			match: &gatewayRouteMetadataMatchProperty{
//   				exact: jsii.String("exact"),
//   				prefix: jsii.String("prefix"),
//   				range: &gatewayRouteRangeMatchProperty{
//   					end: jsii.Number(123),
//   					start: jsii.Number(123),
//   				},
//   				regex: jsii.String("regex"),
//   				suffix: jsii.String("suffix"),
//   			},
//   		},
//   	},
//   	port: jsii.Number(123),
//   	serviceName: jsii.String("serviceName"),
//   }
//
type CfnGatewayRoute_GrpcGatewayRouteMatchProperty struct {
	// The gateway route host name to be matched on.
	Hostname interface{} `field:"optional" json:"hostname" yaml:"hostname"`
	// The gateway route metadata to be matched on.
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
	// `CfnGatewayRoute.GrpcGatewayRouteMatchProperty.Port`.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The fully qualified domain name for the service to match from the request.
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
}

