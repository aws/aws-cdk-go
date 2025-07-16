package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   advancedConfigurationProperty := &AdvancedConfigurationProperty{
//   	AlternateTargetGroupArn: jsii.String("alternateTargetGroupArn"),
//
//   	// the properties below are optional
//   	ProductionListenerRule: jsii.String("productionListenerRule"),
//   	RoleArn: jsii.String("roleArn"),
//   	TestListenerRule: jsii.String("testListenerRule"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-advancedconfiguration.html
//
type CfnService_AdvancedConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-advancedconfiguration.html#cfn-ecs-service-advancedconfiguration-alternatetargetgrouparn
	//
	AlternateTargetGroupArn *string `field:"required" json:"alternateTargetGroupArn" yaml:"alternateTargetGroupArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-advancedconfiguration.html#cfn-ecs-service-advancedconfiguration-productionlistenerrule
	//
	ProductionListenerRule *string `field:"optional" json:"productionListenerRule" yaml:"productionListenerRule"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-advancedconfiguration.html#cfn-ecs-service-advancedconfiguration-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-advancedconfiguration.html#cfn-ecs-service-advancedconfiguration-testlistenerrule
	//
	TestListenerRule *string `field:"optional" json:"testListenerRule" yaml:"testListenerRule"`
}

