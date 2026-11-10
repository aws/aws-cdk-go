package awsecs


// The advanced settings for VPC Lattice used in blue/green deployments.
//
// Specify the alternate target group and listener rules required for traffic shifting during blue/green deployments. For more information, see [Required resources for Amazon ECS blue/green deployments](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/blue-green-deployment-implementation.html) in the *Amazon Elastic Container Service Developer Guide*.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   vpcLatticeAdvancedConfigurationProperty := &VpcLatticeAdvancedConfigurationProperty{
//   	AlternateTargetGroupArn: jsii.String("alternateTargetGroupArn"),
//   	ProductionListenerRule: jsii.String("productionListenerRule"),
//   	TestListenerRule: jsii.String("testListenerRule"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-vpclatticeadvancedconfiguration.html
//
type CfnServicePropsMixin_VpcLatticeAdvancedConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the alternate target group associated with the VPC Lattice Configuration for Amazon ECS blue/green deployments.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-vpclatticeadvancedconfiguration.html#cfn-ecs-service-vpclatticeadvancedconfiguration-alternatetargetgrouparn
	//
	AlternateTargetGroupArn *string `field:"optional" json:"alternateTargetGroupArn" yaml:"alternateTargetGroupArn"`
	// The Amazon Resource Name (ARN) that identifies the production listener rule or listener for routing production traffic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-vpclatticeadvancedconfiguration.html#cfn-ecs-service-vpclatticeadvancedconfiguration-productionlistenerrule
	//
	ProductionListenerRule *string `field:"optional" json:"productionListenerRule" yaml:"productionListenerRule"`
	// The Amazon Resource Name (ARN) that identifies the test listener rule or listener for routing test traffic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-vpclatticeadvancedconfiguration.html#cfn-ecs-service-vpclatticeadvancedconfiguration-testlistenerrule
	//
	TestListenerRule *string `field:"optional" json:"testListenerRule" yaml:"testListenerRule"`
}

