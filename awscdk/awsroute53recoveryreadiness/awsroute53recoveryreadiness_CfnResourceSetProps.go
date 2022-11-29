package awsroute53recoveryreadiness

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnResourceSet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResourceSetProps := &cfnResourceSetProps{
//   	resources: []interface{}{
//   		&resourceProperty{
//   			componentId: jsii.String("componentId"),
//   			dnsTargetResource: &dNSTargetResourceProperty{
//   				domainName: jsii.String("domainName"),
//   				hostedZoneArn: jsii.String("hostedZoneArn"),
//   				recordSetId: jsii.String("recordSetId"),
//   				recordType: jsii.String("recordType"),
//   				targetResource: &targetResourceProperty{
//   					nlbResource: &nLBResourceProperty{
//   						arn: jsii.String("arn"),
//   					},
//   					r53Resource: &r53ResourceRecordProperty{
//   						domainName: jsii.String("domainName"),
//   						recordSetId: jsii.String("recordSetId"),
//   					},
//   				},
//   			},
//   			readinessScopes: []*string{
//   				jsii.String("readinessScopes"),
//   			},
//   			resourceArn: jsii.String("resourceArn"),
//   		},
//   	},
//   	resourceSetType: jsii.String("resourceSetType"),
//
//   	// the properties below are optional
//   	resourceSetName: jsii.String("resourceSetName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnResourceSetProps struct {
	// A list of resource objects in the resource set.
	Resources interface{} `field:"required" json:"resources" yaml:"resources"`
	// The resource type of the resources in the resource set. Enter one of the following values for resource type:.
	//
	// AWS::AutoScaling::AutoScalingGroup, AWS::CloudWatch::Alarm, AWS::EC2::CustomerGateway, AWS::DynamoDB::Table, AWS::EC2::Volume, AWS::ElasticLoadBalancing::LoadBalancer, AWS::ElasticLoadBalancingV2::LoadBalancer, AWS::MSK::Cluster, AWS::RDS::DBCluster, AWS::Route53::HealthCheck, AWS::SQS::Queue, AWS::SNS::Topic, AWS::SNS::Subscription, AWS::EC2::VPC, AWS::EC2::VPNConnection, AWS::EC2::VPNGateway, AWS::Route53RecoveryReadiness::DNSTargetResource.
	//
	// Note that AWS::Route53RecoveryReadiness::DNSTargetResource is only used for this setting. It isn't an actual AWS CloudFormation resource type.
	ResourceSetType *string `field:"required" json:"resourceSetType" yaml:"resourceSetType"`
	// The name of the resource set to create.
	ResourceSetName *string `field:"optional" json:"resourceSetName" yaml:"resourceSetName"`
	// A tag to associate with the parameters for a resource set.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

