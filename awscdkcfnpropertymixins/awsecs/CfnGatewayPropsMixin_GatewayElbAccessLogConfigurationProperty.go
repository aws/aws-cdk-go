package awsecs


// The access log configuration for the gateway ALB.
//
// Access logging is enabled by the presence of a destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   gatewayElbAccessLogConfigurationProperty := &GatewayElbAccessLogConfigurationProperty{
//   	Destination: &GatewayElbAccessLogDestinationProperty{
//   		S3: &GatewayElbAccessLogS3DestinationProperty{
//   			Bucket: jsii.String("bucket"),
//   			Prefix: jsii.String("prefix"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-gateway-gatewayelbaccesslogconfiguration.html
//
type CfnGatewayPropsMixin_GatewayElbAccessLogConfigurationProperty struct {
	// The destination for gateway ALB access logs.
	//
	// Exactly one destination type must be specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-gateway-gatewayelbaccesslogconfiguration.html#cfn-ecs-gateway-gatewayelbaccesslogconfiguration-destination
	//
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
}

