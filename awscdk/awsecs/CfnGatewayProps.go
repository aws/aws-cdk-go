package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnGateway`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGatewayProps := &CfnGatewayProps{
//   	Configuration: &GatewayConfigurationProperty{
//   		HostSharedAlb: &HostSharedAlbConfigurationProperty{
//   			AccessLogs: &GatewayElbAccessLogConfigurationProperty{
//   				Destination: &GatewayElbAccessLogDestinationProperty{
//   					S3: &GatewayElbAccessLogS3DestinationProperty{
//   						Bucket: jsii.String("bucket"),
//
//   						// the properties below are optional
//   						Prefix: jsii.String("prefix"),
//   					},
//   				},
//   			},
//   			CurrentAlbCount: jsii.Number(123),
//   			HttpDesync: jsii.String("httpDesync"),
//   			MaxAlbCount: jsii.Number(123),
//   			NetworkConfiguration: &GatewayElbNetworkConfigurationProperty{
//   				InboundSecurityGroup: jsii.String("inboundSecurityGroup"),
//   				IpAddressType: jsii.String("ipAddressType"),
//   				Subnets: []*string{
//   					jsii.String("subnets"),
//   				},
//   			},
//   			Port: jsii.Number(123),
//   			Protocol: jsii.String("protocol"),
//   			Scheme: jsii.String("scheme"),
//   			TestTrafficRules: &GatewayAlbTestTrafficRulesProperty{
//   				Header: &GatewayTestTrafficHeaderProperty{
//   					Name: jsii.String("name"),
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   			},
//   			WebAclArn: jsii.String("webAclArn"),
//   		},
//   	},
//   	GatewayName: jsii.String("gatewayName"),
//   	InfrastructureRoleArn: jsii.String("infrastructureRoleArn"),
//
//   	// the properties below are optional
//   	DeletionProtection: jsii.Boolean(false),
//   	PropagateTags: jsii.String("propagateTags"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-gateway.html
//
type CfnGatewayProps struct {
	// The configuration for the gateway.
	//
	// Currently only supports hostSharedAlb.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-gateway.html#cfn-ecs-gateway-configuration
	//
	Configuration interface{} `field:"required" json:"configuration" yaml:"configuration"`
	// The name of the gateway.
	//
	// Up to 255 letters (uppercase and lowercase), numbers, and hyphens are allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-gateway.html#cfn-ecs-gateway-gatewayname
	//
	GatewayName *string `field:"required" json:"gatewayName" yaml:"gatewayName"`
	// The ARN of the IAM role that ECS uses to manage infrastructure for the gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-gateway.html#cfn-ecs-gateway-infrastructurerolearn
	//
	InfrastructureRoleArn *string `field:"required" json:"infrastructureRoleArn" yaml:"infrastructureRoleArn"`
	// Whether deletion protection is enabled for the gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-gateway.html#cfn-ecs-gateway-deletionprotection
	//
	// Default: - true.
	//
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Specifies whether to propagate the tags from the gateway to other resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-gateway.html#cfn-ecs-gateway-propagatetags
	//
	PropagateTags *string `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// The metadata that you apply to the gateway to help you categorize and organize them.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-gateway.html#cfn-ecs-gateway-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

