package awsecs


// The host shared ALB configuration for the gateway.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hostSharedAlbConfigurationProperty := &HostSharedAlbConfigurationProperty{
//   	AccessLogs: &GatewayElbAccessLogConfigurationProperty{
//   		Destination: &GatewayElbAccessLogDestinationProperty{
//   			S3: &GatewayElbAccessLogS3DestinationProperty{
//   				Bucket: jsii.String("bucket"),
//
//   				// the properties below are optional
//   				Prefix: jsii.String("prefix"),
//   			},
//   		},
//   	},
//   	CurrentAlbCount: jsii.Number(123),
//   	HttpDesync: jsii.String("httpDesync"),
//   	MaxAlbCount: jsii.Number(123),
//   	NetworkConfiguration: &GatewayElbNetworkConfigurationProperty{
//   		InboundSecurityGroup: jsii.String("inboundSecurityGroup"),
//   		IpAddressType: jsii.String("ipAddressType"),
//   		Subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   	},
//   	Port: jsii.Number(123),
//   	Protocol: jsii.String("protocol"),
//   	Scheme: jsii.String("scheme"),
//   	TestTrafficRules: &GatewayAlbTestTrafficRulesProperty{
//   		Header: &GatewayTestTrafficHeaderProperty{
//   			Name: jsii.String("name"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	WebAclArn: jsii.String("webAclArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-gateway-hostsharedalbconfiguration.html
//
type CfnGateway_HostSharedAlbConfigurationProperty struct {
	// The access log configuration for the gateway ALB.
	//
	// Access logging is enabled by the presence of a destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-gateway-hostsharedalbconfiguration.html#cfn-ecs-gateway-hostsharedalbconfiguration-accesslogs
	//
	AccessLogs interface{} `field:"optional" json:"accessLogs" yaml:"accessLogs"`
	// The current number of ALBs provisioned for the gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-gateway-hostsharedalbconfiguration.html#cfn-ecs-gateway-hostsharedalbconfiguration-currentalbcount
	//
	CurrentAlbCount *float64 `field:"optional" json:"currentAlbCount" yaml:"currentAlbCount"`
	// The HTTP desync mitigation mode for the ALB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-gateway-hostsharedalbconfiguration.html#cfn-ecs-gateway-hostsharedalbconfiguration-httpdesync
	//
	HttpDesync *string `field:"optional" json:"httpDesync" yaml:"httpDesync"`
	// The maximum number of ALBs to provision for the gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-gateway-hostsharedalbconfiguration.html#cfn-ecs-gateway-hostsharedalbconfiguration-maxalbcount
	//
	MaxAlbCount *float64 `field:"optional" json:"maxAlbCount" yaml:"maxAlbCount"`
	// The network configuration for the gateway ALB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-gateway-hostsharedalbconfiguration.html#cfn-ecs-gateway-hostsharedalbconfiguration-networkconfiguration
	//
	NetworkConfiguration interface{} `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// The port on which the ALB listens.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-gateway-hostsharedalbconfiguration.html#cfn-ecs-gateway-hostsharedalbconfiguration-port
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The protocol for the ALB listener.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-gateway-hostsharedalbconfiguration.html#cfn-ecs-gateway-hostsharedalbconfiguration-protocol
	//
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The scheme for the ALB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-gateway-hostsharedalbconfiguration.html#cfn-ecs-gateway-hostsharedalbconfiguration-scheme
	//
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
	// The test traffic rules for the gateway ALB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-gateway-hostsharedalbconfiguration.html#cfn-ecs-gateway-hostsharedalbconfiguration-testtrafficrules
	//
	TestTrafficRules interface{} `field:"optional" json:"testTrafficRules" yaml:"testTrafficRules"`
	// The ARN of the WAFv2 web ACL to associate with the ALB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-gateway-hostsharedalbconfiguration.html#cfn-ecs-gateway-hostsharedalbconfiguration-webaclarn
	//
	WebAclArn *string `field:"optional" json:"webAclArn" yaml:"webAclArn"`
}

