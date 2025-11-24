package mixinsawsecs


// The advanced settings for a load balancer used in blue/green deployments.
//
// Specify the alternate target group, listener rules, and IAM role required for traffic shifting during blue/green deployments. For more information, see [Required resources for Amazon ECS blue/green deployments](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/blue-green-deployment-implementation.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   advancedConfigurationProperty := &AdvancedConfigurationProperty{
//   	AlternateTargetGroupArn: jsii.String("alternateTargetGroupArn"),
//   	ProductionListenerRule: jsii.String("productionListenerRule"),
//   	RoleArn: jsii.String("roleArn"),
//   	TestListenerRule: jsii.String("testListenerRule"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-advancedconfiguration.html
//
type CfnServicePropsMixin_AdvancedConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the alternate target group for Amazon ECS blue/green deployments.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-advancedconfiguration.html#cfn-ecs-service-advancedconfiguration-alternatetargetgrouparn
	//
	AlternateTargetGroupArn *string `field:"optional" json:"alternateTargetGroupArn" yaml:"alternateTargetGroupArn"`
	// The Amazon Resource Name (ARN) that that identifies the production listener rule (in the case of an Application Load Balancer) or listener (in the case for an Network Load Balancer) for routing production traffic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-advancedconfiguration.html#cfn-ecs-service-advancedconfiguration-productionlistenerrule
	//
	ProductionListenerRule *string `field:"optional" json:"productionListenerRule" yaml:"productionListenerRule"`
	// The Amazon Resource Name (ARN) of the IAM role that grants Amazon ECS permission to call the Elastic Load Balancing APIs for you.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-advancedconfiguration.html#cfn-ecs-service-advancedconfiguration-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The Amazon Resource Name (ARN) that identifies ) that identifies the test listener rule (in the case of an Application Load Balancer) or listener (in the case for an Network Load Balancer) for routing test traffic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-advancedconfiguration.html#cfn-ecs-service-advancedconfiguration-testlistenerrule
	//
	TestListenerRule *string `field:"optional" json:"testListenerRule" yaml:"testListenerRule"`
}

