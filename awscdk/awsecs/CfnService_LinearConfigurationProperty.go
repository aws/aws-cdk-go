package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   linearConfigurationProperty := &LinearConfigurationProperty{
//   	StepBakeTimeInMinutes: jsii.Number(123),
//   	StepPercent: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-linearconfiguration.html
//
type CfnService_LinearConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-linearconfiguration.html#cfn-ecs-service-linearconfiguration-stepbaketimeinminutes
	//
	StepBakeTimeInMinutes *float64 `field:"optional" json:"stepBakeTimeInMinutes" yaml:"stepBakeTimeInMinutes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-linearconfiguration.html#cfn-ecs-service-linearconfiguration-steppercent
	//
	StepPercent *float64 `field:"optional" json:"stepPercent" yaml:"stepPercent"`
}

