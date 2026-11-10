package awsecs


// The configuration for the gateway.
//
// Currently only supports hostSharedAlb.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   gatewayConfigurationProperty := &GatewayConfigurationProperty{
//   	HostSharedAlb: &HostSharedAlbConfigurationProperty{
//   		AccessLogs: &GatewayElbAccessLogConfigurationProperty{
//   			Destination: &GatewayElbAccessLogDestinationProperty{
//   				S3: &GatewayElbAccessLogS3DestinationProperty{
//   					Bucket: jsii.String("bucket"),
//   					Prefix: jsii.String("prefix"),
//   				},
//   			},
//   		},
//   		CurrentAlbCount: jsii.Number(123),
//   		HttpDesync: jsii.String("httpDesync"),
//   		MaxAlbCount: jsii.Number(123),
//   		NetworkConfiguration: &GatewayElbNetworkConfigurationProperty{
//   			InboundSecurityGroup: jsii.String("inboundSecurityGroup"),
//   			IpAddressType: jsii.String("ipAddressType"),
//   			Subnets: []*string{
//   				jsii.String("subnets"),
//   			},
//   		},
//   		Port: jsii.Number(123),
//   		Protocol: jsii.String("protocol"),
//   		Scheme: jsii.String("scheme"),
//   		TestTrafficRules: &GatewayAlbTestTrafficRulesProperty{
//   			Header: &GatewayTestTrafficHeaderProperty{
//   				Name: jsii.String("name"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   		WebAclArn: jsii.String("webAclArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-gateway-gatewayconfiguration.html
//
type CfnGatewayPropsMixin_GatewayConfigurationProperty struct {
	// The host shared ALB configuration for the gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-gateway-gatewayconfiguration.html#cfn-ecs-gateway-gatewayconfiguration-hostsharedalb
	//
	HostSharedAlb interface{} `field:"optional" json:"hostSharedAlb" yaml:"hostSharedAlb"`
}

