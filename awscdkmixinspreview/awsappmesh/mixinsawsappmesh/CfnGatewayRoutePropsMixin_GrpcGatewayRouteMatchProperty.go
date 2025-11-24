package mixinsawsappmesh


// An object that represents the criteria for determining a request match.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   grpcGatewayRouteMatchProperty := &GrpcGatewayRouteMatchProperty{
//   	Hostname: &GatewayRouteHostnameMatchProperty{
//   		Exact: jsii.String("exact"),
//   		Suffix: jsii.String("suffix"),
//   	},
//   	Metadata: []interface{}{
//   		&GrpcGatewayRouteMetadataProperty{
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
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	Port: jsii.Number(123),
//   	ServiceName: jsii.String("serviceName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-grpcgatewayroutematch.html
//
type CfnGatewayRoutePropsMixin_GrpcGatewayRouteMatchProperty struct {
	// The gateway route host name to be matched on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-grpcgatewayroutematch.html#cfn-appmesh-gatewayroute-grpcgatewayroutematch-hostname
	//
	Hostname interface{} `field:"optional" json:"hostname" yaml:"hostname"`
	// The gateway route metadata to be matched on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-grpcgatewayroutematch.html#cfn-appmesh-gatewayroute-grpcgatewayroutematch-metadata
	//
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
	// The gateway route port to be matched on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-grpcgatewayroutematch.html#cfn-appmesh-gatewayroute-grpcgatewayroutematch-port
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The fully qualified domain name for the service to match from the request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-grpcgatewayroutematch.html#cfn-appmesh-gatewayroute-grpcgatewayroutematch-servicename
	//
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
}

