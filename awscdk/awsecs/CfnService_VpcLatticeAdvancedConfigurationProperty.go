package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcLatticeAdvancedConfigurationProperty := &VpcLatticeAdvancedConfigurationProperty{
//   	AlternateTargetGroupArn: jsii.String("alternateTargetGroupArn"),
//   	ProductionListenerRule: jsii.String("productionListenerRule"),
//   	TestListenerRule: jsii.String("testListenerRule"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-vpclatticeadvancedconfiguration.html
//
type CfnService_VpcLatticeAdvancedConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-vpclatticeadvancedconfiguration.html#cfn-ecs-service-vpclatticeadvancedconfiguration-alternatetargetgrouparn
	//
	AlternateTargetGroupArn *string `field:"optional" json:"alternateTargetGroupArn" yaml:"alternateTargetGroupArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-vpclatticeadvancedconfiguration.html#cfn-ecs-service-vpclatticeadvancedconfiguration-productionlistenerrule
	//
	ProductionListenerRule *string `field:"optional" json:"productionListenerRule" yaml:"productionListenerRule"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-vpclatticeadvancedconfiguration.html#cfn-ecs-service-vpclatticeadvancedconfiguration-testlistenerrule
	//
	TestListenerRule *string `field:"optional" json:"testListenerRule" yaml:"testListenerRule"`
}

