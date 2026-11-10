package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gatewayConnectorHealthCheckMatcherProperty := &GatewayConnectorHealthCheckMatcherProperty{
//   	GrpcCode: jsii.String("grpcCode"),
//   	HttpCode: jsii.String("httpCode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-gatewayconnectorhealthcheckmatcher.html
//
type CfnService_GatewayConnectorHealthCheckMatcherProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-gatewayconnectorhealthcheckmatcher.html#cfn-ecs-service-gatewayconnectorhealthcheckmatcher-grpccode
	//
	GrpcCode *string `field:"optional" json:"grpcCode" yaml:"grpcCode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-gatewayconnectorhealthcheckmatcher.html#cfn-ecs-service-gatewayconnectorhealthcheckmatcher-httpcode
	//
	HttpCode *string `field:"optional" json:"httpCode" yaml:"httpCode"`
}

