package awsgamelift


// Provides details about how to drain old tasks and replace them with new updated tasks.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deploymentConfigurationProperty := &DeploymentConfigurationProperty{
//   	ImpairmentStrategy: jsii.String("impairmentStrategy"),
//   	MinimumHealthyPercentage: jsii.Number(123),
//   	ProtectionStrategy: jsii.String("protectionStrategy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-deploymentconfiguration.html
//
type CfnContainerFleet_DeploymentConfigurationProperty struct {
	// The strategy to apply in case of impairment;
	//
	// defaults to MAINTAIN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-deploymentconfiguration.html#cfn-gamelift-containerfleet-deploymentconfiguration-impairmentstrategy
	//
	ImpairmentStrategy *string `field:"optional" json:"impairmentStrategy" yaml:"impairmentStrategy"`
	// The minimum percentage of healthy required;
	//
	// defaults to 75.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-deploymentconfiguration.html#cfn-gamelift-containerfleet-deploymentconfiguration-minimumhealthypercentage
	//
	MinimumHealthyPercentage *float64 `field:"optional" json:"minimumHealthyPercentage" yaml:"minimumHealthyPercentage"`
	// The protection strategy for deployment on the container fleet;
	//
	// defaults to WITH_PROTECTION.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-deploymentconfiguration.html#cfn-gamelift-containerfleet-deploymentconfiguration-protectionstrategy
	//
	ProtectionStrategy *string `field:"optional" json:"protectionStrategy" yaml:"protectionStrategy"`
}

