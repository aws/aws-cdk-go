package awsroute53recoveryreadiness

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnResourceSet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResourceSetProps := &CfnResourceSetProps{
//   	Resources: []interface{}{
//   		&ResourceProperty{
//   			ComponentId: jsii.String("componentId"),
//   			DnsTargetResource: &DNSTargetResourceProperty{
//   				DomainName: jsii.String("domainName"),
//   				HostedZoneArn: jsii.String("hostedZoneArn"),
//   				RecordSetId: jsii.String("recordSetId"),
//   				RecordType: jsii.String("recordType"),
//   				TargetResource: &TargetResourceProperty{
//   					NlbResource: &NLBResourceProperty{
//   						Arn: jsii.String("arn"),
//   					},
//   					R53Resource: &R53ResourceRecordProperty{
//   						DomainName: jsii.String("domainName"),
//   						RecordSetId: jsii.String("recordSetId"),
//   					},
//   				},
//   			},
//   			ReadinessScopes: []*string{
//   				jsii.String("readinessScopes"),
//   			},
//   			ResourceArn: jsii.String("resourceArn"),
//   		},
//   	},
//   	ResourceSetType: jsii.String("resourceSetType"),
//
//   	// the properties below are optional
//   	ResourceSetName: jsii.String("resourceSetName"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnResourceSetProps struct {
	// A list of resource objects in the resource set.
	Resources interface{} `field:"required" json:"resources" yaml:"resources"`
	// The resource type of the resources in the resource set. Enter one of the following values for resource type:.
	//
	// AWS::ApiGateway::Stage, AWS::ApiGatewayV2::Stage, AWS::AutoScaling::AutoScalingGroup, AWS::CloudWatch::Alarm, AWS::EC2::CustomerGateway, AWS::DynamoDB::Table, AWS::EC2::Volume, AWS::ElasticLoadBalancing::LoadBalancer, AWS::ElasticLoadBalancingV2::LoadBalancer, AWS::Lambda::Function, AWS::MSK::Cluster, AWS::RDS::DBCluster, AWS::Route53::HealthCheck, AWS::SQS::Queue, AWS::SNS::Topic, AWS::SNS::Subscription, AWS::EC2::VPC, AWS::EC2::VPNConnection, AWS::EC2::VPNGateway, AWS::Route53RecoveryReadiness::DNSTargetResource.
	//
	// Note that AWS::Route53RecoveryReadiness::DNSTargetResource is only used for this setting. It isn't an actual AWS CloudFormation resource type.
	ResourceSetType *string `field:"required" json:"resourceSetType" yaml:"resourceSetType"`
	// The name of the resource set to create.
	ResourceSetName *string `field:"optional" json:"resourceSetName" yaml:"resourceSetName"`
	// A tag to associate with the parameters for a resource set.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

