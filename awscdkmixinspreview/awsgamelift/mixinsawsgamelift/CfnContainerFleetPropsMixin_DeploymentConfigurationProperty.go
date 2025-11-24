package mixinsawsgamelift


// Set of rules for processing a deployment for a container fleet update.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   deploymentConfigurationProperty := &DeploymentConfigurationProperty{
//   	ImpairmentStrategy: jsii.String("impairmentStrategy"),
//   	MinimumHealthyPercentage: jsii.Number(123),
//   	ProtectionStrategy: jsii.String("protectionStrategy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-deploymentconfiguration.html
//
type CfnContainerFleetPropsMixin_DeploymentConfigurationProperty struct {
	// Determines what actions to take if a deployment fails.
	//
	// If the fleet is multi-location, this strategy applies across all fleet locations. With a rollback strategy, updated fleet instances are rolled back to the last successful deployment. Alternatively, you can maintain a few impaired containers for the purpose of debugging, while all other tasks return to the last successful deployment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-deploymentconfiguration.html#cfn-gamelift-containerfleet-deploymentconfiguration-impairmentstrategy
	//
	ImpairmentStrategy *string `field:"optional" json:"impairmentStrategy" yaml:"impairmentStrategy"`
	// Sets a minimum level of healthy tasks to maintain during deployment activity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-deploymentconfiguration.html#cfn-gamelift-containerfleet-deploymentconfiguration-minimumhealthypercentage
	//
	MinimumHealthyPercentage *float64 `field:"optional" json:"minimumHealthyPercentage" yaml:"minimumHealthyPercentage"`
	// Determines how fleet deployment activity affects active game sessions on the fleet.
	//
	// With protection, a deployment honors game session protection, and delays actions that would interrupt a protected active game session until the game session ends. Without protection, deployment activity can shut down all running tasks, including active game sessions, regardless of game session protection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-deploymentconfiguration.html#cfn-gamelift-containerfleet-deploymentconfiguration-protectionstrategy
	//
	ProtectionStrategy *string `field:"optional" json:"protectionStrategy" yaml:"protectionStrategy"`
}

