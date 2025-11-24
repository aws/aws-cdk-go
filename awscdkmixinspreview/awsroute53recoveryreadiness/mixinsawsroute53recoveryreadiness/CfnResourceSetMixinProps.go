package mixinsawsroute53recoveryreadiness

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnResourceSetPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnResourceSetMixinProps := &CfnResourceSetMixinProps{
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
//   	ResourceSetName: jsii.String("resourceSetName"),
//   	ResourceSetType: jsii.String("resourceSetType"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53recoveryreadiness-resourceset.html
//
type CfnResourceSetMixinProps struct {
	// A list of resource objects in the resource set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53recoveryreadiness-resourceset.html#cfn-route53recoveryreadiness-resourceset-resources
	//
	Resources interface{} `field:"optional" json:"resources" yaml:"resources"`
	// The name of the resource set to create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53recoveryreadiness-resourceset.html#cfn-route53recoveryreadiness-resourceset-resourcesetname
	//
	ResourceSetName *string `field:"optional" json:"resourceSetName" yaml:"resourceSetName"`
	// The resource type of the resources in the resource set. Enter one of the following values for resource type:.
	//
	// AWS::ApiGateway::Stage, AWS::ApiGatewayV2::Stage, AWS::AutoScaling::AutoScalingGroup, AWS::CloudWatch::Alarm, AWS::EC2::CustomerGateway, AWS::DynamoDB::Table, AWS::EC2::Volume, AWS::ElasticLoadBalancing::LoadBalancer, AWS::ElasticLoadBalancingV2::LoadBalancer, AWS::Lambda::Function, AWS::MSK::Cluster, AWS::RDS::DBCluster, AWS::Route53::HealthCheck, AWS::SQS::Queue, AWS::SNS::Topic, AWS::SNS::Subscription, AWS::EC2::VPC, AWS::EC2::VPNConnection, AWS::EC2::VPNGateway, AWS::Route53RecoveryReadiness::DNSTargetResource.
	//
	// Note that AWS::Route53RecoveryReadiness::DNSTargetResource is only used for this setting. It isn't an actual AWS CloudFormation resource type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53recoveryreadiness-resourceset.html#cfn-route53recoveryreadiness-resourceset-resourcesettype
	//
	ResourceSetType *string `field:"optional" json:"resourceSetType" yaml:"resourceSetType"`
	// A tag to associate with the parameters for a resource set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53recoveryreadiness-resourceset.html#cfn-route53recoveryreadiness-resourceset-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

