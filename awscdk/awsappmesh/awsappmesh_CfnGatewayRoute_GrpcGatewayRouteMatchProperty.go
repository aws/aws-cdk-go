package awsappmesh


// An object that represents the criteria for determining a request match.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   grpcGatewayRouteMatchProperty := &GrpcGatewayRouteMatchProperty{
//   	Hostname: &GatewayRouteHostnameMatchProperty{
//   		Exact: jsii.String("exact"),
//   		Suffix: jsii.String("suffix"),
//   	},
//   	Metadata: []interface{}{
//   		&GrpcGatewayRouteMetadataProperty{
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			Invert: jsii.Boolean(false),
//   			Match: &GatewayRouteMetadataMatchProperty{
//   				Exact: jsii.String("exact"),
//   				Prefix: jsii.String("prefix"),
//   				Range: &GatewayRouteRangeMatchProperty{
//   					End: jsii.Number(123),
//   					Start: jsii.Number(123),
//   				},
//   				Regex: jsii.String("regex"),
//   				Suffix: jsii.String("suffix"),
//   			},
//   		},
//   	},
//   	Port: jsii.Number(123),
//   	ServiceName: jsii.String("serviceName"),
//   }
//
type CfnGatewayRoute_GrpcGatewayRouteMatchProperty struct {
	// The gateway route host name to be matched on.
	Hostname interface{} `field:"optional" json:"hostname" yaml:"hostname"`
	// The gateway route metadata to be matched on.
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
	// The gateway route port to be matched on.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The fully qualified domain name for the service to match from the request.
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
}

