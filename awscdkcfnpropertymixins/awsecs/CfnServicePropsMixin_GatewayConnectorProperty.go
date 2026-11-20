package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   gatewayConnectorProperty := &GatewayConnectorProperty{
//   	DeregistrationDelay: &GatewayConnectorDeregistrationDelayProperty{
//   		TimeoutSeconds: jsii.Number(123),
//   	},
//   	Filters: &GatewayConnectorFiltersProperty{
//   		AccessTypes: []*string{
//   			jsii.String("accessTypes"),
//   		},
//   	},
//   	HealthCheck: &GatewayConnectorHealthCheckProperty{
//   		HealthCheckIntervalSeconds: jsii.Number(123),
//   		HealthCheckPath: jsii.String("healthCheckPath"),
//   		HealthCheckPort: jsii.Number(123),
//   		HealthCheckProtocol: jsii.String("healthCheckProtocol"),
//   		HealthCheckTimeoutSeconds: jsii.Number(123),
//   		HealthyThresholdCount: jsii.Number(123),
//   		Matcher: &GatewayConnectorHealthCheckMatcherProperty{
//   			GrpcCode: jsii.String("grpcCode"),
//   			HttpCode: jsii.String("httpCode"),
//   		},
//   		UnhealthyThresholdCount: jsii.Number(123),
//   	},
//   	PortName: jsii.String("portName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-gatewayconnector.html
//
type CfnServicePropsMixin_GatewayConnectorProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-gatewayconnector.html#cfn-ecs-service-gatewayconnector-deregistrationdelay
	//
	DeregistrationDelay interface{} `field:"optional" json:"deregistrationDelay" yaml:"deregistrationDelay"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-gatewayconnector.html#cfn-ecs-service-gatewayconnector-filters
	//
	Filters interface{} `field:"optional" json:"filters" yaml:"filters"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-gatewayconnector.html#cfn-ecs-service-gatewayconnector-healthcheck
	//
	HealthCheck interface{} `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-gatewayconnector.html#cfn-ecs-service-gatewayconnector-portname
	//
	PortName *string `field:"optional" json:"portName" yaml:"portName"`
}

