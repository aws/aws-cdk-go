package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gatewayConnectorHealthCheckProperty := &GatewayConnectorHealthCheckProperty{
//   	HealthCheckIntervalSeconds: jsii.Number(123),
//   	HealthCheckPath: jsii.String("healthCheckPath"),
//   	HealthCheckPort: jsii.Number(123),
//   	HealthCheckProtocol: jsii.String("healthCheckProtocol"),
//   	HealthCheckTimeoutSeconds: jsii.Number(123),
//   	HealthyThresholdCount: jsii.Number(123),
//   	Matcher: &GatewayConnectorHealthCheckMatcherProperty{
//   		GrpcCode: jsii.String("grpcCode"),
//   		HttpCode: jsii.String("httpCode"),
//   	},
//   	UnhealthyThresholdCount: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-gatewayconnectorhealthcheck.html
//
type CfnService_GatewayConnectorHealthCheckProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-gatewayconnectorhealthcheck.html#cfn-ecs-service-gatewayconnectorhealthcheck-healthcheckintervalseconds
	//
	HealthCheckIntervalSeconds *float64 `field:"optional" json:"healthCheckIntervalSeconds" yaml:"healthCheckIntervalSeconds"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-gatewayconnectorhealthcheck.html#cfn-ecs-service-gatewayconnectorhealthcheck-healthcheckpath
	//
	HealthCheckPath *string `field:"optional" json:"healthCheckPath" yaml:"healthCheckPath"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-gatewayconnectorhealthcheck.html#cfn-ecs-service-gatewayconnectorhealthcheck-healthcheckport
	//
	HealthCheckPort *float64 `field:"optional" json:"healthCheckPort" yaml:"healthCheckPort"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-gatewayconnectorhealthcheck.html#cfn-ecs-service-gatewayconnectorhealthcheck-healthcheckprotocol
	//
	HealthCheckProtocol *string `field:"optional" json:"healthCheckProtocol" yaml:"healthCheckProtocol"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-gatewayconnectorhealthcheck.html#cfn-ecs-service-gatewayconnectorhealthcheck-healthchecktimeoutseconds
	//
	HealthCheckTimeoutSeconds *float64 `field:"optional" json:"healthCheckTimeoutSeconds" yaml:"healthCheckTimeoutSeconds"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-gatewayconnectorhealthcheck.html#cfn-ecs-service-gatewayconnectorhealthcheck-healthythresholdcount
	//
	HealthyThresholdCount *float64 `field:"optional" json:"healthyThresholdCount" yaml:"healthyThresholdCount"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-gatewayconnectorhealthcheck.html#cfn-ecs-service-gatewayconnectorhealthcheck-matcher
	//
	Matcher interface{} `field:"optional" json:"matcher" yaml:"matcher"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-gatewayconnectorhealthcheck.html#cfn-ecs-service-gatewayconnectorhealthcheck-unhealthythresholdcount
	//
	UnhealthyThresholdCount *float64 `field:"optional" json:"unhealthyThresholdCount" yaml:"unhealthyThresholdCount"`
}

